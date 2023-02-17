/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.voting";

export interface Credibility {
  index: string;
  score: string;
  forX: string;
}

export interface AppliedX {
  X: string;
}

const baseCredibility: object = { index: "", score: "", forX: "" };

export const Credibility = {
  encode(message: Credibility, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.score !== "") {
      writer.uint32(18).string(message.score);
    }
    if (message.forX !== "") {
      writer.uint32(26).string(message.forX);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Credibility {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCredibility } as Credibility;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.score = reader.string();
          break;
        case 3:
          message.forX = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Credibility {
    const message = { ...baseCredibility } as Credibility;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = String(object.score);
    } else {
      message.score = "";
    }
    if (object.forX !== undefined && object.forX !== null) {
      message.forX = String(object.forX);
    } else {
      message.forX = "";
    }
    return message;
  },

  toJSON(message: Credibility): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.score !== undefined && (obj.score = message.score);
    message.forX !== undefined && (obj.forX = message.forX);
    return obj;
  },

  fromPartial(object: DeepPartial<Credibility>): Credibility {
    const message = { ...baseCredibility } as Credibility;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = "";
    }
    if (object.forX !== undefined && object.forX !== null) {
      message.forX = object.forX;
    } else {
      message.forX = "";
    }
    return message;
  },
};

const baseAppliedX: object = { X: "" };

export const AppliedX = {
  encode(message: AppliedX, writer: Writer = Writer.create()): Writer {
    if (message.X !== "") {
      writer.uint32(10).string(message.X);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AppliedX {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAppliedX } as AppliedX;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.X = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppliedX {
    const message = { ...baseAppliedX } as AppliedX;
    if (object.X !== undefined && object.X !== null) {
      message.X = String(object.X);
    } else {
      message.X = "";
    }
    return message;
  },

  toJSON(message: AppliedX): unknown {
    const obj: any = {};
    message.X !== undefined && (obj.X = message.X);
    return obj;
  },

  fromPartial(object: DeepPartial<AppliedX>): AppliedX {
    const message = { ...baseAppliedX } as AppliedX;
    if (object.X !== undefined && object.X !== null) {
      message.X = object.X;
    } else {
      message.X = "";
    }
    return message;
  },
};

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
