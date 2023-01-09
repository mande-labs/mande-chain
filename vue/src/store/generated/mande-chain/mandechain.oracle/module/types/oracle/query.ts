/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../oracle/params";
import { NetworkConstantResult } from "../oracle/network_constant";

export const protobufPackage = "mandechain.oracle";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryNetworkConstantRequest {
  request_id: number;
}

export interface QueryNetworkConstantResponse {
  result: NetworkConstantResult | undefined;
}

export interface QueryLastNetworkConstantIdRequest {}

export interface QueryLastNetworkConstantIdResponse {
  request_id: number;
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

const baseQueryNetworkConstantRequest: object = { request_id: 0 };

export const QueryNetworkConstantRequest = {
  encode(
    message: QueryNetworkConstantRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.request_id !== 0) {
      writer.uint32(8).int64(message.request_id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryNetworkConstantRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryNetworkConstantRequest,
    } as QueryNetworkConstantRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.request_id = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryNetworkConstantRequest {
    const message = {
      ...baseQueryNetworkConstantRequest,
    } as QueryNetworkConstantRequest;
    if (object.request_id !== undefined && object.request_id !== null) {
      message.request_id = Number(object.request_id);
    } else {
      message.request_id = 0;
    }
    return message;
  },

  toJSON(message: QueryNetworkConstantRequest): unknown {
    const obj: any = {};
    message.request_id !== undefined && (obj.request_id = message.request_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryNetworkConstantRequest>
  ): QueryNetworkConstantRequest {
    const message = {
      ...baseQueryNetworkConstantRequest,
    } as QueryNetworkConstantRequest;
    if (object.request_id !== undefined && object.request_id !== null) {
      message.request_id = object.request_id;
    } else {
      message.request_id = 0;
    }
    return message;
  },
};

const baseQueryNetworkConstantResponse: object = {};

export const QueryNetworkConstantResponse = {
  encode(
    message: QueryNetworkConstantResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.result !== undefined) {
      NetworkConstantResult.encode(
        message.result,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryNetworkConstantResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryNetworkConstantResponse,
    } as QueryNetworkConstantResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.result = NetworkConstantResult.decode(
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

  fromJSON(object: any): QueryNetworkConstantResponse {
    const message = {
      ...baseQueryNetworkConstantResponse,
    } as QueryNetworkConstantResponse;
    if (object.result !== undefined && object.result !== null) {
      message.result = NetworkConstantResult.fromJSON(object.result);
    } else {
      message.result = undefined;
    }
    return message;
  },

  toJSON(message: QueryNetworkConstantResponse): unknown {
    const obj: any = {};
    message.result !== undefined &&
      (obj.result = message.result
        ? NetworkConstantResult.toJSON(message.result)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryNetworkConstantResponse>
  ): QueryNetworkConstantResponse {
    const message = {
      ...baseQueryNetworkConstantResponse,
    } as QueryNetworkConstantResponse;
    if (object.result !== undefined && object.result !== null) {
      message.result = NetworkConstantResult.fromPartial(object.result);
    } else {
      message.result = undefined;
    }
    return message;
  },
};

const baseQueryLastNetworkConstantIdRequest: object = {};

export const QueryLastNetworkConstantIdRequest = {
  encode(
    _: QueryLastNetworkConstantIdRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastNetworkConstantIdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastNetworkConstantIdRequest,
    } as QueryLastNetworkConstantIdRequest;
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

  fromJSON(_: any): QueryLastNetworkConstantIdRequest {
    const message = {
      ...baseQueryLastNetworkConstantIdRequest,
    } as QueryLastNetworkConstantIdRequest;
    return message;
  },

  toJSON(_: QueryLastNetworkConstantIdRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryLastNetworkConstantIdRequest>
  ): QueryLastNetworkConstantIdRequest {
    const message = {
      ...baseQueryLastNetworkConstantIdRequest,
    } as QueryLastNetworkConstantIdRequest;
    return message;
  },
};

const baseQueryLastNetworkConstantIdResponse: object = { request_id: 0 };

export const QueryLastNetworkConstantIdResponse = {
  encode(
    message: QueryLastNetworkConstantIdResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.request_id !== 0) {
      writer.uint32(8).int64(message.request_id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryLastNetworkConstantIdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryLastNetworkConstantIdResponse,
    } as QueryLastNetworkConstantIdResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.request_id = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryLastNetworkConstantIdResponse {
    const message = {
      ...baseQueryLastNetworkConstantIdResponse,
    } as QueryLastNetworkConstantIdResponse;
    if (object.request_id !== undefined && object.request_id !== null) {
      message.request_id = Number(object.request_id);
    } else {
      message.request_id = 0;
    }
    return message;
  },

  toJSON(message: QueryLastNetworkConstantIdResponse): unknown {
    const obj: any = {};
    message.request_id !== undefined && (obj.request_id = message.request_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryLastNetworkConstantIdResponse>
  ): QueryLastNetworkConstantIdResponse {
    const message = {
      ...baseQueryLastNetworkConstantIdResponse,
    } as QueryLastNetworkConstantIdResponse;
    if (object.request_id !== undefined && object.request_id !== null) {
      message.request_id = object.request_id;
    } else {
      message.request_id = 0;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** NetworkConstantResult defines a rpc handler method for MsgNetworkConstantData. */
  NetworkConstantResult(
    request: QueryNetworkConstantRequest
  ): Promise<QueryNetworkConstantResponse>;
  /** LastNetworkConstantId query the last NetworkConstant result id */
  LastNetworkConstantId(
    request: QueryLastNetworkConstantIdRequest
  ): Promise<QueryLastNetworkConstantIdResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("mandechain.oracle.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  NetworkConstantResult(
    request: QueryNetworkConstantRequest
  ): Promise<QueryNetworkConstantResponse> {
    const data = QueryNetworkConstantRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.oracle.Query",
      "NetworkConstantResult",
      data
    );
    return promise.then((data) =>
      QueryNetworkConstantResponse.decode(new Reader(data))
    );
  }

  LastNetworkConstantId(
    request: QueryLastNetworkConstantIdRequest
  ): Promise<QueryLastNetworkConstantIdResponse> {
    const data = QueryLastNetworkConstantIdRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.oracle.Query",
      "LastNetworkConstantId",
      data
    );
    return promise.then((data) =>
      QueryLastNetworkConstantIdResponse.decode(new Reader(data))
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
