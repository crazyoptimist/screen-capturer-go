import type { DataProvider } from "@refinedev/core";
import { API_URL } from "../config";

const fetcher = async (url: string, options?: RequestInit) =>
  fetch(url, {
    ...options,
    headers: {
      ...options?.headers,
    },
  });

export const dataProvider: DataProvider = {
  getList: async ({ resource, pagination, filters, sorters, meta }) => {
    const params = new URLSearchParams();

    if (pagination && pagination.current && pagination.pageSize) {
      const offset = (pagination.current - 1) * pagination.pageSize;
      const limit = pagination.pageSize;
      params.append("_offset", offset.toString());
      params.append("_limit", limit.toString());
    }

    if (sorters && sorters.length > 0) {
      // MUI free version only supports single field sort,
      // even though below code is for multi-sort as well
      params.append("_sort", sorters.map((sorter) => sorter.field).join(","));
      params.append("_order", sorters.map((sorter) => sorter.order).join(","));
    }

    if (filters && filters.length > 0) {
      // MUI free version only supports single field filter, just like sorting
      // We only support "eq" filter for now
      filters.forEach((filter) => {
        if ("field" in filter && filter.operator === "eq") {
          params.append(filter.field, filter.value);
        }
      });
    }

    const response = await fetcher(
      `${API_URL}/${resource}?${params.toString()}`
    );

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    const total = Number(response.headers.get("X-Total-Count"));

    return {
      data,
      total,
    };
  },
  getMany: async ({ resource, ids, meta }) => {
    const params = new URLSearchParams();

    if (ids) {
      ids.forEach((id) => params.append("id", id as string));
    }

    const response = await fetcher(
      `${API_URL}/${resource}?${params.toString()}`
    );

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    return { data };
  },
  getOne: async ({ resource, id, meta }) => {
    const response = await fetcher(`${API_URL}/${resource}/${id}`);

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    return { data };
  },
  create: async ({ resource, variables }) => {
    const response = await fetcher(`${API_URL}/${resource}`, {
      method: "POST",
      body: JSON.stringify(variables),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    return { data };
  },
  update: async ({ resource, id, variables }) => {
    // Workaround for setting patch
    const patchUrl = id
      ? `${API_URL}/${resource}/${id}`
      : `${API_URL}/${resource}`;

    const response = await fetcher(patchUrl, {
      method: "PATCH",
      body: JSON.stringify(variables),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    return { data };
  },
  getApiUrl: () => API_URL,
  deleteOne: async ({ resource, id }) => {
    const response = await fetcher(`${API_URL}/${resource}/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    return { data };
  },
  custom: async ({
    url,
    method,
    payload,
    // filters,
    // sorters,
    // query,
    // headers,
    // meta,
  }) => {
    const response = await fetcher(`${url}`, {
      method,
      body: JSON.stringify(payload),
    });

    if (response.status < 200 || response.status > 299) throw response;

    const data = await response.json();

    return { data };
  },
};
