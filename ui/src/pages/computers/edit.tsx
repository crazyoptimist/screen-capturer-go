import { HttpError } from "@refinedev/core";
import { Edit } from "@refinedev/mui";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import { useForm } from "@refinedev/react-hook-form";

import { IComputer, Nullable } from "../../interfaces";

export const ComputerEdit: React.FC = () => {
  const {
    saveButtonProps,
    register,
    formState: { errors },
  } = useForm<IComputer, HttpError, Nullable<IComputer>>();

  return (
    <Edit saveButtonProps={saveButtonProps}>
      <Box
        component="form"
        sx={{ display: "flex", flexDirection: "column" }}
        autoComplete="off"
      >
        <TextField
          {...register("name")}
          error={!!errors.name}
          helperText={errors.name?.message}
          margin="normal"
          fullWidth
          label="Name"
          InputLabelProps={{ shrink: true }}
          autoFocus
        />
      </Box>
    </Edit>
  );
};
