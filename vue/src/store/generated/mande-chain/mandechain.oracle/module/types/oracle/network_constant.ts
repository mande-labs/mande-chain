/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.oracle";

export interface NetworkConstantCallData {
  repeat: number;
}

export interface NetworkConstantResult {
  response: string;
}

const baseNetworkConstantCallData: object = { repeat: 0 };

export const NetworkConstantCallData = {
  encode(
    message: NetworkConstantCallData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.repeat !== 0) {
      writer.uint32(8).uint64(message.repeat);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NetworkConstantCallData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseNetworkConstantCallData,
    } as NetworkConstantCallData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.repeat = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NetworkConstantCallData {
    const message = {
      ...baseNetworkConstantCallData,
    } as NetworkConstantCallData;
    if (object.repeat !== undefined && object.repeat !== null) {
      message.repeat = Number(object.repeat);
    } else {
      message.repeat = 0;
    }
    return message;
  },

  toJSON(message: NetworkConstantCallData): unknown {
    const obj: any = {};
    message.repeat !== undefined && (obj.repeat = message.repeat);
    return obj;
  },

  fromPartial(
    object: DeepPartial<NetworkConstantCallData>
  ): NetworkConstantCallData {
    const message = {
      ...baseNetworkConstantCallData,
    } as NetworkConstantCallData;
    if (object.repeat !== undefined && object.repeat !== null) {
      message.repeat = object.repeat;
    } else {
      message.repeat = 0;
    }
    return message;
  },
};

const baseNetworkConstantResult: object = { response: "" };

export const NetworkConstantResult = {
  encode(
    message: NetworkConstantResult,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.response !== "") {
      writer.uint32(10).string(message.response);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NetworkConstantResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNetworkConstantResult } as NetworkConstantResult;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.response = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NetworkConstantResult {
    const message = { ...baseNetworkConstantResult } as NetworkConstantResult;
    if (object.response !== undefined && object.response !== null) {
      message.response = String(object.response);
    } else {
      message.response = "";
    }
    return message;
  },

  toJSON(message: NetworkConstantResult): unknown {
    const obj: any = {};
    message.response !== undefined && (obj.response = message.response);
    return obj;
  },

  fromPartial(
    object: DeepPartial<NetworkConstantResult>
  ): NetworkConstantResult {
    const message = { ...baseNetworkConstantResult } as NetworkConstantResult;
    if (object.response !== undefined && object.response !== null) {
      message.response = object.response;
    } else {
      message.response = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
