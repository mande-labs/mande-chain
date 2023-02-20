/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface SsiClaim {
  id?: string;
  currentStatus?: string;
  statusReason?: string;
}

export interface SsiCredential {
  claim?: SsiClaim;
  issuer?: string;
  issuanceDate?: string;
  expirationDate?: string;
  credentialHash?: string;
  proof?: SsiCredentialProof;
}

export interface SsiCredentialProof {
  type?: string;
  created?: string;
  updated?: string;
  verificationMethod?: string;
  proofPurpose?: string;
  proofValue?: string;
}

export interface SsiCredentialStatus {
  claim?: SsiClaim;
  issuer?: string;
  issuanceDate?: string;
  expirationDate?: string;
  credentialHash?: string;
}

export interface SsiDid {
  context?: string[];
  id?: string;
  controller?: string[];
  alsoKnownAs?: string[];
  verificationMethod?: SsiVerificationMethod[];
  authentication?: string[];
  assertionMethod?: string[];
  keyAgreement?: string[];
  capabilityInvocation?: string[];
  capabilityDelegation?: string[];
  service?: SsiService[];
}

export interface SsiMetadata {
  created?: string;
  updated?: string;
  deactivated?: boolean;
  versionId?: string;
}

export interface SsiMsgCreateDIDResponse {
  /** @format uint64 */
  id?: string;
}

export interface SsiMsgCreateSchemaResponse {
  /** @format uint64 */
  id?: string;
}

export interface SsiMsgDeactivateDIDResponse {
  /** @format uint64 */
  id?: string;
}

export interface SsiMsgRegisterCredentialStatusResponse {
  /** @format uint64 */
  id?: string;
}

export interface SsiMsgUpdateDIDResponse {
  updateId?: string;
}

export interface SsiQueryCredentialResponse {
  credStatus?: SsiCredential;
}

export interface SsiQueryCredentialsResponse {
  /** @format uint64 */
  totalCount?: string;
  credentials?: SsiCredential[];
}

export interface SsiQueryDidDocumentResponse {
  didDocument?: SsiDid;
  didDocumentMetadata?: SsiMetadata;
}

export interface SsiQueryDidDocumentsResponse {
  /** @format uint64 */
  totalDidCount?: string;
  didDocList?: SsiQueryDidDocumentResponse[];
}

export interface SsiQuerySchemaResponse {
  schema?: SsiSchema[];
}

export interface SsiQuerySchemasResponse {
  /** @format uint64 */
  totalCount?: string;
  schemaList?: SsiSchema[];
}

export interface SsiSchema {
  type?: string;
  modelVersion?: string;
  id?: string;
  name?: string;
  author?: string;
  authored?: string;
  schema?: SsiSchemaProperty;
  proof?: SsiSchemaProof;
}

export interface SsiSchemaDocument {
  type?: string;
  modelVersion?: string;
  id?: string;
  name?: string;
  author?: string;
  authored?: string;
  schema?: SsiSchemaProperty;
}

export interface SsiSchemaProof {
  type?: string;
  created?: string;
  verificationMethod?: string;
  proofPurpose?: string;
  proofValue?: string;
}

export interface SsiSchemaProperty {
  schema?: string;
  description?: string;
  type?: string;
  properties?: string;
  required?: string[];
  additionalProperties?: boolean;
}

export interface SsiService {
  id?: string;
  type?: string;
  serviceEndpoint?: string;
}

export interface SsiSignInfo {
  verification_method_id?: string;
  signature?: string;
}

export interface SsiVerificationMethod {
  id?: string;
  type?: string;
  controller?: string;
  publicKeyMultibase?: string;
  blockchainAccountId?: string;
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title ssi/v1/credential.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryCredentials
   * @summary Get all the registed Credential Statuses
   * @request GET:/mande-chain/ssi/credential
   */
  queryQueryCredentials = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<SsiQueryCredentialsResponse, RpcStatus>({
      path: `/mande-chain/ssi/credential`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryCredential
   * @summary Get the Credential Status for a given credential id
   * @request GET:/mande-chain/ssi/credential/{credId}
   */
  queryQueryCredential = (credId: string, params: RequestParams = {}) =>
    this.request<SsiQueryCredentialResponse, RpcStatus>({
      path: `/mande-chain/ssi/credential/${credId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryDidDocuments
   * @summary Get the count and list of registered Did Documents
   * @request GET:/mande-chain/ssi/did
   */
  queryQueryDidDocuments = (
    query?: {
      count?: boolean;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<SsiQueryDidDocumentsResponse, RpcStatus>({
      path: `/mande-chain/ssi/did`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQueryDidDocument
   * @summary Get the Did Document for a specified DID id
   * @request GET:/mande-chain/ssi/did/{didId}
   */
  queryQueryDidDocument = (didId: string, params: RequestParams = {}) =>
    this.request<SsiQueryDidDocumentResponse, RpcStatus>({
      path: `/mande-chain/ssi/did/${didId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQuerySchemas
   * @summary Get the count and list of registered Schemas
   * @request GET:/mande-chain/ssi/schema
   */
  queryQuerySchemas = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<SsiQuerySchemasResponse, RpcStatus>({
      path: `/mande-chain/ssi/schema`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryQuerySchema
   * @summary Get the Schema Document for a specified schema id
   * @request GET:/mande-chain/ssi/schema/{schemaId}
   */
  queryQuerySchema = (schemaId: string, params: RequestParams = {}) =>
    this.request<SsiQuerySchemaResponse, RpcStatus>({
      path: `/mande-chain/ssi/schema/${schemaId}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
