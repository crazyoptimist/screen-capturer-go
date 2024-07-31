export type Nullable<T> = {
  [P in keyof T]: T[P] | null;
};

export interface IComputer {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  ipAddress: string;
  isActive?: boolean;
}
