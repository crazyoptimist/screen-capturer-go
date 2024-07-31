import React from "react";
import { EditButton, DeleteButton, List, useDataGrid } from "@refinedev/mui";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import dayjs from "dayjs";

import { IComputer } from "../../interfaces";

export const ComputerList: React.FC = () => {
  const { dataGridProps } = useDataGrid<IComputer>({
    errorNotification: (error, _values, _resource) => {
      return {
        message: `Error (status code: ${error?.status})`,
        description: `${error?.statusText}`,
        type: "error",
      };
    },
  });

  const columns = React.useMemo<GridColDef<IComputer>[]>(
    () => [
      {
        field: "name",
        filterable: false,
        sortable: false,
        headerName: "Name",
        minWidth: 100,
        flex: 1,
      },
      {
        field: "isActive",
        filterable: false,
        sortable: false,
        headerName: "Is Active",
        minWidth: 100,
        flex: 1,
      },
      {
        field: "updatedAt",
        filterable: false,
        sortable: false,
        headerName: "Status Checked At",
        renderCell: function render({ row }) {
          return dayjs(row.updatedAt).format("lll");
        },
        minWidth: 200,
        flex: 2,
      },
      {
        field: "actions",
        filterable: false,
        sortable: false,
        headerName: "Actions",
        renderCell: function render({ row }) {
          return (
            <>
              <EditButton hideText recordItemId={row.id} />
              <DeleteButton hideText recordItemId={row.id} />
            </>
          );
        },
        align: "center",
        headerAlign: "center",
        minWidth: 100,
        flex: 1,
      },
    ],
    []
  );

  return (
    <List>
      <DataGrid {...dataGridProps} columns={columns} autoHeight />
    </List>
  );
};
