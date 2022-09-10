/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../voting/params";
import { VoteBook } from "../voting/vote_book";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { AggregateVoteCount } from "../voting/aggregate_vote_count";

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

export interface QueryGetAggregateVoteCountRequest {
  index: string;
}

export interface QueryGetAggregateVoteCountResponse {
  aggregateVoteCount: AggregateVoteCount | undefined;
}

export interface QueryAllAggregateVoteCountRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAggregateVoteCountResponse {
  aggregateVoteCount: AggregateVoteCount[];
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

const baseQueryGetAggregateVoteCountRequest: object = { index: "" };

export const QueryGetAggregateVoteCountRequest = {
  encode(
    message: QueryGetAggregateVoteCountRequest,
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
  ): QueryGetAggregateVoteCountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAggregateVoteCountRequest,
    } as QueryGetAggregateVoteCountRequest;
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

  fromJSON(object: any): QueryGetAggregateVoteCountRequest {
    const message = {
      ...baseQueryGetAggregateVoteCountRequest,
    } as QueryGetAggregateVoteCountRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetAggregateVoteCountRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAggregateVoteCountRequest>
  ): QueryGetAggregateVoteCountRequest {
    const message = {
      ...baseQueryGetAggregateVoteCountRequest,
    } as QueryGetAggregateVoteCountRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetAggregateVoteCountResponse: object = {};

export const QueryGetAggregateVoteCountResponse = {
  encode(
    message: QueryGetAggregateVoteCountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.aggregateVoteCount !== undefined) {
      AggregateVoteCount.encode(
        message.aggregateVoteCount,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAggregateVoteCountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAggregateVoteCountResponse,
    } as QueryGetAggregateVoteCountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aggregateVoteCount = AggregateVoteCount.decode(
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

  fromJSON(object: any): QueryGetAggregateVoteCountResponse {
    const message = {
      ...baseQueryGetAggregateVoteCountResponse,
    } as QueryGetAggregateVoteCountResponse;
    if (
      object.aggregateVoteCount !== undefined &&
      object.aggregateVoteCount !== null
    ) {
      message.aggregateVoteCount = AggregateVoteCount.fromJSON(
        object.aggregateVoteCount
      );
    } else {
      message.aggregateVoteCount = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAggregateVoteCountResponse): unknown {
    const obj: any = {};
    message.aggregateVoteCount !== undefined &&
      (obj.aggregateVoteCount = message.aggregateVoteCount
        ? AggregateVoteCount.toJSON(message.aggregateVoteCount)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAggregateVoteCountResponse>
  ): QueryGetAggregateVoteCountResponse {
    const message = {
      ...baseQueryGetAggregateVoteCountResponse,
    } as QueryGetAggregateVoteCountResponse;
    if (
      object.aggregateVoteCount !== undefined &&
      object.aggregateVoteCount !== null
    ) {
      message.aggregateVoteCount = AggregateVoteCount.fromPartial(
        object.aggregateVoteCount
      );
    } else {
      message.aggregateVoteCount = undefined;
    }
    return message;
  },
};

const baseQueryAllAggregateVoteCountRequest: object = {};

export const QueryAllAggregateVoteCountRequest = {
  encode(
    message: QueryAllAggregateVoteCountRequest,
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
  ): QueryAllAggregateVoteCountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAggregateVoteCountRequest,
    } as QueryAllAggregateVoteCountRequest;
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

  fromJSON(object: any): QueryAllAggregateVoteCountRequest {
    const message = {
      ...baseQueryAllAggregateVoteCountRequest,
    } as QueryAllAggregateVoteCountRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAggregateVoteCountRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAggregateVoteCountRequest>
  ): QueryAllAggregateVoteCountRequest {
    const message = {
      ...baseQueryAllAggregateVoteCountRequest,
    } as QueryAllAggregateVoteCountRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAggregateVoteCountResponse: object = {};

export const QueryAllAggregateVoteCountResponse = {
  encode(
    message: QueryAllAggregateVoteCountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.aggregateVoteCount) {
      AggregateVoteCount.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllAggregateVoteCountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAggregateVoteCountResponse,
    } as QueryAllAggregateVoteCountResponse;
    message.aggregateVoteCount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aggregateVoteCount.push(
            AggregateVoteCount.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllAggregateVoteCountResponse {
    const message = {
      ...baseQueryAllAggregateVoteCountResponse,
    } as QueryAllAggregateVoteCountResponse;
    message.aggregateVoteCount = [];
    if (
      object.aggregateVoteCount !== undefined &&
      object.aggregateVoteCount !== null
    ) {
      for (const e of object.aggregateVoteCount) {
        message.aggregateVoteCount.push(AggregateVoteCount.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAggregateVoteCountResponse): unknown {
    const obj: any = {};
    if (message.aggregateVoteCount) {
      obj.aggregateVoteCount = message.aggregateVoteCount.map((e) =>
        e ? AggregateVoteCount.toJSON(e) : undefined
      );
    } else {
      obj.aggregateVoteCount = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAggregateVoteCountResponse>
  ): QueryAllAggregateVoteCountResponse {
    const message = {
      ...baseQueryAllAggregateVoteCountResponse,
    } as QueryAllAggregateVoteCountResponse;
    message.aggregateVoteCount = [];
    if (
      object.aggregateVoteCount !== undefined &&
      object.aggregateVoteCount !== null
    ) {
      for (const e of object.aggregateVoteCount) {
        message.aggregateVoteCount.push(AggregateVoteCount.fromPartial(e));
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
  /** Queries a AggregateVoteCount by index. */
  AggregateVoteCount(
    request: QueryGetAggregateVoteCountRequest
  ): Promise<QueryGetAggregateVoteCountResponse>;
  /** Queries a list of AggregateVoteCount items. */
  AggregateVoteCountAll(
    request: QueryAllAggregateVoteCountRequest
  ): Promise<QueryAllAggregateVoteCountResponse>;
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

  AggregateVoteCount(
    request: QueryGetAggregateVoteCountRequest
  ): Promise<QueryGetAggregateVoteCountResponse> {
    const data = QueryGetAggregateVoteCountRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "AggregateVoteCount",
      data
    );
    return promise.then((data) =>
      QueryGetAggregateVoteCountResponse.decode(new Reader(data))
    );
  }

  AggregateVoteCountAll(
    request: QueryAllAggregateVoteCountRequest
  ): Promise<QueryAllAggregateVoteCountResponse> {
    const data = QueryAllAggregateVoteCountRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.voting.Query",
      "AggregateVoteCountAll",
      data
    );
    return promise.then((data) =>
      QueryAllAggregateVoteCountResponse.decode(new Reader(data))
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
