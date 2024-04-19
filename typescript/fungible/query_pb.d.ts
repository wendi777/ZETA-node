// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file fungible/query.proto (package zetachain.zetacore.fungible, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { ForeignCoins } from "./foreign_coins_pb.js";
import type { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination_pb.js";
import type { SystemContract } from "./system_contract_pb.js";

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetForeignCoinsRequest
 */
export declare class QueryGetForeignCoinsRequest extends Message<QueryGetForeignCoinsRequest> {
  /**
   * @generated from field: string index = 1;
   */
  index: string;

  constructor(data?: PartialMessage<QueryGetForeignCoinsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetForeignCoinsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetForeignCoinsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetForeignCoinsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetForeignCoinsRequest;

  static equals(a: QueryGetForeignCoinsRequest | PlainMessage<QueryGetForeignCoinsRequest> | undefined, b: QueryGetForeignCoinsRequest | PlainMessage<QueryGetForeignCoinsRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetForeignCoinsResponse
 */
export declare class QueryGetForeignCoinsResponse extends Message<QueryGetForeignCoinsResponse> {
  /**
   * @generated from field: zetachain.zetacore.fungible.ForeignCoins foreignCoins = 1;
   */
  foreignCoins?: ForeignCoins;

  constructor(data?: PartialMessage<QueryGetForeignCoinsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetForeignCoinsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetForeignCoinsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetForeignCoinsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetForeignCoinsResponse;

  static equals(a: QueryGetForeignCoinsResponse | PlainMessage<QueryGetForeignCoinsResponse> | undefined, b: QueryGetForeignCoinsResponse | PlainMessage<QueryGetForeignCoinsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryAllForeignCoinsRequest
 */
export declare class QueryAllForeignCoinsRequest extends Message<QueryAllForeignCoinsRequest> {
  /**
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryAllForeignCoinsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryAllForeignCoinsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllForeignCoinsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllForeignCoinsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllForeignCoinsRequest;

  static equals(a: QueryAllForeignCoinsRequest | PlainMessage<QueryAllForeignCoinsRequest> | undefined, b: QueryAllForeignCoinsRequest | PlainMessage<QueryAllForeignCoinsRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryAllForeignCoinsResponse
 */
export declare class QueryAllForeignCoinsResponse extends Message<QueryAllForeignCoinsResponse> {
  /**
   * @generated from field: repeated zetachain.zetacore.fungible.ForeignCoins foreignCoins = 1;
   */
  foreignCoins: ForeignCoins[];

  /**
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryAllForeignCoinsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryAllForeignCoinsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllForeignCoinsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllForeignCoinsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllForeignCoinsResponse;

  static equals(a: QueryAllForeignCoinsResponse | PlainMessage<QueryAllForeignCoinsResponse> | undefined, b: QueryAllForeignCoinsResponse | PlainMessage<QueryAllForeignCoinsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetSystemContractRequest
 */
export declare class QueryGetSystemContractRequest extends Message<QueryGetSystemContractRequest> {
  constructor(data?: PartialMessage<QueryGetSystemContractRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetSystemContractRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetSystemContractRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetSystemContractRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetSystemContractRequest;

  static equals(a: QueryGetSystemContractRequest | PlainMessage<QueryGetSystemContractRequest> | undefined, b: QueryGetSystemContractRequest | PlainMessage<QueryGetSystemContractRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetSystemContractResponse
 */
export declare class QueryGetSystemContractResponse extends Message<QueryGetSystemContractResponse> {
  /**
   * @generated from field: zetachain.zetacore.fungible.SystemContract SystemContract = 1;
   */
  SystemContract?: SystemContract;

  constructor(data?: PartialMessage<QueryGetSystemContractResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetSystemContractResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetSystemContractResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetSystemContractResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetSystemContractResponse;

  static equals(a: QueryGetSystemContractResponse | PlainMessage<QueryGetSystemContractResponse> | undefined, b: QueryGetSystemContractResponse | PlainMessage<QueryGetSystemContractResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetGasStabilityPoolAddress
 */
export declare class QueryGetGasStabilityPoolAddress extends Message<QueryGetGasStabilityPoolAddress> {
  constructor(data?: PartialMessage<QueryGetGasStabilityPoolAddress>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetGasStabilityPoolAddress";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetGasStabilityPoolAddress;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolAddress;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolAddress;

  static equals(a: QueryGetGasStabilityPoolAddress | PlainMessage<QueryGetGasStabilityPoolAddress> | undefined, b: QueryGetGasStabilityPoolAddress | PlainMessage<QueryGetGasStabilityPoolAddress> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetGasStabilityPoolAddressResponse
 */
export declare class QueryGetGasStabilityPoolAddressResponse extends Message<QueryGetGasStabilityPoolAddressResponse> {
  /**
   * @generated from field: string cosmos_address = 1;
   */
  cosmosAddress: string;

  /**
   * @generated from field: string evm_address = 2;
   */
  evmAddress: string;

  constructor(data?: PartialMessage<QueryGetGasStabilityPoolAddressResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetGasStabilityPoolAddressResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetGasStabilityPoolAddressResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolAddressResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolAddressResponse;

  static equals(a: QueryGetGasStabilityPoolAddressResponse | PlainMessage<QueryGetGasStabilityPoolAddressResponse> | undefined, b: QueryGetGasStabilityPoolAddressResponse | PlainMessage<QueryGetGasStabilityPoolAddressResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetGasStabilityPoolBalance
 */
export declare class QueryGetGasStabilityPoolBalance extends Message<QueryGetGasStabilityPoolBalance> {
  /**
   * @generated from field: int64 chain_id = 1;
   */
  chainId: bigint;

  constructor(data?: PartialMessage<QueryGetGasStabilityPoolBalance>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetGasStabilityPoolBalance";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetGasStabilityPoolBalance;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolBalance;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolBalance;

  static equals(a: QueryGetGasStabilityPoolBalance | PlainMessage<QueryGetGasStabilityPoolBalance> | undefined, b: QueryGetGasStabilityPoolBalance | PlainMessage<QueryGetGasStabilityPoolBalance> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryGetGasStabilityPoolBalanceResponse
 */
export declare class QueryGetGasStabilityPoolBalanceResponse extends Message<QueryGetGasStabilityPoolBalanceResponse> {
  /**
   * @generated from field: string balance = 2;
   */
  balance: string;

  constructor(data?: PartialMessage<QueryGetGasStabilityPoolBalanceResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryGetGasStabilityPoolBalanceResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetGasStabilityPoolBalanceResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolBalanceResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetGasStabilityPoolBalanceResponse;

  static equals(a: QueryGetGasStabilityPoolBalanceResponse | PlainMessage<QueryGetGasStabilityPoolBalanceResponse> | undefined, b: QueryGetGasStabilityPoolBalanceResponse | PlainMessage<QueryGetGasStabilityPoolBalanceResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalance
 */
export declare class QueryAllGasStabilityPoolBalance extends Message<QueryAllGasStabilityPoolBalance> {
  constructor(data?: PartialMessage<QueryAllGasStabilityPoolBalance>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalance";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllGasStabilityPoolBalance;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllGasStabilityPoolBalance;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllGasStabilityPoolBalance;

  static equals(a: QueryAllGasStabilityPoolBalance | PlainMessage<QueryAllGasStabilityPoolBalance> | undefined, b: QueryAllGasStabilityPoolBalance | PlainMessage<QueryAllGasStabilityPoolBalance> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalanceResponse
 */
export declare class QueryAllGasStabilityPoolBalanceResponse extends Message<QueryAllGasStabilityPoolBalanceResponse> {
  /**
   * @generated from field: repeated zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalanceResponse.Balance balances = 1;
   */
  balances: QueryAllGasStabilityPoolBalanceResponse_Balance[];

  constructor(data?: PartialMessage<QueryAllGasStabilityPoolBalanceResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalanceResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllGasStabilityPoolBalanceResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllGasStabilityPoolBalanceResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllGasStabilityPoolBalanceResponse;

  static equals(a: QueryAllGasStabilityPoolBalanceResponse | PlainMessage<QueryAllGasStabilityPoolBalanceResponse> | undefined, b: QueryAllGasStabilityPoolBalanceResponse | PlainMessage<QueryAllGasStabilityPoolBalanceResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalanceResponse.Balance
 */
export declare class QueryAllGasStabilityPoolBalanceResponse_Balance extends Message<QueryAllGasStabilityPoolBalanceResponse_Balance> {
  /**
   * @generated from field: int64 chain_id = 1;
   */
  chainId: bigint;

  /**
   * @generated from field: string balance = 2;
   */
  balance: string;

  constructor(data?: PartialMessage<QueryAllGasStabilityPoolBalanceResponse_Balance>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryAllGasStabilityPoolBalanceResponse.Balance";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllGasStabilityPoolBalanceResponse_Balance;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllGasStabilityPoolBalanceResponse_Balance;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllGasStabilityPoolBalanceResponse_Balance;

  static equals(a: QueryAllGasStabilityPoolBalanceResponse_Balance | PlainMessage<QueryAllGasStabilityPoolBalanceResponse_Balance> | undefined, b: QueryAllGasStabilityPoolBalanceResponse_Balance | PlainMessage<QueryAllGasStabilityPoolBalanceResponse_Balance> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryCodeHashRequest
 */
export declare class QueryCodeHashRequest extends Message<QueryCodeHashRequest> {
  /**
   * @generated from field: string address = 1;
   */
  address: string;

  constructor(data?: PartialMessage<QueryCodeHashRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryCodeHashRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryCodeHashRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryCodeHashRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryCodeHashRequest;

  static equals(a: QueryCodeHashRequest | PlainMessage<QueryCodeHashRequest> | undefined, b: QueryCodeHashRequest | PlainMessage<QueryCodeHashRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.fungible.QueryCodeHashResponse
 */
export declare class QueryCodeHashResponse extends Message<QueryCodeHashResponse> {
  /**
   * @generated from field: string code_hash = 1;
   */
  codeHash: string;

  constructor(data?: PartialMessage<QueryCodeHashResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.fungible.QueryCodeHashResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryCodeHashResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryCodeHashResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryCodeHashResponse;

  static equals(a: QueryCodeHashResponse | PlainMessage<QueryCodeHashResponse> | undefined, b: QueryCodeHashResponse | PlainMessage<QueryCodeHashResponse> | undefined): boolean;
}

