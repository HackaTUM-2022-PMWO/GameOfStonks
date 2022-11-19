import vanillaCreate from "zustand/vanilla";
import create from "zustand";
import { StonksServiceClient } from "../services/stonk-client";
import { getClient } from "../services/transport";
import { StonkInfo, User } from "../services/vo-stonks";

export type StonksState = {
  loading: boolean;
  // set after successfull regoster/newUser
  // removed after lobby closed or when 401 returned from any service call
  username?: string;
  sessionUsers?: User[];

  stonkInfos?: StonkInfo[];
};

export type StonksModifiers = {
  register: (name: string) => ReturnType<StonksServiceClient["newUser"]>;

  getStonkInfo: (
    stonkName: string
  ) => ReturnType<StonksServiceClient["getStonkInfo"]>;
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
      loading: false,

      register: (name: string) => {
        return withLoading(client.newUser("name")).then((resp) => {
          set({ username: name });

          // TODO: store incomming users in lobby
          // TODO: start SSE stream here
          return resp;
        });
        // FIXME: handle server error
      },

      getStonkInfo: (stonk: string) => {
        return withLoading(client.getStonkInfo(stonk));
      },
    };
  }
);

export const useStonkState = create(vanillaStore);
