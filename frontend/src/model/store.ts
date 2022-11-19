import vanillaCreate from "zustand/vanilla";
import create from "zustand";
import { StonksClient } from "../services/client";
import { getClient } from "../services/transport";

export const vanillaStore = vanillaCreate((set, get) => {
  const client = getClient(StonksClient);

  client.hello("test").then((resp) => console.log(resp));

  return {};
});

export const useStonkState = create(vanillaStore);
