/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.ssi";

/** GenesisState defines the ssi module's genesis state. */
export interface GenesisState {
  chain_namespace: string;
}

const baseGenesisState: object = { chain_namespace: "" };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.chain_namespace !== "") {
      writer.uint32(10).string(message.chain_namespace);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chain_namespace = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (
      object.chain_namespace !== undefined &&
      object.chain_namespace !== null
    ) {
      message.chain_namespace = String(object.chain_namespace);
    } else {
      message.chain_namespace = "";
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.chain_namespace !== undefined &&
      (obj.chain_namespace = message.chain_namespace);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (
      object.chain_namespace !== undefined &&
      object.chain_namespace !== null
    ) {
      message.chain_namespace = object.chain_namespace;
    } else {
      message.chain_namespace = "";
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
