import { ResourceProps } from "@refinedev/core";

export const resources: ResourceProps[] = [
  {
    name: "computers",
    list: "/computers",
    create: "/computers/create",
    edit: "/computers/:id/edit",
  },
];
