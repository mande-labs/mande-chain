/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.voting";

export interface AggregateVotesCasted {
  index: string;
  positive: number;
  negative: number;
}

const baseAggregateVotesCasted: object = {
  index: "",
  positive: 0,
  negative: 0,
};

export const AggregateVotesCasted = {
  encode(
    message: AggregateVotesCasted,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.positive !== 0) {
      writer.uint32(16).uint64(message.positive);
    }
    if (message.negative !== 0) {
      writer.uint32(24).uint64(message.negative);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AggregateVotesCasted {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAggregateVotesCasted } as AggregateVotesCasted;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.positive = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.negative = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AggregateVotesCasted {
    const message = { ...baseAggregateVotesCasted } as AggregateVotesCasted;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
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

  toJSON(message: AggregateVotesCasted): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.positive !== undefined && (obj.positive = message.positive);
    message.negative !== undefined && (obj.negative = message.negative);
    return obj;
  },

  fromPartial(object: DeepPartial<AggregateVotesCasted>): AggregateVotesCasted {
    const message = { ...baseAggregateVotesCasted } as AggregateVotesCasted;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
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
