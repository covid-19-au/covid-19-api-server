import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as common_pb from './common_pb';

export class FlightDetail extends jspb.Message {
  getFlightNum(): string;
  setFlightNum(value: string): void;

  getIataCode(): string;
  setIataCode(value: string): void;

  getIcaoCode(): string;
  setIcaoCode(value: string): void;

  getDepDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDepDate(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasDepDate(): boolean;
  clearDepDate(): void;

  getArrDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setArrDate(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasArrDate(): boolean;
  clearArrDate(): void;

  getDepAirport(): string;
  setDepAirport(value: string): void;

  getArrAirport(): string;
  setArrAirport(value: string): void;

  getAffectedRowsList(): Array<number>;
  setAffectedRowsList(value: Array<number>): void;
  clearAffectedRowsList(): void;
  addAffectedRows(value: number, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FlightDetail.AsObject;
  static toObject(includeInstance: boolean, msg: FlightDetail): FlightDetail.AsObject;
  static serializeBinaryToWriter(message: FlightDetail, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FlightDetail;
  static deserializeBinaryFromReader(message: FlightDetail, reader: jspb.BinaryReader): FlightDetail;
}

export namespace FlightDetail {
  export type AsObject = {
    flightNum: string,
    iataCode: string,
    icaoCode: string,
    depDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    arrDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    depAirport: string,
    arrAirport: string,
    affectedRowsList: Array<number>,
  }
}

export class GetFlightsRequest extends jspb.Message {
  getFlightNum(): string;
  setFlightNum(value: string): void;

  getDepDate(): number;
  setDepDate(value: number): void;

  getDepAirport(): string;
  setDepAirport(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFlightsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFlightsRequest): GetFlightsRequest.AsObject;
  static serializeBinaryToWriter(message: GetFlightsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFlightsRequest;
  static deserializeBinaryFromReader(message: GetFlightsRequest, reader: jspb.BinaryReader): GetFlightsRequest;
}

export namespace GetFlightsRequest {
  export type AsObject = {
    flightNum: string,
    depDate: number,
    depAirport: string,
  }
}

export class GetFlightsResponse extends jspb.Message {
  getPagination(): common_pb.Pagination | undefined;
  setPagination(value?: common_pb.Pagination): void;
  hasPagination(): boolean;
  clearPagination(): void;

  getFlightsList(): Array<FlightDetail>;
  setFlightsList(value: Array<FlightDetail>): void;
  clearFlightsList(): void;
  addFlights(value?: FlightDetail, index?: number): FlightDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFlightsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetFlightsResponse): GetFlightsResponse.AsObject;
  static serializeBinaryToWriter(message: GetFlightsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFlightsResponse;
  static deserializeBinaryFromReader(message: GetFlightsResponse, reader: jspb.BinaryReader): GetFlightsResponse;
}

export namespace GetFlightsResponse {
  export type AsObject = {
    pagination?: common_pb.Pagination.AsObject,
    flightsList: Array<FlightDetail.AsObject>,
  }
}

export class AddFlightsRequest extends jspb.Message {
  getFlightsList(): Array<FlightDetail>;
  setFlightsList(value: Array<FlightDetail>): void;
  clearFlightsList(): void;
  addFlights(value?: FlightDetail, index?: number): FlightDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddFlightsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddFlightsRequest): AddFlightsRequest.AsObject;
  static serializeBinaryToWriter(message: AddFlightsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddFlightsRequest;
  static deserializeBinaryFromReader(message: AddFlightsRequest, reader: jspb.BinaryReader): AddFlightsRequest;
}

export namespace AddFlightsRequest {
  export type AsObject = {
    flightsList: Array<FlightDetail.AsObject>,
  }
}

export class DelFlightsRequest extends jspb.Message {
  getFlightsList(): Array<DelFlightsRequest.FlightIdentifier>;
  setFlightsList(value: Array<DelFlightsRequest.FlightIdentifier>): void;
  clearFlightsList(): void;
  addFlights(value?: DelFlightsRequest.FlightIdentifier, index?: number): DelFlightsRequest.FlightIdentifier;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DelFlightsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DelFlightsRequest): DelFlightsRequest.AsObject;
  static serializeBinaryToWriter(message: DelFlightsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DelFlightsRequest;
  static deserializeBinaryFromReader(message: DelFlightsRequest, reader: jspb.BinaryReader): DelFlightsRequest;
}

export namespace DelFlightsRequest {
  export type AsObject = {
    flightsList: Array<DelFlightsRequest.FlightIdentifier.AsObject>,
  }

  export class FlightIdentifier extends jspb.Message {
    getFlightNum(): string;
    setFlightNum(value: string): void;

    getDepDate(): number;
    setDepDate(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FlightIdentifier.AsObject;
    static toObject(includeInstance: boolean, msg: FlightIdentifier): FlightIdentifier.AsObject;
    static serializeBinaryToWriter(message: FlightIdentifier, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FlightIdentifier;
    static deserializeBinaryFromReader(message: FlightIdentifier, reader: jspb.BinaryReader): FlightIdentifier;
  }

  export namespace FlightIdentifier {
    export type AsObject = {
      flightNum: string,
      depDate: number,
    }
  }

}

