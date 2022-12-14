/* eslint-disable */
import { Params } from "../voting/params";
import { VoteBook } from "../voting/vote_book";
import { AggregateVoteCount } from "../voting/aggregate_vote_count";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.voting";

/** GenesisState defines the voting module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  voteBookList: VoteBook[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  aggregateVoteCountList: AggregateVoteCount[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.voteBookList) {
      VoteBook.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.aggregateVoteCountList) {
      AggregateVoteCount.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.voteBookList = [];
    message.aggregateVoteCountList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.voteBookList.push(VoteBook.decode(reader, reader.uint32()));
          break;
        case 3:
          message.aggregateVoteCountList.push(
            AggregateVoteCount.decode(reader, reader.uint32())
          );
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
    message.voteBookList = [];
    message.aggregateVoteCountList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.voteBookList !== undefined && object.voteBookList !== null) {
      for (const e of object.voteBookList) {
        message.voteBookList.push(VoteBook.fromJSON(e));
      }
    }
    if (
      object.aggregateVoteCountList !== undefined &&
      object.aggregateVoteCountList !== null
    ) {
      for (const e of object.aggregateVoteCountList) {
        message.aggregateVoteCountList.push(AggregateVoteCount.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.voteBookList) {
      obj.voteBookList = message.voteBookList.map((e) =>
        e ? VoteBook.toJSON(e) : undefined
      );
    } else {
      obj.voteBookList = [];
    }
    if (message.aggregateVoteCountList) {
      obj.aggregateVoteCountList = message.aggregateVoteCountList.map((e) =>
        e ? AggregateVoteCount.toJSON(e) : undefined
      );
    } else {
      obj.aggregateVoteCountList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.voteBookList = [];
    message.aggregateVoteCountList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.voteBookList !== undefined && object.voteBookList !== null) {
      for (const e of object.voteBookList) {
        message.voteBookList.push(VoteBook.fromPartial(e));
      }
    }
    if (
      object.aggregateVoteCountList !== undefined &&
      object.aggregateVoteCountList !== null
    ) {
      for (const e of object.aggregateVoteCountList) {
        message.aggregateVoteCountList.push(AggregateVoteCount.fromPartial(e));
      }
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
