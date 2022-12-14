/* eslint-disable */
import { Action, actionFromJSON, actionToJSON } from "../claim/claim_record";
import { Timestamp } from "../google/protobuf/timestamp";
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mun.claim";

export interface ClaimAuthorization {
  contract_address: string;
  action: Action;
}

/** Params defines the parameters for the module. */
export interface Params {
  airdrop_enabled: boolean;
  airdrop_start_time: Date | undefined;
  duration_until_decay: Duration | undefined;
  duration_of_decay: Duration | undefined;
  /** denom of claimable asset */
  claim_denom: string;
  /** list of contracts and their allowed claim actions */
  allowed_claimers: ClaimAuthorization[];
}

const baseClaimAuthorization: object = { contract_address: "", action: 0 };

export const ClaimAuthorization = {
  encode(
    message: ClaimAuthorization,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.contract_address !== "") {
      writer.uint32(10).string(message.contract_address);
    }
    if (message.action !== 0) {
      writer.uint32(16).int32(message.action);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ClaimAuthorization {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClaimAuthorization } as ClaimAuthorization;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.contract_address = reader.string();
          break;
        case 2:
          message.action = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimAuthorization {
    const message = { ...baseClaimAuthorization } as ClaimAuthorization;
    if (
      object.contract_address !== undefined &&
      object.contract_address !== null
    ) {
      message.contract_address = String(object.contract_address);
    } else {
      message.contract_address = "";
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = actionFromJSON(object.action);
    } else {
      message.action = 0;
    }
    return message;
  },

  toJSON(message: ClaimAuthorization): unknown {
    const obj: any = {};
    message.contract_address !== undefined &&
      (obj.contract_address = message.contract_address);
    message.action !== undefined && (obj.action = actionToJSON(message.action));
    return obj;
  },

  fromPartial(object: DeepPartial<ClaimAuthorization>): ClaimAuthorization {
    const message = { ...baseClaimAuthorization } as ClaimAuthorization;
    if (
      object.contract_address !== undefined &&
      object.contract_address !== null
    ) {
      message.contract_address = object.contract_address;
    } else {
      message.contract_address = "";
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action;
    } else {
      message.action = 0;
    }
    return message;
  },
};

const baseParams: object = { airdrop_enabled: false, claim_denom: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.airdrop_enabled === true) {
      writer.uint32(8).bool(message.airdrop_enabled);
    }
    if (message.airdrop_start_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.airdrop_start_time),
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.duration_until_decay !== undefined) {
      Duration.encode(
        message.duration_until_decay,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.duration_of_decay !== undefined) {
      Duration.encode(
        message.duration_of_decay,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.claim_denom !== "") {
      writer.uint32(42).string(message.claim_denom);
    }
    for (const v of message.allowed_claimers) {
      ClaimAuthorization.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.allowed_claimers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdrop_enabled = reader.bool();
          break;
        case 2:
          message.airdrop_start_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.duration_until_decay = Duration.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          message.duration_of_decay = Duration.decode(reader, reader.uint32());
          break;
        case 5:
          message.claim_denom = reader.string();
          break;
        case 6:
          message.allowed_claimers.push(
            ClaimAuthorization.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    message.allowed_claimers = [];
    if (
      object.airdrop_enabled !== undefined &&
      object.airdrop_enabled !== null
    ) {
      message.airdrop_enabled = Boolean(object.airdrop_enabled);
    } else {
      message.airdrop_enabled = false;
    }
    if (
      object.airdrop_start_time !== undefined &&
      object.airdrop_start_time !== null
    ) {
      message.airdrop_start_time = fromJsonTimestamp(object.airdrop_start_time);
    } else {
      message.airdrop_start_time = undefined;
    }
    if (
      object.duration_until_decay !== undefined &&
      object.duration_until_decay !== null
    ) {
      message.duration_until_decay = Duration.fromJSON(
        object.duration_until_decay
      );
    } else {
      message.duration_until_decay = undefined;
    }
    if (
      object.duration_of_decay !== undefined &&
      object.duration_of_decay !== null
    ) {
      message.duration_of_decay = Duration.fromJSON(object.duration_of_decay);
    } else {
      message.duration_of_decay = undefined;
    }
    if (object.claim_denom !== undefined && object.claim_denom !== null) {
      message.claim_denom = String(object.claim_denom);
    } else {
      message.claim_denom = "";
    }
    if (
      object.allowed_claimers !== undefined &&
      object.allowed_claimers !== null
    ) {
      for (const e of object.allowed_claimers) {
        message.allowed_claimers.push(ClaimAuthorization.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.airdrop_enabled !== undefined &&
      (obj.airdrop_enabled = message.airdrop_enabled);
    message.airdrop_start_time !== undefined &&
      (obj.airdrop_start_time =
        message.airdrop_start_time !== undefined
          ? message.airdrop_start_time.toISOString()
          : null);
    message.duration_until_decay !== undefined &&
      (obj.duration_until_decay = message.duration_until_decay
        ? Duration.toJSON(message.duration_until_decay)
        : undefined);
    message.duration_of_decay !== undefined &&
      (obj.duration_of_decay = message.duration_of_decay
        ? Duration.toJSON(message.duration_of_decay)
        : undefined);
    message.claim_denom !== undefined &&
      (obj.claim_denom = message.claim_denom);
    if (message.allowed_claimers) {
      obj.allowed_claimers = message.allowed_claimers.map((e) =>
        e ? ClaimAuthorization.toJSON(e) : undefined
      );
    } else {
      obj.allowed_claimers = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.allowed_claimers = [];
    if (
      object.airdrop_enabled !== undefined &&
      object.airdrop_enabled !== null
    ) {
      message.airdrop_enabled = object.airdrop_enabled;
    } else {
      message.airdrop_enabled = false;
    }
    if (
      object.airdrop_start_time !== undefined &&
      object.airdrop_start_time !== null
    ) {
      message.airdrop_start_time = object.airdrop_start_time;
    } else {
      message.airdrop_start_time = undefined;
    }
    if (
      object.duration_until_decay !== undefined &&
      object.duration_until_decay !== null
    ) {
      message.duration_until_decay = Duration.fromPartial(
        object.duration_until_decay
      );
    } else {
      message.duration_until_decay = undefined;
    }
    if (
      object.duration_of_decay !== undefined &&
      object.duration_of_decay !== null
    ) {
      message.duration_of_decay = Duration.fromPartial(
        object.duration_of_decay
      );
    } else {
      message.duration_of_decay = undefined;
    }
    if (object.claim_denom !== undefined && object.claim_denom !== null) {
      message.claim_denom = object.claim_denom;
    } else {
      message.claim_denom = "";
    }
    if (
      object.allowed_claimers !== undefined &&
      object.allowed_claimers !== null
    ) {
      for (const e of object.allowed_claimers) {
        message.allowed_claimers.push(ClaimAuthorization.fromPartial(e));
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}
