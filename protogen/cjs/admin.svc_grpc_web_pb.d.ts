import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as case_def_pb from './case.def_pb';
import * as flight_def_pb from './flight.def_pb';

export class DataAdminClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  isServiceReady(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

  addCases(
    request: case_def_pb.AddCasesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

  putCases(
    request: case_def_pb.PutCasesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

  delCases(
    request: case_def_pb.DelCasesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

  addFlights(
    request: flight_def_pb.AddFlightsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

  delFlights(
    request: flight_def_pb.DelFlightsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

}

export class DataAdminPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  isServiceReady(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

  addCases(
    request: case_def_pb.AddCasesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

  putCases(
    request: case_def_pb.PutCasesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

  delCases(
    request: case_def_pb.DelCasesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

  addFlights(
    request: flight_def_pb.AddFlightsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

  delFlights(
    request: flight_def_pb.DelFlightsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

}

