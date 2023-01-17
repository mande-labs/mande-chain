/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../voting/params";
import { VoteBook } from "../voting/vote_book";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { AggregateVotesCasted } from "../voting/aggregate_votes_casted";
import { AggregateVotesReceived } from "../voting/aggregate_votes_received";

export const protobufPackage = "mandechain.voting";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetVoteBookRequest {
  index: string;
}

export interface QueryGetVoteBookResponse {
  voteBook: VoteBook | undefined;
}

export interface QueryAllVoteBookRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVoteBookResponse {
  voteBook: VoteBook[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAggregateVotesCastedRequest {
  index: string;
}

export interface QueryGetAggregateVotesCastedResponse {
  aggregateVotesCasted: AggregateVotesCasted | undefined;
}

export interface QueryAllAggregateVotesCastedRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAggregateVotesCastedResponse {
  aggregateVotesCasted: AggregateVotesCasted[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAggregateVotesReceivedRequest {
  index: string;
}

export interface QueryGetAggregateVotesReceivedResponse {
  aggregateVotesReceived: AggregateVotesReceived | undefined;
}

export interface QueryAllAggregateVotesReceivedRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAggregateVotesReceivedResponse {
  aggregateVotesReceived: AggregateVotesReceived[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetVoteBookRequest: object = { index: "" };

export const QueryGetVoteBookRequest = {
  encode(
    message: QueryGetVoteBookRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVoteBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVoteBookRequest,
    } as QueryGetVoteBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVoteBookRequest {
    const message = {
      ...baseQueryGetVoteBookRequest,
    } as QueryGetVoteBookRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetVoteBookRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVoteBookRequest>
  ): QueryGetVoteBookRequest {
    const message = {
      ...baseQueryGetVoteBookRequest,
    } as QueryGetVoteBookRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetVoteBookResponse: object = {};

export const QueryGetVoteBookResponse = {
  encode(
    message: QueryGetVoteBookResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.voteBook !== undefined) {
      VoteBook.encode(message.voteBook, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetVoteBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetVoteBookResponse,
    } as QueryGetVoteBookResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.voteBook = VoteBook.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetVoteBookResponse {
    const message = {
      ...baseQueryGetVoteBookResponse,
    } as QueryGetVoteBookResponse;
    if (object.voteBook !== undefined && object.voteBook !== null) {
      message.voteBook = VoteBook.fromJSON(object.voteBook);
    } else {
      message.voteBook = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetVoteBookResponse): unknown {
    const obj: any = {};
    message.voteBook !== undefined &&
      (obj.voteBook = message.voteBook
        ? VoteBook.toJSON(message.voteBook)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetVoteBookResponse>
  ): QueryGetVoteBookResponse {
    const message = {
      ...baseQueryGetVoteBookResponse,
    } as QueryGetVoteBookResponse;
    if (object.voteBook !== undefined && object.voteBook !== null) {
      message.voteBook = VoteBook.fromPartial(object.voteBook);
    } else {
      message.voteBook = undefined;
    }
    return message;
  },
};

const baseQueryAllVoteBookRequest: object = {};

export const QueryAllVoteBookRequest = {
  encode(
    message: QueryAllVoteBookRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVoteBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVoteBookRequest,
    } as QueryAllVoteBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVoteBookRequest {
    const message = {
      ...baseQueryAllVoteBookRequest,
    } as QueryAllVoteBookRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVoteBookRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVoteBookRequest>
  ): QueryAllVoteBookRequest {
    const message = {
      ...baseQueryAllVoteBookRequest,
    } as QueryAllVoteBookRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllVoteBookResponse: object = {};

export const QueryAllVoteBookResponse = {
  encode(
    message: QueryAllVoteBookResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.voteBook) {
      VoteBook.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllVoteBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllVoteBookResponse,
    } as QueryAllVoteBookResponse;
    message.voteBook = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.voteBook.push(VoteBook.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllVoteBookResponse {
    const message = {
      ...baseQueryAllVoteBookResponse,
    } as QueryAllVoteBookResponse;
    message.voteBook = [];
    if (object.voteBook !== undefined && object.voteBook !== null) {
      for (const e of object.voteBook) {
        message.voteBook.push(VoteBook.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllVoteBookResponse): unknown {
    const obj: any = {};
    if (message.voteBook) {
      obj.voteBook = message.voteBook.map((e) =>
        e ? VoteBook.toJSON(e) : undefined
      );
    } else {
      obj.voteBook = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllVoteBookResponse>
  ): QueryAllVoteBookResponse {
    const message = {
      ...baseQueryAllVoteBookResponse,
    } as QueryAllVoteBookResponse;
    message.voteBook = [];
    if (object.voteBook !== undefined && object.voteBook !== null) {
      for (const e of object.voteBook) {
        message.voteBook.push(VoteBook.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAggregateVotesCastedRequest: object = { index: "" };

export const QueryGetAggregateVotesCastedRequest = {
  encode(
    message: QueryGetAggregateVotesCastedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAggregateVotesCastedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAggregateVotesCastedRequest,
    } as QueryGetAggregateVotesCastedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAggregateVotesCastedRequest {
    const message = {
      ...baseQueryGetAggregateVotesCastedRequest,
    } as QueryGetAggregateVotesCastedRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetAggregateVotesCastedRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAggregateVotesCastedRequest>
  ): QueryGetAggregateVotesCastedRequest {
    const message = {
      ...baseQueryGetAggregateVotesCastedRequest,
    } as QueryGetAggregateVotesCastedRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetAggregateVotesCastedResponse: object = {};

export const QueryGetAggregateVotesCastedResponse = {
  encode(
    message: QueryGetAggregateVotesCastedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.aggregateVotesCasted !== undefined) {
      AggregateVotesCasted.encode(
        message.aggregateVotesCasted,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAggregateVotesCastedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAggregateVotesCastedResponse,
    } as QueryGetAggregateVotesCastedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aggregateVotesCasted = AggregateVotesCasted.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAggregateVotesCastedResponse {
    const message = {
      ...baseQueryGetAggregateVotesCastedResponse,
    } as QueryGetAggregateVotesCastedResponse;
    if (
      object.aggregateVotesCasted !== undefined &&
      object.aggregateVotesCasted !== null
    ) {
      message.aggregateVotesCasted = AggregateVotesCasted.fromJSON(
        object.aggregateVotesCasted
      );
    } else {
      message.aggregateVotesCasted = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAggregateVotesCastedResponse): unknown {
    const obj: any = {};
    message.aggregateVotesCasted !== undefined &&
      (obj.aggregateVotesCasted = message.aggregateVotesCasted
        ? AggregateVotesCasted.toJSON(message.aggregateVotesCasted)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAggregateVotesCastedResponse>
  ): QueryGetAggregateVotesCastedResponse {
    const message = {
      ...baseQueryGetAggregateVotesCastedResponse,
    } as QueryGetAggregateVotesCastedResponse;
    if (
      object.aggregateVotesCasted !== undefined &&
      object.aggregateVotesCasted !== null
    ) {
      message.aggregateVotesCasted = AggregateVotesCasted.fromPartial(
        object.aggregateVotesCasted
      );
    } else {
      message.aggregateVotesCasted = undefined;
    }
    return message;
  },
};

const baseQueryAllAggregateVotesCastedRequest: object = {};

export const QueryAllAggregateVotesCastedRequest = {
  encode(
    message: QueryAllAggregateVotesCastedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllAggregateVotesCastedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAggregateVotesCastedRequest,
    } as QueryAllAggregateVotesCastedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAggregateVotesCastedRequest {
    const message = {
      ...baseQueryAllAggregateVotesCastedRequest,
    } as QueryAllAggregateVotesCastedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAggregateVotesCastedRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAggregateVotesCastedRequest>
  ): QueryAllAggregateVotesCastedRequest {
    const message = {
      ...baseQueryAllAggregateVotesCastedRequest,
    } as QueryAllAggregateVotesCastedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAggregateVotesCastedResponse: object = {};

export const QueryAllAggregateVotesCastedResponse = {
  encode(
    message: QueryAllAggregateVotesCastedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.aggregateVotesCasted) {
      AggregateVotesCasted.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllAggregateVotesCastedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAggregateVotesCastedResponse,
    } as QueryAllAggregateVotesCastedResponse;
    message.aggregateVotesCasted = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aggregateVotesCasted.push(
            AggregateVotesCasted.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAggregateVotesCastedResponse {
    const message = {
      ...baseQueryAllAggregateVotesCastedResponse,
    } as QueryAllAggregateVotesCastedResponse;
    message.aggregateVotesCasted = [];
    if (
      object.aggregateVotesCasted !== undefined &&
      object.aggregateVotesCasted !== null
    ) {
      for (const e of object.aggregateVotesCasted) {
        message.aggregateVotesCasted.push(AggregateVotesCasted.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAggregateVotesCastedResponse): unknown {
    const obj: any = {};
    if (message.aggregateVotesCasted) {
      obj.aggregateVotesCasted = message.aggregateVotesCasted.map((e) =>
        e ? AggregateVotesCasted.toJSON(e) : undefined
      );
    } else {
      obj.aggregateVotesCasted = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAggregateVotesCastedResponse>
  ): QueryAllAggregateVotesCastedResponse {
    const message = {
      ...baseQueryAllAggregateVotesCastedResponse,
    } as QueryAllAggregateVotesCastedResponse;
    message.aggregateVotesCasted = [];
    if (
      object.aggregateVotesCasted !== undefined &&
      object.aggregateVotesCasted !== null
    ) {
      for (const e of object.aggregateVotesCasted) {
        message.aggregateVotesCasted.push(AggregateVotesCasted.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetAggregateVotesReceivedRequest: object = { index: "" };

export const QueryGetAggregateVotesReceivedRequest = {
  encode(
    message: QueryGetAggregateVotesReceivedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAggregateVotesReceivedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAggregateVotesReceivedRequest,
    } as QueryGetAggregateVotesReceivedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAggregateVotesReceivedRequest {
    const message = {
      ...baseQueryGetAggregateVotesReceivedRequest,
    } as QueryGetAggregateVotesReceivedRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetAggregateVotesReceivedRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAggregateVotesReceivedRequest>
  ): QueryGetAggregateVotesReceivedRequest {
    const message = {
      ...baseQueryGetAggregateVotesReceivedRequest,
    } as QueryGetAggregateVotesReceivedRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetAggregateVotesReceivedResponse: object = {};

export const QueryGetAggregateVotesReceivedResponse = {
  encode(
    message: QueryGetAggregateVotesReceivedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.aggregateVotesReceived !== undefined) {
      AggregateVotesReceived.encode(
        message.aggregateVotesReceived,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAggregateVotesReceivedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAggregateVotesReceivedResponse,
    } as QueryGetAggregateVotesReceivedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aggregateVotesReceived = AggregateVotesReceived.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAggregateVotesReceivedResponse {
    const message = {
      ...baseQueryGetAggregateVotesReceivedResponse,
    } as QueryGetAggregateVotesReceivedResponse;
    if (
      object.aggregateVotesReceived !== undefined &&
      object.aggregateVotesReceived !== null
    ) {
      message.aggregateVotesReceived = AggregateVotesReceived.fromJSON(
        object.aggregateVotesReceived
      );
    } else {
      message.aggregateVotesReceived = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAggregateVotesReceivedResponse): unknown {
    const obj: any = {};
    message.aggregateVotesReceived !== undefined &&
      (obj.aggregateVotesReceived = message.aggregateVotesReceived
        ? AggregateVotesReceived.toJSON(message.aggregateVotesReceived)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAggregateVotesReceivedResponse>
  ): QueryGetAggregateVotesReceivedResponse {
    const message = {
      ...baseQueryGetAggregateVotesReceivedResponse,
    } as QueryGetAggregateVotesReceivedResponse;
    if (
      object.aggregateVotesReceived !== undefined &&
      object.aggregateVotesReceived !== null
    ) {
      message.aggregateVotesReceived = AggregateVotesReceived.fromPartial(
        object.aggregateVotesReceived
      );
    } else {
      message.aggregateVotesReceived = undefined;
    }
    return message;
  },
};

const baseQueryAllAggregateVotesReceivedRequest: object = {};

export const QueryAllAggregateVotesReceivedRequest = {
  encode(
    message: QueryAllAggregateVotesReceivedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllAggregateVotesReceivedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAggregateVotesReceivedRequest,
    } as QueryAllAggregateVotesReceivedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAggregateVotesReceivedRequest {
    const message = {
      ...baseQueryAllAggregateVotesReceivedRequest,
    } as QueryAllAggregateVotesReceivedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAggregateVotesReceivedRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAggregateVotesReceivedRequest>
  ): QueryAllAggregateVotesReceivedRequest {
    const message = {
      ...baseQueryAllAggregateVotesReceivedRequest,
    } as QueryAllAggregateVotesReceivedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAggregateVotesReceivedResponse: object = {};

export const QueryAllAggregateVotesReceivedResponse = {
  encode(
    message: QueryAllAggregateVotesReceivedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.aggregateVotesReceived) {
      AggregateVotesReceived.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllAggregateVotesReceivedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAggregateVotesReceivedResponse,
    } as QueryAllAggregateVotesReceivedResponse;
    message.aggregateVotesReceived = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aggregateVotesReceived.push(
            AggregateVotesReceived.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllAggregateVotesReceivedResponse {
    const message = {
      ...baseQueryAllAggregateVotesReceivedResponse,
    } as QueryAllAggregateVotesReceivedResponse;
    message.aggregateVotesReceived = [];
    if (
      object.aggregateVotesReceived !== undefined &&
      object.aggregateVotesReceived !== null
    ) {
      for (const e of object.aggregateVotesReceived) {
        message.aggregateVotesReceived.push(AggregateVotesReceived.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAggregateVotesReceivedResponse): unknown {
    const obj: any = {};
    if (message.aggregateVotesReceived) {
      obj.aggregateVotesReceived = message.aggregateVotesReceived.map((e) =>
        e ? AggregateVotesReceived.toJSON(e) : undefined
      );
    } else {
      obj.aggregateVotesReceived = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAggregateVotesReceivedResponse>
  ): QueryAllAggregateVotesReceivedResponse {
    const message = {
      ...baseQueryAllAggregateVotesReceivedResponse,
    } as QueryAllAggregateVotesReceivedResponse;
    message.aggregateVotesReceived = [];
    if (
      object.aggregateVotesReceived !== undefined &&
      object.aggregateVotesReceived !== null
    ) {
      for (const e of object.aggregateVotesReceived) {
        message.aggregateVotesReceived.push(
          AggregateVotesReceived.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a VoteBook by index. */
  VoteBook(request: QueryGetVoteBookRequest): Promise<QueryGetVoteBookResponse>;
  /** Queries a list of VoteBook items. */
  VoteBookAll(
    request: QueryAllVoteBookRequest
  ): Promise<QueryAllVoteBookResponse>;
  /** Queries a AggregateVotesCasted by index. */
  AggregateVotesCasted(
    request: QueryGetAggregateVotesCastedRequest
  ): Promise<QueryGetAggregateVotesCastedResponse>;
  /** Queries a list of AggregateVotesCasted items. */
  AggregateVotesCastedAll(
    request: QueryAllAggregateVotesCastedRequest
  ): Promise<QueryAllAggregateVotesCastedResponse>;
  /** Queries a AggregateVotesReceived by index. */
  AggregateVotesReceived(
    request: QueryGetAggregateVotesReceivedRequest
  ): Promise<QueryGetAggregateVotesReceivedResponse>;
  /** Queries a list of AggregateVotesReceived items. */
  AggregateVotesReceivedAll(
    request: QueryAllAggregateVotesReceivedRequest
  ): Promise<QueryAllAggregateVotesReceivedResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("mandechain.voting.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  VoteBook(
    request: QueryGetVoteBookRequest
  ): Promise<QueryGetVoteBookResponse> {
    const data = QueryGetVoteBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "VoteBook",
      data
    );
    return promise.then((data) =>
      QueryGetVoteBookResponse.decode(new Reader(data))
    );
  }

  VoteBookAll(
    request: QueryAllVoteBookRequest
  ): Promise<QueryAllVoteBookResponse> {
    const data = QueryAllVoteBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "VoteBookAll",
      data
    );
    return promise.then((data) =>
      QueryAllVoteBookResponse.decode(new Reader(data))
    );
  }

  AggregateVotesCasted(
    request: QueryGetAggregateVotesCastedRequest
  ): Promise<QueryGetAggregateVotesCastedResponse> {
    const data = QueryGetAggregateVotesCastedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "AggregateVotesCasted",
      data
    );
    return promise.then((data) =>
      QueryGetAggregateVotesCastedResponse.decode(new Reader(data))
    );
  }

  AggregateVotesCastedAll(
    request: QueryAllAggregateVotesCastedRequest
  ): Promise<QueryAllAggregateVotesCastedResponse> {
    const data = QueryAllAggregateVotesCastedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "AggregateVotesCastedAll",
      data
    );
    return promise.then((data) =>
      QueryAllAggregateVotesCastedResponse.decode(new Reader(data))
    );
  }

  AggregateVotesReceived(
    request: QueryGetAggregateVotesReceivedRequest
  ): Promise<QueryGetAggregateVotesReceivedResponse> {
    const data = QueryGetAggregateVotesReceivedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "AggregateVotesReceived",
      data
    );
    return promise.then((data) =>
      QueryGetAggregateVotesReceivedResponse.decode(new Reader(data))
    );
  }

  AggregateVotesReceivedAll(
    request: QueryAllAggregateVotesReceivedRequest
  ): Promise<QueryAllAggregateVotesReceivedResponse> {
    const data = QueryAllAggregateVotesReceivedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "AggregateVotesReceivedAll",
      data
    );
    return promise.then((data) =>
      QueryAllAggregateVotesReceivedResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
