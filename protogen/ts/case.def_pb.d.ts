import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as common_pb from './common_pb';
import * as geo_def_pb from './geo.def_pb';

export class ExistingCaseDetail extends jspb.Message {
  getCaseUuid(): string;
  setCaseUuid(value: string): void;

  getPatientId(): string;
  setPatientId(value: string): void;

  getReportedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReportedTime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasReportedTime(): boolean;
  clearReportedTime(): void;

  getState(): CaseState;
  setState(value: CaseState): void;

  getInfectSrc(): InfectionSource;
  setInfectSrc(value: InfectionSource): void;

  getNamedLocation(): geo_def_pb.NamedLocation | undefined;
  setNamedLocation(value?: geo_def_pb.NamedLocation): void;
  hasNamedLocation(): boolean;
  clearNamedLocation(): void;

  getGpsLocation(): geo_def_pb.GPSLocation | undefined;
  setGpsLocation(value?: geo_def_pb.GPSLocation): void;
  hasGpsLocation(): boolean;
  clearGpsLocation(): void;

  getLocationCase(): ExistingCaseDetail.LocationCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistingCaseDetail.AsObject;
  static toObject(includeInstance: boolean, msg: ExistingCaseDetail): ExistingCaseDetail.AsObject;
  static serializeBinaryToWriter(message: ExistingCaseDetail, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistingCaseDetail;
  static deserializeBinaryFromReader(message: ExistingCaseDetail, reader: jspb.BinaryReader): ExistingCaseDetail;
}

export namespace ExistingCaseDetail {
  export type AsObject = {
    caseUuid: string,
    patientId: string,
    reportedTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    state: CaseState,
    infectSrc: InfectionSource,
    namedLocation?: geo_def_pb.NamedLocation.AsObject,
    gpsLocation?: geo_def_pb.GPSLocation.AsObject,
  }

  export enum LocationCase { 
    LOCATION_NOT_SET = 0,
    NAMED_LOCATION = 21,
    GPS_LOCATION = 22,
  }
}

export class NewCaseDetail extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getReportedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReportedTime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasReportedTime(): boolean;
  clearReportedTime(): void;

  getState(): CaseState;
  setState(value: CaseState): void;

  getInfectSrc(): InfectionSource;
  setInfectSrc(value: InfectionSource): void;

  getNamedLocation(): geo_def_pb.NamedLocation | undefined;
  setNamedLocation(value?: geo_def_pb.NamedLocation): void;
  hasNamedLocation(): boolean;
  clearNamedLocation(): void;

  getGpsLocation(): geo_def_pb.GPSLocation | undefined;
  setGpsLocation(value?: geo_def_pb.GPSLocation): void;
  hasGpsLocation(): boolean;
  clearGpsLocation(): void;

  getLocationCase(): NewCaseDetail.LocationCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewCaseDetail.AsObject;
  static toObject(includeInstance: boolean, msg: NewCaseDetail): NewCaseDetail.AsObject;
  static serializeBinaryToWriter(message: NewCaseDetail, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewCaseDetail;
  static deserializeBinaryFromReader(message: NewCaseDetail, reader: jspb.BinaryReader): NewCaseDetail;
}

export namespace NewCaseDetail {
  export type AsObject = {
    patientId: string,
    reportedTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    state: CaseState,
    infectSrc: InfectionSource,
    namedLocation?: geo_def_pb.NamedLocation.AsObject,
    gpsLocation?: geo_def_pb.GPSLocation.AsObject,
  }

  export enum LocationCase { 
    LOCATION_NOT_SET = 0,
    NAMED_LOCATION = 21,
    GPS_LOCATION = 22,
  }
}

export class GetCasesRequest extends jspb.Message {
  getNamedLocation(): geo_def_pb.NamedLocation | undefined;
  setNamedLocation(value?: geo_def_pb.NamedLocation): void;
  hasNamedLocation(): boolean;
  clearNamedLocation(): void;

  getGpsLocation(): geo_def_pb.GPSLocation | undefined;
  setGpsLocation(value?: geo_def_pb.GPSLocation): void;
  hasGpsLocation(): boolean;
  clearGpsLocation(): void;

  getStatesList(): Array<CaseState>;
  setStatesList(value: Array<CaseState>): void;
  clearStatesList(): void;
  addStates(value: CaseState, index?: number): void;

  getStartTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartTime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasStartTime(): boolean;
  clearStartTime(): void;

  getEndTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEndTime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasEndTime(): boolean;
  clearEndTime(): void;

  getLocationCase(): GetCasesRequest.LocationCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCasesRequest): GetCasesRequest.AsObject;
  static serializeBinaryToWriter(message: GetCasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCasesRequest;
  static deserializeBinaryFromReader(message: GetCasesRequest, reader: jspb.BinaryReader): GetCasesRequest;
}

export namespace GetCasesRequest {
  export type AsObject = {
    namedLocation?: geo_def_pb.NamedLocation.AsObject,
    gpsLocation?: geo_def_pb.GPSLocation.AsObject,
    statesList: Array<CaseState>,
    startTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    endTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export enum LocationCase { 
    LOCATION_NOT_SET = 0,
    NAMED_LOCATION = 1,
    GPS_LOCATION = 2,
  }
}

export class GetCasesResponse extends jspb.Message {
  getPagination(): common_pb.Pagination | undefined;
  setPagination(value?: common_pb.Pagination): void;
  hasPagination(): boolean;
  clearPagination(): void;

  getCasesList(): Array<ExistingCaseDetail>;
  setCasesList(value: Array<ExistingCaseDetail>): void;
  clearCasesList(): void;
  addCases(value?: ExistingCaseDetail, index?: number): ExistingCaseDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCasesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCasesResponse): GetCasesResponse.AsObject;
  static serializeBinaryToWriter(message: GetCasesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCasesResponse;
  static deserializeBinaryFromReader(message: GetCasesResponse, reader: jspb.BinaryReader): GetCasesResponse;
}

export namespace GetCasesResponse {
  export type AsObject = {
    pagination?: common_pb.Pagination.AsObject,
    casesList: Array<ExistingCaseDetail.AsObject>,
  }
}

export class AddCasesRequest extends jspb.Message {
  getCasesList(): Array<NewCaseDetail>;
  setCasesList(value: Array<NewCaseDetail>): void;
  clearCasesList(): void;
  addCases(value?: NewCaseDetail, index?: number): NewCaseDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddCasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddCasesRequest): AddCasesRequest.AsObject;
  static serializeBinaryToWriter(message: AddCasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddCasesRequest;
  static deserializeBinaryFromReader(message: AddCasesRequest, reader: jspb.BinaryReader): AddCasesRequest;
}

export namespace AddCasesRequest {
  export type AsObject = {
    casesList: Array<NewCaseDetail.AsObject>,
  }
}

export class PutCasesRequest extends jspb.Message {
  getCasesList(): Array<ExistingCaseDetail>;
  setCasesList(value: Array<ExistingCaseDetail>): void;
  clearCasesList(): void;
  addCases(value?: ExistingCaseDetail, index?: number): ExistingCaseDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PutCasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PutCasesRequest): PutCasesRequest.AsObject;
  static serializeBinaryToWriter(message: PutCasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PutCasesRequest;
  static deserializeBinaryFromReader(message: PutCasesRequest, reader: jspb.BinaryReader): PutCasesRequest;
}

export namespace PutCasesRequest {
  export type AsObject = {
    casesList: Array<ExistingCaseDetail.AsObject>,
  }
}

export class DelCasesRequest extends jspb.Message {
  getCaseUuidsList(): Array<string>;
  setCaseUuidsList(value: Array<string>): void;
  clearCaseUuidsList(): void;
  addCaseUuids(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DelCasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DelCasesRequest): DelCasesRequest.AsObject;
  static serializeBinaryToWriter(message: DelCasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DelCasesRequest;
  static deserializeBinaryFromReader(message: DelCasesRequest, reader: jspb.BinaryReader): DelCasesRequest;
}

export namespace DelCasesRequest {
  export type AsObject = {
    caseUuidsList: Array<string>,
  }
}

export enum CaseState { 
  _UNASSIGNED_CASE_STATE = 0,
  UNKNOWN = 11,
  UNDISCLOSED = 12,
  CONFIRMED = 21,
  RECOVERED = 22,
  DEATH = 23,
}
export enum InfectionSource { 
  _UNASSIGNED_INF_SRC = 0,
  DOMESTIC = 1,
  INTL_IMPORTED = 2,
  SHIP_IMPORTED = 3,
}
