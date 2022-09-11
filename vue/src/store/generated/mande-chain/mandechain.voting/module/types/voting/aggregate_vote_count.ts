/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.voting";

export interface AggregateVoteCount {
  index: string;
  voter: string;
  casted: number;
  positiveReceived: number;
  negativeReceived: number;
}

const baseAggregateVoteCount: object = {
  index: "",
  voter: "",
  casted: 0,
  positiveReceived: 0,
  negativeReceived: 0,
};

export const AggregateVoteCount = {
  encode(
    message: AggregateVoteCount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.voter !== "") {
      writer.uint32(18).string(message.voter);
    }
    if (message.casted !== 0) {
      writer.uint32(24).uint64(message.casted);
    }
    if (message.positiveReceived !== 0) {
      writer.uint32(32).uint64(message.positiveReceived);
    }
    if (message.negativeReceived !== 0) {
      writer.uint32(40).uint64(message.negativeReceived);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AggregateVoteCount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAggregateVoteCount } as AggregateVoteCount;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.voter = reader.string();
          break;
        case 3:
          message.casted = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.positiveReceived = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.negativeReceived = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AggregateVoteCount {
    const message = { ...baseAggregateVoteCount } as AggregateVoteCount;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    if (object.casted !== undefined && object.casted !== null) {
      message.casted = Number(object.casted);
    } else {
      message.casted = 0;
    }
    if (
      object.positiveReceived !== undefined &&
      object.positiveReceived !== null
    ) {
      message.positiveReceived = Number(object.positiveReceived);
    } else {
      message.positiveReceived = 0;
    }
    if (
      object.negativeReceived !== undefined &&
      object.negativeReceived !== null
    ) {
      message.negativeReceived = Number(object.negativeReceived);
    } else {
      message.negativeReceived = 0;
    }
    return message;
  },

  toJSON(message: AggregateVoteCount): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.voter !== undefined && (obj.voter = message.voter);
    message.casted !== undefined && (obj.casted = message.casted);
    message.positiveReceived !== undefined &&
      (obj.positiveReceived = message.positiveReceived);
    message.negativeReceived !== undefined &&
      (obj.negativeReceived = message.negativeReceived);
    return obj;
  },

  fromPartial(object: DeepPartial<AggregateVoteCount>): AggregateVoteCount {
    const message = { ...baseAggregateVoteCount } as AggregateVoteCount;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
    }
    if (object.casted !== undefined && object.casted !== null) {
      message.casted = object.casted;
    } else {
      message.casted = 0;
    }
    if (
      object.positiveReceived !== undefined &&
      object.positiveReceived !== null
    ) {
      message.positiveReceived = object.positiveReceived;
    } else {
      message.positiveReceived = 0;
    }
    if (
      object.negativeReceived !== undefined &&
      object.negativeReceived !== null
    ) {
      message.negativeReceived = object.negativeReceived;
    } else {
      message.negativeReceived = 0;
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
