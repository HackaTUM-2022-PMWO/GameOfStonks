import vanillaCreate from "zustand/vanilla";
import create from "zustand";
import { StonksServiceClient } from "../services/stonk-client";
import { Err } from "../services/vo-stonks";
import { getClient } from "../services/transport";
import { persist } from "zustand/middleware";
import {
  PlaceOrderCmd,
  StonkInfo,
  StonkName,
  User,
} from "../services/vo-stonks";
import { Routes } from "../router/router";

export type StonksState = {
  roundDuration: number;

  loading: boolean;
  // set after successfull regoster/newUser
  // removed after lobby closed or when 401 returned from any service call
  username?: string;
  currentUser?: User;
  sessionUsers?: User[];
  gameStarted: boolean;
  stonkInfos?: StonkInfo[];
};

const handleAuthError = (err: Err) => {
  if (err.message === "user is not an active user") {
    // do not redirect if already on landing page
    if (window.location.pathname !== "/" && window.location.pathname !== "") {
      window.location.href = "/";
    }
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
      roundDuration: 0,

      username: undefined,
      stockDetails: [],
      gameStarted: false,
      loading: true,

      register: (name, navigate) => {
        return withLoading(client.newUser(name)).then((resp) => {
          set({ username: name, sessionUsers: (resp.ret as any) ?? [] });

          // start SSE handling
          const evtSource = new EventSource("/stream");
          evtSource.onmessage = (evt) => {
            const payload = JSON.parse(evt.data) as {
              reload?: boolean;
              start?: User[];
              finish?: User[];
            };
            // ready to play baby
            if (payload.start) {
              set({
                gameStarted: true,
                currentUser: payload.start.find((u) => u.Name === name),
                sessionUsers: payload.start,
                // roundDuration: payloa
              });
              navigate(Routes.StartStocks);
              console.log("staring game");
            } else if (payload.finish) {
              set({ gameStarted: false });
              navigate(Routes.Result);
              console.log("game over");
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
        let resp = await Promise.all(
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

        // filter out empty infos
        resp = resp.filter((r) => r.ret.Name !== "");

        set({
          stonkInfos: resp.map((resp) => {
            return resp.ret;
          }),
        });
      },

      updateState: () => {
        withLoading(client.getUserInfo()).then(
          ({ ret: user, ret_1: users, ret_2: err }) => {
            if (err != null) {
              handleAuthError(err);
              return;
            }
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
