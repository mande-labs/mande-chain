/* eslint-disable */
import { Params } from "../voting/params";
import { VoteBook } from "../voting/vote_book";
import { AggregateVotesCasted } from "../voting/aggregate_votes_casted";
import { AggregateVotesReceived } from "../voting/aggregate_votes_received";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mandechain.voting";

/** GenesisState defines the voting module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  voteBookList: VoteBook[];
  aggregateVotesCastedList: AggregateVotesCasted[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  aggregateVotesReceivedList: AggregateVotesReceived[];
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
    for (const v of message.aggregateVotesCastedList) {
      AggregateVotesCasted.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.aggregateVotesReceivedList) {
      AggregateVotesReceived.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.voteBookList = [];
    message.aggregateVotesCastedList = [];
    message.aggregateVotesReceivedList = [];
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
          message.aggregateVotesCastedList.push(
            AggregateVotesCasted.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.aggregateVotesReceivedList.push(
            AggregateVotesReceived.decode(reader, reader.uint32())
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
    message.aggregateVotesCastedList = [];
    message.aggregateVotesReceivedList = [];
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
      object.aggregateVotesCastedList !== undefined &&
      object.aggregateVotesCastedList !== null
    ) {
      for (const e of object.aggregateVotesCastedList) {
        message.aggregateVotesCastedList.push(AggregateVotesCasted.fromJSON(e));
      }
    }
    if (
      object.aggregateVotesReceivedList !== undefined &&
      object.aggregateVotesReceivedList !== null
    ) {
      for (const e of object.aggregateVotesReceivedList) {
        message.aggregateVotesReceivedList.push(
          AggregateVotesReceived.fromJSON(e)
        );
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
    if (message.aggregateVotesCastedList) {
      obj.aggregateVotesCastedList = message.aggregateVotesCastedList.map((e) =>
        e ? AggregateVotesCasted.toJSON(e) : undefined
      );
    } else {
      obj.aggregateVotesCastedList = [];
    }
    if (message.aggregateVotesReceivedList) {
      obj.aggregateVotesReceivedList = message.aggregateVotesReceivedList.map(
        (e) => (e ? AggregateVotesReceived.toJSON(e) : undefined)
      );
    } else {
      obj.aggregateVotesReceivedList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.voteBookList = [];
    message.aggregateVotesCastedList = [];
    message.aggregateVotesReceivedList = [];
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
      object.aggregateVotesCastedList !== undefined &&
      object.aggregateVotesCastedList !== null
    ) {
      for (const e of object.aggregateVotesCastedList) {
        message.aggregateVotesCastedList.push(
          AggregateVotesCasted.fromPartial(e)
        );
      }
    }
    if (
      object.aggregateVotesReceivedList !== undefined &&
      object.aggregateVotesReceivedList !== null
    ) {
      for (const e of object.aggregateVotesReceivedList) {
        message.aggregateVotesReceivedList.push(
          AggregateVotesReceived.fromPartial(e)
        );
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
