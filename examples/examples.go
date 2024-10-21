package examples

import (
	"archive/tar"
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"path"
	"sort"
	"strings"
	"sync"

	"golang.org/x/sync/singleflight"
	"golang.org/x/xerrors"

	"github.com/coder/coder/v2/codersdk"
)

var (
	// Only some templates are embedded that we want to display inside the UI.
	// The metadata in examples.gen.json is generated via scripts/examplegen.
	//go:embed examples.gen.json
	//go:embed templates/aws-devcontainer
	//go:embed templates/aws-linux
	//go:embed templates/aws-windows
	//go:embed templates/azure-linux
	//go:embed templates/do-linux
	//go:embed templates/docker
	//go:embed templates/devcontainer-docker
	//go:embed templates/devcontainer-kubernetes
	//go:embed templates/gcp-devcontainer
	//go:embed templates/gcp-linux
	//go:embed templates/gcp-vm-container
	//go:embed templates/gcp-windows
	//go:embed templates/kubernetes
	//go:embed templates/nomad-docker
	//go:embed templates/scratch
	files embed.FS

	exampleBasePath = "https://github.com/coder/coder/tree/main/examples/templates/"
	examplesJSON    = "examples.gen.json"
	parsedExamples  []codersdk.TemplateExample
	parseExamples   sync.Once
	archives        singleflight.Group
	ErrNotFound     = xerrors.New("example not found")
)

const rootDir = "templates"

// List returns all embedded examples.
func List() ([]codersdk.TemplateExample, error) {
	var err error
	parseExamples.Do(func() {
		parsedExamples, err = parseAndVerifyExamples()
	})
	return parsedExamples, err
}

func parseAndVerifyExamples() (examples []codersdk.TemplateExample, err error) {
	f, err := files.Open(examplesJSON)
	if err != nil {
		return nil, xerrors.Errorf("open %s: %w", examplesJSON, err)
	}
	defer f.Close()

	b := bufio.NewReader(f)

	// Discard the first line (code generated by-comment).
	_, _, err = b.ReadLine()
	if err != nil {
		return nil, xerrors.Errorf("read %s: %w", examplesJSON, err)
	}

	err = json.NewDecoder(b).Decode(&examples)
	if err != nil {
		return nil, xerrors.Errorf("decode %s: %w", examplesJSON, err)
	}

	// Sanity-check: Verify that the examples in the JSON file match the
	// embedded files.
	var wantEmbedFiles []string
	for i, example := range examples {
		examples[i].URL = exampleBasePath + example.ID
		wantEmbedFiles = append(wantEmbedFiles, example.ID)
	}

	files, err := fs.Sub(files, rootDir)
	if err != nil {
		return nil, xerrors.Errorf("get templates fs: %w", err)
	}
	dirs, err := fs.ReadDir(files, ".")
	if err != nil {
		return nil, xerrors.Errorf("read templates dir: %w", err)
	}
	var gotEmbedFiles []string
	for _, dir := range dirs {
		if dir.IsDir() {
			gotEmbedFiles = append(gotEmbedFiles, dir.Name())
		}
	}

	sort.Strings(wantEmbedFiles)
	sort.Strings(gotEmbedFiles)
	want := strings.Join(wantEmbedFiles, ", ")
	got := strings.Join(gotEmbedFiles, ", ")
	if want != got {
		return nil, xerrors.Errorf("mismatch between %s and embedded files: want %q, got %q", examplesJSON, want, got)
	}

	return examples, nil
}

// Archive returns a tar by example ID.
func Archive(exampleID string) ([]byte, error) {
	rawData, err, _ := archives.Do(exampleID, func() (interface{}, error) {
		examples, err := List()
		if err != nil {
			return nil, xerrors.Errorf("list: %w", err)
		}

		var selected codersdk.TemplateExample
		for _, example := range examples {
			if example.ID != exampleID {
				continue
			}
			selected = example
			break
		}

		if selected.ID == "" {
			return nil, xerrors.Errorf("example with id %q not found: %w", exampleID, ErrNotFound)
		}

		exampleFiles, err := fs.Sub(files, path.Join(rootDir, exampleID))
		if err != nil {
			return nil, xerrors.Errorf("get example fs: %w", err)
		}

		var buffer bytes.Buffer
		tarWriter := tar.NewWriter(&buffer)

		err = fs.WalkDir(exampleFiles, ".", func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if path == "." {
				// Tar files don't have a root directory.
				return nil
			}

			info, err := entry.Info()
			if err != nil {
				return xerrors.Errorf("stat file: %w", err)
			}

			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return xerrors.Errorf("get file header: %w", err)
			}
			header.Name = strings.TrimPrefix(path, "./")
			header.Mode = 0o644

			if entry.IsDir() {
				// Trailing slash on entry name is not required. Our tar
				// creation code for tarring up a local directory doesn't
				// include slashes so this we don't include them here for
				// consistency.
				// header.Name += "/"
				header.Mode = 0o755
				header.Typeflag = tar.TypeDir
				err = tarWriter.WriteHeader(header)
				if err != nil {
					return xerrors.Errorf("write file: %w", err)
				}
			} else {
				file, err := exampleFiles.Open(path)
				if err != nil {
					return xerrors.Errorf("open file %s: %w", path, err)
				}
				defer file.Close()

				err = tarWriter.WriteHeader(header)
				if err != nil {
					return xerrors.Errorf("write file: %w", err)
				}

				_, err = io.Copy(tarWriter, file)
				if err != nil {
					return xerrors.Errorf("write: %w", err)
				}
			}
			return nil
		})
		if err != nil {
			return nil, xerrors.Errorf("walk example directory: %w", err)
		}

		err = tarWriter.Close()
		if err != nil {
			return nil, xerrors.Errorf("close archive: %w", err)
		}

		return buffer.Bytes(), nil
	})
	if err != nil {
		return nil, err
	}
	data, valid := rawData.([]byte)
	if !valid {
		panic("dev error: data must be a byte slice")
	}
	return data, nil
}
