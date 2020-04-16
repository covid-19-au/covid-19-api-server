import * as jspb from "google-protobuf"

import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';
import * as common_pb from './common_pb';
import * as geo_def_pb from './geo.def_pb';

export class ExistingCaseDetail extends jspb.Message {
  getCaseId(): string;
  setCaseId(value: string): void;

  getPatientId(): string;
  setPatientId(value: string): void;

  getReportedTime(): number;
  setReportedTime(value: number): void;

  getState(): CaseState;
  setState(value: CaseState): void;

  getInfectSrc(): InfectionSource;
  setInfectSrc(value: InfectionSource): void;

  getLocation(): geo_def_pb.Location | undefined;
  setLocation(value?: geo_def_pb.Location): void;
  hasLocation(): boolean;
  clearLocation(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExistingCaseDetail.AsObject;
  static toObject(includeInstance: boolean, msg: ExistingCaseDetail): ExistingCaseDetail.AsObject;
  static serializeBinaryToWriter(message: ExistingCaseDetail, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExistingCaseDetail;
  static deserializeBinaryFromReader(message: ExistingCaseDetail, reader: jspb.BinaryReader): ExistingCaseDetail;
}

export namespace ExistingCaseDetail {
  export type AsObject = {
    caseId: string,
    patientId: string,
    reportedTime: number,
    state: CaseState,
    infectSrc: InfectionSource,
    location?: geo_def_pb.Location.AsObject,
  }
}

export class NewCaseDetail extends jspb.Message {
  getPatientId(): string;
  setPatientId(value: string): void;

  getReportedTime(): number;
  setReportedTime(value: number): void;

  getState(): CaseState;
  setState(value: CaseState): void;

  getInfectSrc(): InfectionSource;
  setInfectSrc(value: InfectionSource): void;

  getLocation(): geo_def_pb.Location | undefined;
  setLocation(value?: geo_def_pb.Location): void;
  hasLocation(): boolean;
  clearLocation(): void;

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
    reportedTime: number,
    state: CaseState,
    infectSrc: InfectionSource,
    location?: geo_def_pb.Location.AsObject,
  }
}

export class GetCasesRequest extends jspb.Message {
  getLocation(): geo_def_pb.Location | undefined;
  setLocation(value?: geo_def_pb.Location): void;
  hasLocation(): boolean;
  clearLocation(): void;

  getStatesList(): Array<CaseState>;
  setStatesList(value: Array<CaseState>): void;
  clearStatesList(): void;
  addStates(value: CaseState, index?: number): void;

  getStartTime(): number;
  setStartTime(value: number): void;

  getEndTime(): number;
  setEndTime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCasesRequest): GetCasesRequest.AsObject;
  static serializeBinaryToWriter(message: GetCasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCasesRequest;
  static deserializeBinaryFromReader(message: GetCasesRequest, reader: jspb.BinaryReader): GetCasesRequest;
}

export namespace GetCasesRequest {
  export type AsObject = {
    location?: geo_def_pb.Location.AsObject,
    statesList: Array<CaseState>,
    startTime: number,
    endTime: number,
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

export class GetCaseStatsRequest extends jspb.Message {
  getNoCache(): boolean;
  setNoCache(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCaseStatsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCaseStatsRequest): GetCaseStatsRequest.AsObject;
  static serializeBinaryToWriter(message: GetCaseStatsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCaseStatsRequest;
  static deserializeBinaryFromReader(message: GetCaseStatsRequest, reader: jspb.BinaryReader): GetCaseStatsRequest;
}

export namespace GetCaseStatsRequest {
  export type AsObject = {
    noCache: boolean,
  }
}

export class GetCaseStatsResponse extends jspb.Message {
  getLastUpdTime(): number;
  setLastUpdTime(value: number): void;

  getConfirmedCount(): number;
  setConfirmedCount(value: number): void;

  getConfirmedDiffYesterday(): number;
  setConfirmedDiffYesterday(value: number): void;

  getDeathsCount(): number;
  setDeathsCount(value: number): void;

  getDeathsDiffYesterday(): number;
  setDeathsDiffYesterday(value: number): void;

  getRecoveredCount(): number;
  setRecoveredCount(value: number): void;

  getRecoveredDiffYesterday(): number;
  setRecoveredDiffYesterday(value: number): void;

  getTestedCount(): number;
  setTestedCount(value: number): void;

  getTestedDiffYesterday(): number;
  setTestedDiffYesterday(value: number): void;

  getInHospitalCount(): number;
  setInHospitalCount(value: number): void;

  getInHospitalDiffYesterday(): number;
  setInHospitalDiffYesterday(value: number): void;

  getInIcuCount(): number;
  setInIcuCount(value: number): void;

  getInIcuDiffYesterday(): number;
  setInIcuDiffYesterday(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCaseStatsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCaseStatsResponse): GetCaseStatsResponse.AsObject;
  static serializeBinaryToWriter(message: GetCaseStatsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCaseStatsResponse;
  static deserializeBinaryFromReader(message: GetCaseStatsResponse, reader: jspb.BinaryReader): GetCaseStatsResponse;
}

export namespace GetCaseStatsResponse {
  export type AsObject = {
    lastUpdTime: number,
    confirmedCount: number,
    confirmedDiffYesterday: number,
    deathsCount: number,
    deathsDiffYesterday: number,
    recoveredCount: number,
    recoveredDiffYesterday: number,
    testedCount: number,
    testedDiffYesterday: number,
    inHospitalCount: number,
    inHospitalDiffYesterday: number,
    inIcuCount: number,
    inIcuDiffYesterday: number,
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
  getCaseIdsList(): Array<string>;
  setCaseIdsList(value: Array<string>): void;
  clearCaseIdsList(): void;
  addCaseIds(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DelCasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DelCasesRequest): DelCasesRequest.AsObject;
  static serializeBinaryToWriter(message: DelCasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DelCasesRequest;
  static deserializeBinaryFromReader(message: DelCasesRequest, reader: jspb.BinaryReader): DelCasesRequest;
}

export namespace DelCasesRequest {
  export type AsObject = {
    caseIdsList: Array<string>,
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
