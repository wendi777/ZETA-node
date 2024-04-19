// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file lightclient/verification_flags.proto (package zetachain.zetacore.lightclient, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * VerificationFlags is a structure containing information which chain types are enabled for block header verification
 *
 * @generated from message zetachain.zetacore.lightclient.VerificationFlags
 */
export declare class VerificationFlags extends Message<VerificationFlags> {
  /**
   * @generated from field: bool ethTypeChainEnabled = 1;
   */
  ethTypeChainEnabled: boolean;

  /**
   * @generated from field: bool btcTypeChainEnabled = 2;
   */
  btcTypeChainEnabled: boolean;

  constructor(data?: PartialMessage<VerificationFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.VerificationFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VerificationFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VerificationFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VerificationFlags;

  static equals(a: VerificationFlags | PlainMessage<VerificationFlags> | undefined, b: VerificationFlags | PlainMessage<VerificationFlags> | undefined): boolean;
}

