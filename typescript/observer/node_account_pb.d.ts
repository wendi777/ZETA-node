// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file observer/node_account.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { PubKeySet } from "../pkg/crypto/crypto_pb.js";

/**
 * @generated from enum zetachain.zetacore.observer.NodeStatus
 */
export declare enum NodeStatus {
  /**
   * @generated from enum value: Unknown = 0;
   */
  Unknown = 0,

  /**
   * @generated from enum value: Whitelisted = 1;
   */
  Whitelisted = 1,

  /**
   * @generated from enum value: Standby = 2;
   */
  Standby = 2,

  /**
   * @generated from enum value: Ready = 3;
   */
  Ready = 3,

  /**
   * @generated from enum value: Active = 4;
   */
  Active = 4,

  /**
   * @generated from enum value: Disabled = 5;
   */
  Disabled = 5,
}

/**
 * @generated from message zetachain.zetacore.observer.NodeAccount
 */
export declare class NodeAccount extends Message<NodeAccount> {
  /**
   * @generated from field: string operator = 1;
   */
  operator: string;

  /**
   * @generated from field: string granteeAddress = 2;
   */
  granteeAddress: string;

  /**
   * @generated from field: crypto.PubKeySet granteePubkey = 3;
   */
  granteePubkey?: PubKeySet;

  /**
   * @generated from field: zetachain.zetacore.observer.NodeStatus nodeStatus = 4;
   */
  nodeStatus: NodeStatus;

  constructor(data?: PartialMessage<NodeAccount>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.NodeAccount";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeAccount;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeAccount;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeAccount;

  static equals(a: NodeAccount | PlainMessage<NodeAccount> | undefined, b: NodeAccount | PlainMessage<NodeAccount> | undefined): boolean;
}

