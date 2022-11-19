import vanillaCreate from "zustand/vanilla";
import create from "zustand";
import { StonksServiceClient } from "../services/stonk-client";
import { getClient } from "../services/transport";
import {
  OrderType,
  PlaceOrderCmd,
  StonkInfo,
  StonkName,
  User,
} from "../services/vo-stonks";

export type StonksState = {
  loading: boolean;
  // set after successfull regoster/newUser
  // removed after lobby closed or when 401 returned from any service call
  username?: string;
  sessionUsers?: User[];
  gameStarted: boolean;
  stonkInfos?: StonkInfo[];
};

export type StonksModifiers = {
  register: (name: string) => ReturnType<StonksServiceClient["newUser"]>;

  getStonkInfo: (
    stonkName: StonkName
  ) => ReturnType<StonksServiceClient["getStonkInfo"]>;

  placeOrder: (
    cmd: PlaceOrderCmd
  ) => ReturnType<StonksServiceClient["placeOrder"]>;
};

export const vanillaStore = vanillaCreate<StonksState & StonksModifiers>(
  (set, get) => {
    const client = getClient(StonksServiceClient);

    const withLoading = <T>(promise: Promise<T>) => {
      set({ loading: true });
      promise.finally(() => set({ loading: false }));
      return promise;
    };

    return {
      username: undefined,
      stockDetails: [],
      gameStarted: false,
      loading: false,

      register: (name: string) => {
        return withLoading(client.newUser(name)).then((resp) => {
          set({ username: name, sessionUsers: resp.ret ?? [] });
          // TODO: start SSE stream here
          return resp;
        });
        // FIXME: handle server error
      },

      getStonkInfo: (stonk: StonkName) => {
        return withLoading(client.getStonkInfo(stonk));
      },

      placeOrder: (cmd: PlaceOrderCmd) => {
        return withLoading(client.placeOrder({ ...cmd })).then((resp) => {
          return resp;
        });
      },
    };
  }
);

export const useStonkState = create(vanillaStore);
