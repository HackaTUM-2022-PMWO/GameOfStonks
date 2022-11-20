import vanillaCreate from "zustand/vanilla";
import create from "zustand";
import { StonksServiceClient } from "../services/stonk-client";
import {} from "../services/vo-stonks";
import { getClient } from "../services/transport";
import {
  OrderType,
  PlaceOrderCmd,
  StonkInfo,
  StonkName,
  User,
} from "../services/vo-stonks";
import userEvent from "@testing-library/user-event";
import { Routes } from "../router/router";

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
  register: (
    name: string,
    navigate: (url: string) => void
  ) => ReturnType<StonksServiceClient["newUser"]>;

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

      register: (name, navigate) => {
        return withLoading(client.newUser(name)).then((resp) => {
          set({ username: name, sessionUsers: (resp.ret as any) ?? [] });

          const evtSource = new EventSource("/stream");
          evtSource.onmessage = (evt) => {
            const payload = JSON.parse(evt.data);
            console.log("msg", payload);

            // ready to play baby
            if (payload.start) {
              navigate(Routes.StartStocks);
            }
          };
          evtSource.onopen = (evt) => {
            console.log("channel opened");
          };
          evtSource.onerror = (evt) => console.error(evt);

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

      updateState: () => {
        withLoading(client.getUserInfo()).then(
          ({ ret: user, ret_1: users, ret_2: err }) => {
            // if (err == nil)
          }
        );
      },
    };
  }
);

export const useStonkState = create(vanillaStore);
