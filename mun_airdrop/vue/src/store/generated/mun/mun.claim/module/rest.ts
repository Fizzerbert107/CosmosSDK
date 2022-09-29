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

export enum ClaimAction {
  ActionInitialClaim = "ActionInitialClaim",
  ActionStake = "ActionStake",
  ActionVote = "ActionVote",
  ActionSwap = "ActionSwap",
}

export interface ClaimClaimAuthorization {
  contract_address?: string;
  action?: ClaimAction;
}

export interface ClaimClaimRecord {
  address?: string;
  initial_claimable_amount?: V1Beta1Coin[];
  action_completed?: boolean[];
}

export interface ClaimMsgClaimForResponse {
  address?: string;
  claimed_amount?: V1Beta1Coin[];
}

export interface ClaimMsgInitialClaimResponse {
  address?: string;
  claimed_amount?: V1Beta1Coin[];
}

/**
 * Params defines the parameters for the module.
 */
export interface ClaimParams {
  airdrop_enabled?: boolean;

  /** @format date-time */
  airdrop_start_time?: string;
  duration_until_decay?: string;
  duration_of_decay?: string;
  claim_denom?: string;
  allowed_claimers?: ClaimClaimAuthorization[];
}

export interface ClaimQueryClaimRecordResponse {
  claim_record?: ClaimClaimRecord;
}

export interface ClaimQueryClaimableForActionResponse {
  coins?: V1Beta1Coin[];
}

export interface ClaimQueryModuleAccountBalanceResponse {
  moduleAccountBalance?: V1Beta1Coin[];
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface ClaimQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: ClaimParams;
}

export interface ClaimQueryTotalClaimableResponse {
  coins?: V1Beta1Coin[];
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
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
 * @title claim/claim_record.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryClaimRecord
   * @summary Queries a list of ClaimRecord items.
   * @request GET:/mun/claim/claim_record/{address}
   */
  queryClaimRecord = (address: string, params: RequestParams = {}) =>
    this.request<ClaimQueryClaimRecordResponse, RpcStatus>({
      path: `/mun/claim/claim_record/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryClaimableForAction
   * @summary Queries a list of ClaimableForAction items.
   * @request GET:/mun/claim/claimable_for_action/{address}
   */
  queryClaimableForAction = (
    address: string,
    query?: { action?: "ActionInitialClaim" | "ActionStake" | "ActionVote" | "ActionSwap" },
    params: RequestParams = {},
  ) =>
    this.request<ClaimQueryClaimableForActionResponse, RpcStatus>({
      path: `/mun/claim/claimable_for_action/${address}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryModuleAccountBalance
   * @summary Queries a list of ModuleAccountBalance items.
   * @request GET:/mun/claim/module_account_balance
   */
  queryModuleAccountBalance = (params: RequestParams = {}) =>
    this.request<ClaimQueryModuleAccountBalanceResponse, RpcStatus>({
      path: `/mun/claim/module_account_balance`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/mun/claim/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<ClaimQueryParamsResponse, RpcStatus>({
      path: `/mun/claim/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTotalClaimable
   * @summary Queries a list of TotalClaimable items.
   * @request GET:/mun/claim/total_claimable/{address}
   */
  queryTotalClaimable = (address: string, params: RequestParams = {}) =>
    this.request<ClaimQueryTotalClaimableResponse, RpcStatus>({
      path: `/mun/claim/total_claimable/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });
}