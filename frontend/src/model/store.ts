import vanillaCreate from "zustand/vanilla";
import create from "zustand";
import { StonksServiceClient } from "../services/stonk-client";
import { Err, Match } from "../services/vo-stonks";
import { getClient } from "../services/transport";
import {
  PlaceOrderCmd,
  StonkInfo,
  StonkName,
  User,
} from "../services/vo-stonks";
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

const handleAuthError = (err: Err) => {
  if (err.message == "user is not an active user") {
    window.location.href = "/";
  }
};

export type StonksModifiers = {
  register: (
    name: string,
    navigate: (url: string) => void
  ) => ReturnType<StonksServiceClient["newUser"]>;

  updateState: () => void;

  getStonkInfo: (
    stonkName: StonkName
  ) => ReturnType<StonksServiceClient["getStonkInfo"]>;

  getStonksHistory: () => Promise<void>;

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

    const withHandleAuthError = <T, E extends Err>(
      promise: Promise<{ ret: T; ret_1: E | null }>
    ) => {
      return promise.then((result) => {
        if (result.ret_1) {
          handleAuthError(result.ret_1);
        }
        return result;
      });
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
            const payload = JSON.parse(evt.data) as {
              reload?: boolean;
              start?: User[];
              finish?: User[];
            };
            console.log("msg", payload);

            // ready to play baby
            if (payload.start) {
              set({ gameStarted: true });
              navigate(Routes.StartStocks);
            } else if (payload.finish) {
              set({ gameStarted: false });
              navigate(Routes.Result);
            } else if (payload.reload) {
              get().updateState();
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
        return withLoading(withHandleAuthError(client.getStonkInfo(stonk)));
      },

      placeOrder: (cmd: PlaceOrderCmd) => {
        return withLoading(client.placeOrder({ ...cmd })).then((resp) => {
          return resp;
        });
      },

      getStonksHistory: async () => {
        const resp = await Promise.all(
          Object.values(StonkName).map((stonkName) =>
            client.getStonkInfo(stonkName)
          )
        );

        for (let r of resp) {
          if (r.ret_1 != null) {
            handleAuthError(r.ret_1);
            return;
          }
        }

        set({
          stonkInfos: resp.map((resp) => {
            return resp.ret;
          }),
        });
      },

      updateState: () => {
        withLoading(client.getUserInfo()).then(
          ({ ret: user, ret_1: users, ret_2: err }) => {
            set({
              sessionUsers:
                (users?.filter((user) => user != null) as any) ?? [],
            });
          }
        );
      },
    };
  }
);

export const useStonkState = create(vanillaStore);
