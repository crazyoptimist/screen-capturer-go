import { HttpError } from "@refinedev/core";
import { Create } from "@refinedev/mui";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import { useForm } from "@refinedev/react-hook-form";

import { Nullable, IComputer } from "../../interfaces";

export const ComputerCreate: React.FC = () => {
  const {
    saveButtonProps,
    register,
    formState: { errors },
  } = useForm<IComputer, HttpError, Nullable<IComputer>>();

  return (
    <Create saveButtonProps={saveButtonProps}>
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
          placeholder="Only use alphabets, numbers, '-', '_' and '.'. Whitespaces and other special characters are not allowed."
        />
      </Box>
    </Create>
  );
};
