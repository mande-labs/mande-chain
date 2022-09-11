/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.voting";

export interface VoteBook {
  index: string;
  caster: string;
  receiver: string;
  positive: number;
  negative: number;
}

const baseVoteBook: object = {
  index: "",
  caster: "",
  receiver: "",
  positive: 0,
  negative: 0,
};

export const VoteBook = {
  encode(message: VoteBook, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.caster !== "") {
      writer.uint32(18).string(message.caster);
    }
    if (message.receiver !== "") {
      writer.uint32(26).string(message.receiver);
    }
    if (message.positive !== 0) {
      writer.uint32(32).uint64(message.positive);
    }
    if (message.negative !== 0) {
      writer.uint32(40).uint64(message.negative);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VoteBook {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVoteBook } as VoteBook;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.caster = reader.string();
          break;
        case 3:
          message.receiver = reader.string();
          break;
        case 4:
          message.positive = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.negative = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VoteBook {
    const message = { ...baseVoteBook } as VoteBook;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.caster !== undefined && object.caster !== null) {
      message.caster = String(object.caster);
    } else {
      message.caster = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    if (object.positive !== undefined && object.positive !== null) {
      message.positive = Number(object.positive);
    } else {
      message.positive = 0;
    }
    if (object.negative !== undefined && object.negative !== null) {
      message.negative = Number(object.negative);
    } else {
      message.negative = 0;
    }
    return message;
  },

  toJSON(message: VoteBook): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.caster !== undefined && (obj.caster = message.caster);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.positive !== undefined && (obj.positive = message.positive);
    message.negative !== undefined && (obj.negative = message.negative);
    return obj;
  },

  fromPartial(object: DeepPartial<VoteBook>): VoteBook {
    const message = { ...baseVoteBook } as VoteBook;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.caster !== undefined && object.caster !== null) {
      message.caster = object.caster;
    } else {
      message.caster = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    if (object.positive !== undefined && object.positive !== null) {
      message.positive = object.positive;
    } else {
      message.positive = 0;
    }
    if (object.negative !== undefined && object.negative !== null) {
      message.negative = object.negative;
    } else {
      message.negative = 0;
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
