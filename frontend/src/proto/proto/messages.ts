/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export interface Stonk {
  name: string;
}

export interface StonkRequest {
  name: string;
}

function createBaseStonk(): Stonk {
  return { name: "" };
}

export const Stonk = {
  encode(message: Stonk, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Stonk {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStonk();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Stonk {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: Stonk): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Stonk>, I>>(object: I): Stonk {
    const message = createBaseStonk();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseStonkRequest(): StonkRequest {
  return { name: "" };
}

export const StonkRequest = {
  encode(message: StonkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StonkRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStonkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StonkRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: StonkRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StonkRequest>, I>>(object: I): StonkRequest {
    const message = createBaseStonkRequest();
    message.name = object.name ?? "";
    return message;
  },
};

export interface StonkServer {
  GetStonkByName(request: StonkRequest): Promise<Stonk>;
}

export class StonkServerClientImpl implements StonkServer {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "StonkServer";
    this.rpc = rpc;
    this.GetStonkByName = this.GetStonkByName.bind(this);
  }
  GetStonkByName(request: StonkRequest): Promise<Stonk> {
    const data = StonkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetStonkByName", data);
    return promise.then((data) => Stonk.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
