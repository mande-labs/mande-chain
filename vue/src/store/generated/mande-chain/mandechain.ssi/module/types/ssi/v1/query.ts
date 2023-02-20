/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Schema } from "../../ssi/v1/schema";
import { PageRequest } from "../../cosmos/base/query/v1beta1/pagination";
import { Credential } from "../../ssi/v1/credential";
import { Did, Metadata } from "../../ssi/v1/did";

export const protobufPackage = "mandechain.ssi";

export interface QuerySchemaRequest {
  schemaId: string;
}

export interface QuerySchemaResponse {
  schema: Schema[];
}

export interface QuerySchemasRequest {
  pagination: PageRequest | undefined;
}

export interface QuerySchemasResponse {
  totalCount: number;
  schemaList: Schema[];
}

export interface QueryCredentialRequest {
  credId: string;
}

export interface QueryCredentialResponse {
  credStatus: Credential | undefined;
}

export interface QueryCredentialsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryCredentialsResponse {
  totalCount: number;
  credentials: Credential[];
}

export interface QueryDidDocumentRequest {
  didId: string;
}

export interface QueryDidDocumentResponse {
  didDocument: Did | undefined;
  didDocumentMetadata: Metadata | undefined;
}

export interface QueryDidDocumentsRequest {
  count: boolean;
  pagination: PageRequest | undefined;
}

export interface QueryDidDocumentsResponse {
  totalDidCount: number;
  didDocList: QueryDidDocumentResponse[];
}

const baseQuerySchemaRequest: object = { schemaId: "" };

export const QuerySchemaRequest = {
  encode(
    message: QuerySchemaRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.schemaId !== "") {
      writer.uint32(10).string(message.schemaId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QuerySchemaRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQuerySchemaRequest } as QuerySchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySchemaRequest {
    const message = { ...baseQuerySchemaRequest } as QuerySchemaRequest;
    if (object.schemaId !== undefined && object.schemaId !== null) {
      message.schemaId = String(object.schemaId);
    } else {
      message.schemaId = "";
    }
    return message;
  },

  toJSON(message: QuerySchemaRequest): unknown {
    const obj: any = {};
    message.schemaId !== undefined && (obj.schemaId = message.schemaId);
    return obj;
  },

  fromPartial(object: DeepPartial<QuerySchemaRequest>): QuerySchemaRequest {
    const message = { ...baseQuerySchemaRequest } as QuerySchemaRequest;
    if (object.schemaId !== undefined && object.schemaId !== null) {
      message.schemaId = object.schemaId;
    } else {
      message.schemaId = "";
    }
    return message;
  },
};

const baseQuerySchemaResponse: object = {};

export const QuerySchemaResponse = {
  encode(
    message: QuerySchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.schema) {
      Schema.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QuerySchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQuerySchemaResponse } as QuerySchemaResponse;
    message.schema = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schema.push(Schema.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySchemaResponse {
    const message = { ...baseQuerySchemaResponse } as QuerySchemaResponse;
    message.schema = [];
    if (object.schema !== undefined && object.schema !== null) {
      for (const e of object.schema) {
        message.schema.push(Schema.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QuerySchemaResponse): unknown {
    const obj: any = {};
    if (message.schema) {
      obj.schema = message.schema.map((e) =>
        e ? Schema.toJSON(e) : undefined
      );
    } else {
      obj.schema = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QuerySchemaResponse>): QuerySchemaResponse {
    const message = { ...baseQuerySchemaResponse } as QuerySchemaResponse;
    message.schema = [];
    if (object.schema !== undefined && object.schema !== null) {
      for (const e of object.schema) {
        message.schema.push(Schema.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQuerySchemasRequest: object = {};

export const QuerySchemasRequest = {
  encode(
    message: QuerySchemasRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QuerySchemasRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQuerySchemasRequest } as QuerySchemasRequest;
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

  fromJSON(object: any): QuerySchemasRequest {
    const message = { ...baseQuerySchemasRequest } as QuerySchemasRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QuerySchemasRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QuerySchemasRequest>): QuerySchemasRequest {
    const message = { ...baseQuerySchemasRequest } as QuerySchemasRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQuerySchemasResponse: object = { totalCount: 0 };

export const QuerySchemasResponse = {
  encode(
    message: QuerySchemasResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.totalCount !== 0) {
      writer.uint32(8).uint64(message.totalCount);
    }
    for (const v of message.schemaList) {
      Schema.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QuerySchemasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQuerySchemasResponse } as QuerySchemasResponse;
    message.schemaList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.totalCount = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.schemaList.push(Schema.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySchemasResponse {
    const message = { ...baseQuerySchemasResponse } as QuerySchemasResponse;
    message.schemaList = [];
    if (object.totalCount !== undefined && object.totalCount !== null) {
      message.totalCount = Number(object.totalCount);
    } else {
      message.totalCount = 0;
    }
    if (object.schemaList !== undefined && object.schemaList !== null) {
      for (const e of object.schemaList) {
        message.schemaList.push(Schema.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QuerySchemasResponse): unknown {
    const obj: any = {};
    message.totalCount !== undefined && (obj.totalCount = message.totalCount);
    if (message.schemaList) {
      obj.schemaList = message.schemaList.map((e) =>
        e ? Schema.toJSON(e) : undefined
      );
    } else {
      obj.schemaList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QuerySchemasResponse>): QuerySchemasResponse {
    const message = { ...baseQuerySchemasResponse } as QuerySchemasResponse;
    message.schemaList = [];
    if (object.totalCount !== undefined && object.totalCount !== null) {
      message.totalCount = object.totalCount;
    } else {
      message.totalCount = 0;
    }
    if (object.schemaList !== undefined && object.schemaList !== null) {
      for (const e of object.schemaList) {
        message.schemaList.push(Schema.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryCredentialRequest: object = { credId: "" };

export const QueryCredentialRequest = {
  encode(
    message: QueryCredentialRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.credId !== "") {
      writer.uint32(10).string(message.credId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCredentialRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCredentialRequest } as QueryCredentialRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.credId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCredentialRequest {
    const message = { ...baseQueryCredentialRequest } as QueryCredentialRequest;
    if (object.credId !== undefined && object.credId !== null) {
      message.credId = String(object.credId);
    } else {
      message.credId = "";
    }
    return message;
  },

  toJSON(message: QueryCredentialRequest): unknown {
    const obj: any = {};
    message.credId !== undefined && (obj.credId = message.credId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCredentialRequest>
  ): QueryCredentialRequest {
    const message = { ...baseQueryCredentialRequest } as QueryCredentialRequest;
    if (object.credId !== undefined && object.credId !== null) {
      message.credId = object.credId;
    } else {
      message.credId = "";
    }
    return message;
  },
};

const baseQueryCredentialResponse: object = {};

export const QueryCredentialResponse = {
  encode(
    message: QueryCredentialResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.credStatus !== undefined) {
      Credential.encode(message.credStatus, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCredentialResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCredentialResponse,
    } as QueryCredentialResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.credStatus = Credential.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCredentialResponse {
    const message = {
      ...baseQueryCredentialResponse,
    } as QueryCredentialResponse;
    if (object.credStatus !== undefined && object.credStatus !== null) {
      message.credStatus = Credential.fromJSON(object.credStatus);
    } else {
      message.credStatus = undefined;
    }
    return message;
  },

  toJSON(message: QueryCredentialResponse): unknown {
    const obj: any = {};
    message.credStatus !== undefined &&
      (obj.credStatus = message.credStatus
        ? Credential.toJSON(message.credStatus)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCredentialResponse>
  ): QueryCredentialResponse {
    const message = {
      ...baseQueryCredentialResponse,
    } as QueryCredentialResponse;
    if (object.credStatus !== undefined && object.credStatus !== null) {
      message.credStatus = Credential.fromPartial(object.credStatus);
    } else {
      message.credStatus = undefined;
    }
    return message;
  },
};

const baseQueryCredentialsRequest: object = {};

export const QueryCredentialsRequest = {
  encode(
    message: QueryCredentialsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCredentialsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCredentialsRequest,
    } as QueryCredentialsRequest;
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

  fromJSON(object: any): QueryCredentialsRequest {
    const message = {
      ...baseQueryCredentialsRequest,
    } as QueryCredentialsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCredentialsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCredentialsRequest>
  ): QueryCredentialsRequest {
    const message = {
      ...baseQueryCredentialsRequest,
    } as QueryCredentialsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCredentialsResponse: object = { totalCount: 0 };

export const QueryCredentialsResponse = {
  encode(
    message: QueryCredentialsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.totalCount !== 0) {
      writer.uint32(8).uint64(message.totalCount);
    }
    for (const v of message.credentials) {
      Credential.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCredentialsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCredentialsResponse,
    } as QueryCredentialsResponse;
    message.credentials = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.totalCount = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.credentials.push(Credential.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCredentialsResponse {
    const message = {
      ...baseQueryCredentialsResponse,
    } as QueryCredentialsResponse;
    message.credentials = [];
    if (object.totalCount !== undefined && object.totalCount !== null) {
      message.totalCount = Number(object.totalCount);
    } else {
      message.totalCount = 0;
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryCredentialsResponse): unknown {
    const obj: any = {};
    message.totalCount !== undefined && (obj.totalCount = message.totalCount);
    if (message.credentials) {
      obj.credentials = message.credentials.map((e) =>
        e ? Credential.toJSON(e) : undefined
      );
    } else {
      obj.credentials = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCredentialsResponse>
  ): QueryCredentialsResponse {
    const message = {
      ...baseQueryCredentialsResponse,
    } as QueryCredentialsResponse;
    message.credentials = [];
    if (object.totalCount !== undefined && object.totalCount !== null) {
      message.totalCount = object.totalCount;
    } else {
      message.totalCount = 0;
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryDidDocumentRequest: object = { didId: "" };

export const QueryDidDocumentRequest = {
  encode(
    message: QueryDidDocumentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.didId !== "") {
      writer.uint32(10).string(message.didId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDidDocumentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentRequest {
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    if (object.didId !== undefined && object.didId !== null) {
      message.didId = String(object.didId);
    } else {
      message.didId = "";
    }
    return message;
  },

  toJSON(message: QueryDidDocumentRequest): unknown {
    const obj: any = {};
    message.didId !== undefined && (obj.didId = message.didId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentRequest>
  ): QueryDidDocumentRequest {
    const message = {
      ...baseQueryDidDocumentRequest,
    } as QueryDidDocumentRequest;
    if (object.didId !== undefined && object.didId !== null) {
      message.didId = object.didId;
    } else {
      message.didId = "";
    }
    return message;
  },
};

const baseQueryDidDocumentResponse: object = {};

export const QueryDidDocumentResponse = {
  encode(
    message: QueryDidDocumentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.didDocument !== undefined) {
      Did.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
    }
    if (message.didDocumentMetadata !== undefined) {
      Metadata.encode(
        message.didDocumentMetadata,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didDocument = Did.decode(reader, reader.uint32());
          break;
        case 2:
          message.didDocumentMetadata = Metadata.decode(
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

  fromJSON(object: any): QueryDidDocumentResponse {
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = Did.fromJSON(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    if (
      object.didDocumentMetadata !== undefined &&
      object.didDocumentMetadata !== null
    ) {
      message.didDocumentMetadata = Metadata.fromJSON(
        object.didDocumentMetadata
      );
    } else {
      message.didDocumentMetadata = undefined;
    }
    return message;
  },

  toJSON(message: QueryDidDocumentResponse): unknown {
    const obj: any = {};
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? Did.toJSON(message.didDocument)
        : undefined);
    message.didDocumentMetadata !== undefined &&
      (obj.didDocumentMetadata = message.didDocumentMetadata
        ? Metadata.toJSON(message.didDocumentMetadata)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentResponse>
  ): QueryDidDocumentResponse {
    const message = {
      ...baseQueryDidDocumentResponse,
    } as QueryDidDocumentResponse;
    if (object.didDocument !== undefined && object.didDocument !== null) {
      message.didDocument = Did.fromPartial(object.didDocument);
    } else {
      message.didDocument = undefined;
    }
    if (
      object.didDocumentMetadata !== undefined &&
      object.didDocumentMetadata !== null
    ) {
      message.didDocumentMetadata = Metadata.fromPartial(
        object.didDocumentMetadata
      );
    } else {
      message.didDocumentMetadata = undefined;
    }
    return message;
  },
};

const baseQueryDidDocumentsRequest: object = { count: false };

export const QueryDidDocumentsRequest = {
  encode(
    message: QueryDidDocumentsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.count === true) {
      writer.uint32(8).bool(message.count);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentsRequest,
    } as QueryDidDocumentsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.count = reader.bool();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentsRequest {
    const message = {
      ...baseQueryDidDocumentsRequest,
    } as QueryDidDocumentsRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = Boolean(object.count);
    } else {
      message.count = false;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDidDocumentsRequest): unknown {
    const obj: any = {};
    message.count !== undefined && (obj.count = message.count);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentsRequest>
  ): QueryDidDocumentsRequest {
    const message = {
      ...baseQueryDidDocumentsRequest,
    } as QueryDidDocumentsRequest;
    if (object.count !== undefined && object.count !== null) {
      message.count = object.count;
    } else {
      message.count = false;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDidDocumentsResponse: object = { totalDidCount: 0 };

export const QueryDidDocumentsResponse = {
  encode(
    message: QueryDidDocumentsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.totalDidCount !== 0) {
      writer.uint32(8).uint64(message.totalDidCount);
    }
    for (const v of message.didDocList) {
      QueryDidDocumentResponse.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryDidDocumentsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryDidDocumentsResponse,
    } as QueryDidDocumentsResponse;
    message.didDocList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.totalDidCount = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.didDocList.push(
            QueryDidDocumentResponse.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDidDocumentsResponse {
    const message = {
      ...baseQueryDidDocumentsResponse,
    } as QueryDidDocumentsResponse;
    message.didDocList = [];
    if (object.totalDidCount !== undefined && object.totalDidCount !== null) {
      message.totalDidCount = Number(object.totalDidCount);
    } else {
      message.totalDidCount = 0;
    }
    if (object.didDocList !== undefined && object.didDocList !== null) {
      for (const e of object.didDocList) {
        message.didDocList.push(QueryDidDocumentResponse.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryDidDocumentsResponse): unknown {
    const obj: any = {};
    message.totalDidCount !== undefined &&
      (obj.totalDidCount = message.totalDidCount);
    if (message.didDocList) {
      obj.didDocList = message.didDocList.map((e) =>
        e ? QueryDidDocumentResponse.toJSON(e) : undefined
      );
    } else {
      obj.didDocList = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryDidDocumentsResponse>
  ): QueryDidDocumentsResponse {
    const message = {
      ...baseQueryDidDocumentsResponse,
    } as QueryDidDocumentsResponse;
    message.didDocList = [];
    if (object.totalDidCount !== undefined && object.totalDidCount !== null) {
      message.totalDidCount = object.totalDidCount;
    } else {
      message.totalDidCount = 0;
    }
    if (object.didDocList !== undefined && object.didDocList !== null) {
      for (const e of object.didDocList) {
        message.didDocList.push(QueryDidDocumentResponse.fromPartial(e));
      }
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Get the Schema Document for a specified schema id */
  QuerySchema(request: QuerySchemaRequest): Promise<QuerySchemaResponse>;
  /** Get the count and list of registered Schemas */
  QuerySchemas(request: QuerySchemasRequest): Promise<QuerySchemasResponse>;
  /** Get the Did Document for a specified DID id */
  QueryDidDocument(
    request: QueryDidDocumentRequest
  ): Promise<QueryDidDocumentResponse>;
  /** Get the count and list of registered Did Documents */
  QueryDidDocuments(
    request: QueryDidDocumentsRequest
  ): Promise<QueryDidDocumentsResponse>;
  /** Get the Credential Status for a given credential id */
  QueryCredential(
    request: QueryCredentialRequest
  ): Promise<QueryCredentialResponse>;
  /** Get all the registed Credential Statuses */
  QueryCredentials(
    request: QueryCredentialsRequest
  ): Promise<QueryCredentialsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  QuerySchema(request: QuerySchemaRequest): Promise<QuerySchemaResponse> {
    const data = QuerySchemaRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.ssi.Query",
      "QuerySchema",
      data
    );
    return promise.then((data) => QuerySchemaResponse.decode(new Reader(data)));
  }

  QuerySchemas(request: QuerySchemasRequest): Promise<QuerySchemasResponse> {
    const data = QuerySchemasRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.ssi.Query",
      "QuerySchemas",
      data
    );
    return promise.then((data) =>
      QuerySchemasResponse.decode(new Reader(data))
    );
  }

  QueryDidDocument(
    request: QueryDidDocumentRequest
  ): Promise<QueryDidDocumentResponse> {
    const data = QueryDidDocumentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.ssi.Query",
      "QueryDidDocument",
      data
    );
    return promise.then((data) =>
      QueryDidDocumentResponse.decode(new Reader(data))
    );
  }

  QueryDidDocuments(
    request: QueryDidDocumentsRequest
  ): Promise<QueryDidDocumentsResponse> {
    const data = QueryDidDocumentsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.ssi.Query",
      "QueryDidDocuments",
      data
    );
    return promise.then((data) =>
      QueryDidDocumentsResponse.decode(new Reader(data))
    );
  }

  QueryCredential(
    request: QueryCredentialRequest
  ): Promise<QueryCredentialResponse> {
    const data = QueryCredentialRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.ssi.Query",
      "QueryCredential",
      data
    );
    return promise.then((data) =>
      QueryCredentialResponse.decode(new Reader(data))
    );
  }

  QueryCredentials(
    request: QueryCredentialsRequest
  ): Promise<QueryCredentialsResponse> {
    const data = QueryCredentialsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "mandechain.ssi.Query",
      "QueryCredentials",
      data
    );
    return promise.then((data) =>
      QueryCredentialsResponse.decode(new Reader(data))
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
