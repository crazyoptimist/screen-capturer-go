import { Refine, Authenticated, IResourceItem } from "@refinedev/core";
import {
  ThemedLayoutV2,
  ErrorComponent,
  RefineThemes,
  useNotificationProvider,
  RefineSnackbarProvider,
} from "@refinedev/mui";
import CssBaseline from "@mui/material/CssBaseline";
import GlobalStyles from "@mui/material/GlobalStyles";
import { ThemeProvider } from "@mui/material/styles";
import routerProvider, {
  NavigateToResource,
  UnsavedChangesNotifier,
  DocumentTitleHandler,
} from "@refinedev/react-router-v6";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";

import { dataProvider } from "./providers/data-provider";
import { resources } from "./resources";
import {
  ComputerList,
  ComputerCreate,
  ComputerEdit,
} from "../src/pages/computers";

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <ThemeProvider theme={RefineThemes.GreenDark}>
        <CssBaseline />
        <GlobalStyles styles={{ html: { WebkitFontSmoothing: "auto" } }} />
        <RefineSnackbarProvider>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Refine
              dataProvider={dataProvider}
              routerProvider={routerProvider}
              notificationProvider={useNotificationProvider}
              resources={resources}
              options={{
                syncWithLocation: true,
                warnWhenUnsavedChanges: true,
                projectId: "1Z3YzY-VcJzxT-kx0Xaf",
                disableTelemetry: true,
              }}
            >
              <Routes>
                <Route element={<Outlet />}>
                  <Route
                    index
                    element={<NavigateToResource resource="computers" />}
                  />

                  <Route path="/computers">
                    <Route index element={<ComputerList />} />
                    <Route path="create" element={<ComputerCreate />} />
                    <Route path=":id/edit" element={<ComputerEdit />} />
                  </Route>
                </Route>

                <Route
                  element={
                    <Authenticated key="catch-all">
                      <ThemedLayoutV2>
                        <Outlet />
                      </ThemedLayoutV2>
                    </Authenticated>
                  }
                >
                  <Route path="*" element={<ErrorComponent />} />
                </Route>
              </Routes>
              <UnsavedChangesNotifier />
              <DocumentTitleHandler handler={docTitleHandler} />
            </Refine>
          </LocalizationProvider>
        </RefineSnackbarProvider>
      </ThemeProvider>
    </BrowserRouter>
  );
};

export default App;

function docTitleHandler({
  resource,
}: {
  resource?: IResourceItem | undefined;
}) {
  let title = "Screen Capturer Control Panel";
  if (resource) {
    title = `${resource.name} | ${title}`;
  }
  return title;
}
