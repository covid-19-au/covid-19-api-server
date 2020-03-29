import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as case_def_pb from './case.def_pb';
import * as flight_def_pb from './flight.def_pb';

export class PublicQueryClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  isServiceReady(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.BoolValue>;

  getCases(
    request: case_def_pb.GetCasesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: case_def_pb.GetCasesResponse) => void
  ): grpcWeb.ClientReadableStream<case_def_pb.GetCasesResponse>;

  getFlights(
    request: flight_def_pb.GetFlightsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: flight_def_pb.GetFlightsResponse) => void
  ): grpcWeb.ClientReadableStream<flight_def_pb.GetFlightsResponse>;

}

export class PublicQueryPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  isServiceReady(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.BoolValue>;

  getCases(
    request: case_def_pb.GetCasesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<case_def_pb.GetCasesResponse>;

  getFlights(
    request: flight_def_pb.GetFlightsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<flight_def_pb.GetFlightsResponse>;

}

