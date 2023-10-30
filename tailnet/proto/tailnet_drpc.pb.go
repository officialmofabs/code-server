// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.33
// source: tailnet/proto/tailnet.proto

package proto

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_tailnet_proto_tailnet_proto struct{}

func (drpcEncoding_File_tailnet_proto_tailnet_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_tailnet_proto_tailnet_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_tailnet_proto_tailnet_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_tailnet_proto_tailnet_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_tailnet_proto_tailnet_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCClientClient interface {
	DRPCConn() drpc.Conn

	StreamDERPMaps(ctx context.Context, in *StreamDERPMapsRequest) (DRPCClient_StreamDERPMapsClient, error)
	CoordinateTailnet(ctx context.Context) (DRPCClient_CoordinateTailnetClient, error)
}

type drpcClientClient struct {
	cc drpc.Conn
}

func NewDRPCClientClient(cc drpc.Conn) DRPCClientClient {
	return &drpcClientClient{cc}
}

func (c *drpcClientClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcClientClient) StreamDERPMaps(ctx context.Context, in *StreamDERPMapsRequest) (DRPCClient_StreamDERPMapsClient, error) {
	stream, err := c.cc.NewStream(ctx, "/coder.tailnet.v2.Client/StreamDERPMaps", drpcEncoding_File_tailnet_proto_tailnet_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcClient_StreamDERPMapsClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_tailnet_proto_tailnet_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCClient_StreamDERPMapsClient interface {
	drpc.Stream
	Recv() (*DERPMap, error)
}

type drpcClient_StreamDERPMapsClient struct {
	drpc.Stream
}

func (x *drpcClient_StreamDERPMapsClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcClient_StreamDERPMapsClient) Recv() (*DERPMap, error) {
	m := new(DERPMap)
	if err := x.MsgRecv(m, drpcEncoding_File_tailnet_proto_tailnet_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcClient_StreamDERPMapsClient) RecvMsg(m *DERPMap) error {
	return x.MsgRecv(m, drpcEncoding_File_tailnet_proto_tailnet_proto{})
}

func (c *drpcClientClient) CoordinateTailnet(ctx context.Context) (DRPCClient_CoordinateTailnetClient, error) {
	stream, err := c.cc.NewStream(ctx, "/coder.tailnet.v2.Client/CoordinateTailnet", drpcEncoding_File_tailnet_proto_tailnet_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcClient_CoordinateTailnetClient{stream}
	return x, nil
}

type DRPCClient_CoordinateTailnetClient interface {
	drpc.Stream
	Send(*CoordinateRequest) error
	Recv() (*CoordinateResponse, error)
}

type drpcClient_CoordinateTailnetClient struct {
	drpc.Stream
}

func (x *drpcClient_CoordinateTailnetClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcClient_CoordinateTailnetClient) Send(m *CoordinateRequest) error {
	return x.MsgSend(m, drpcEncoding_File_tailnet_proto_tailnet_proto{})
}

func (x *drpcClient_CoordinateTailnetClient) Recv() (*CoordinateResponse, error) {
	m := new(CoordinateResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_tailnet_proto_tailnet_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcClient_CoordinateTailnetClient) RecvMsg(m *CoordinateResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_tailnet_proto_tailnet_proto{})
}

type DRPCClientServer interface {
	StreamDERPMaps(*StreamDERPMapsRequest, DRPCClient_StreamDERPMapsStream) error
	CoordinateTailnet(DRPCClient_CoordinateTailnetStream) error
}

type DRPCClientUnimplementedServer struct{}

func (s *DRPCClientUnimplementedServer) StreamDERPMaps(*StreamDERPMapsRequest, DRPCClient_StreamDERPMapsStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientUnimplementedServer) CoordinateTailnet(DRPCClient_CoordinateTailnetStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCClientDescription struct{}

func (DRPCClientDescription) NumMethods() int { return 2 }

func (DRPCClientDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/coder.tailnet.v2.Client/StreamDERPMaps", drpcEncoding_File_tailnet_proto_tailnet_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCClientServer).
					StreamDERPMaps(
						in1.(*StreamDERPMapsRequest),
						&drpcClient_StreamDERPMapsStream{in2.(drpc.Stream)},
					)
			}, DRPCClientServer.StreamDERPMaps, true
	case 1:
		return "/coder.tailnet.v2.Client/CoordinateTailnet", drpcEncoding_File_tailnet_proto_tailnet_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCClientServer).
					CoordinateTailnet(
						&drpcClient_CoordinateTailnetStream{in1.(drpc.Stream)},
					)
			}, DRPCClientServer.CoordinateTailnet, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterClient(mux drpc.Mux, impl DRPCClientServer) error {
	return mux.Register(impl, DRPCClientDescription{})
}

type DRPCClient_StreamDERPMapsStream interface {
	drpc.Stream
	Send(*DERPMap) error
}

type drpcClient_StreamDERPMapsStream struct {
	drpc.Stream
}

func (x *drpcClient_StreamDERPMapsStream) Send(m *DERPMap) error {
	return x.MsgSend(m, drpcEncoding_File_tailnet_proto_tailnet_proto{})
}

type DRPCClient_CoordinateTailnetStream interface {
	drpc.Stream
	Send(*CoordinateResponse) error
	Recv() (*CoordinateRequest, error)
}

type drpcClient_CoordinateTailnetStream struct {
	drpc.Stream
}

func (x *drpcClient_CoordinateTailnetStream) Send(m *CoordinateResponse) error {
	return x.MsgSend(m, drpcEncoding_File_tailnet_proto_tailnet_proto{})
}

func (x *drpcClient_CoordinateTailnetStream) Recv() (*CoordinateRequest, error) {
	m := new(CoordinateRequest)
	if err := x.MsgRecv(m, drpcEncoding_File_tailnet_proto_tailnet_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcClient_CoordinateTailnetStream) RecvMsg(m *CoordinateRequest) error {
	return x.MsgRecv(m, drpcEncoding_File_tailnet_proto_tailnet_proto{})
}
