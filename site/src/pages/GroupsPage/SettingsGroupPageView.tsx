import TextField from "@mui/material/TextField";
import type { Group } from "api/typesGenerated";
import { Button } from "components/Button/Button";
import { FormFooter } from "components/Form/Form";
import { FullPageForm } from "components/FullPageForm/FullPageForm";
import { IconField } from "components/IconField/IconField";
import { Loader } from "components/Loader/Loader";
import { Margins } from "components/Margins/Margins";
import { Spinner } from "components/Spinner/Spinner";
import { Stack } from "components/Stack/Stack";
import { useFormik } from "formik";
import type { FC } from "react";
import {
	getFormHelpers,
	nameValidator,
	onChangeTrimmed,
} from "utils/formUtils";
import { isEveryoneGroup } from "utils/groups";
import * as Yup from "yup";

type FormData = {
	name: string;
	display_name: string;
	avatar_url: string;
	quota_allowance: number;
};

const validationSchema = Yup.object({
	name: nameValidator("Name"),
	quota_allowance: Yup.number().required().min(0).integer(),
});

interface UpdateGroupFormProps {
	group: Group;
	errors: unknown;
	onSubmit: (data: FormData) => void;
	onCancel: () => void;
	isLoading: boolean;
}

const UpdateGroupForm: FC<UpdateGroupFormProps> = ({
	group,
	errors,
	onSubmit,
	onCancel,
	isLoading,
}) => {
	const form = useFormik<FormData>({
		initialValues: {
			name: group.name,
			display_name: group.display_name,
			avatar_url: group.avatar_url,
			quota_allowance: group.quota_allowance,
		},
		validationSchema,
		onSubmit,
	});
	const getFieldHelpers = getFormHelpers<FormData>(form, errors);

	return (
		<FullPageForm title="Group settings">
			<form onSubmit={form.handleSubmit}>
				<Stack spacing={2.5}>
					<TextField
						{...getFieldHelpers("name")}
						onChange={onChangeTrimmed(form)}
						autoComplete="name"
						autoFocus
						fullWidth
						label="Name"
						disabled={isEveryoneGroup(group)}
					/>
					{isEveryoneGroup(group) ? (
						<></>
					) : (
						<>
							<TextField
								{...getFieldHelpers("display_name", {
									helperText: "Optional: keep empty to default to the name.",
								})}
								onChange={onChangeTrimmed(form)}
								autoComplete="display_name"
								autoFocus
								fullWidth
								label="Display Name"
								disabled={isEveryoneGroup(group)}
							/>
							<IconField
								{...getFieldHelpers("avatar_url")}
								onChange={onChangeTrimmed(form)}
								fullWidth
								label="Avatar URL"
								onPickEmoji={(value) => form.setFieldValue("avatar_url", value)}
							/>
						</>
					)}
					<TextField
						{...getFieldHelpers("quota_allowance", {
							helperText: `This group gives ${form.values.quota_allowance} quota credits to each
            of its members.`,
						})}
						onChange={onChangeTrimmed(form)}
						autoFocus
						fullWidth
						type="number"
						label="Quota Allowance"
					/>
				</Stack>

				<FormFooter>
					<Button onClick={onCancel} variant="outline">
						Cancel
					</Button>

					<Button type="submit" disabled={isLoading}>
						<Spinner loading={isLoading} />
						Save
					</Button>
				</FormFooter>
			</form>
		</FullPageForm>
	);
};

export type SettingsGroupPageViewProps = {
	onCancel: () => void;
	onSubmit: (data: FormData) => void;
	group: Group | undefined;
	formErrors: unknown;
	isLoading: boolean;
	isUpdating: boolean;
};

export const SettingsGroupPageView: FC<SettingsGroupPageViewProps> = ({
	onCancel,
	onSubmit,
	group,
	formErrors,
	isLoading,
	isUpdating,
}) => {
	if (isLoading) {
		return <Loader />;
	}

	return (
		<Margins>
			<UpdateGroupForm
				group={group!}
				onCancel={onCancel}
				errors={formErrors}
				isLoading={isUpdating}
				onSubmit={onSubmit}
			/>
		</Margins>
	);
};

export default SettingsGroupPageView;
