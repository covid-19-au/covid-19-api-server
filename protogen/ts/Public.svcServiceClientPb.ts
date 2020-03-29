/**
 * @fileoverview gRPC-Web generated client stub for covid19api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';
import * as case_def_pb from './case.def_pb';
import * as flight_def_pb from './flight.def_pb';

export class PublicQueryClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoIsServiceReady = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_wrappers_pb.BoolValue,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    google_protobuf_wrappers_pb.BoolValue.deserializeBinary
  );

  isServiceReady(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.PublicQuery/IsServiceReady',
      request,
      metadata || {},
      this.methodInfoIsServiceReady,
      callback);
  }

  methodInfoGetCases = new grpcWeb.AbstractClientBase.MethodInfo(
    case_def_pb.GetCasesResponse,
    (request: case_def_pb.GetCasesRequest) => {
      return request.serializeBinary();
    },
    case_def_pb.GetCasesResponse.deserializeBinary
  );

  getCases(
    request: case_def_pb.GetCasesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: case_def_pb.GetCasesResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.PublicQuery/GetCases',
      request,
      metadata || {},
      this.methodInfoGetCases,
      callback);
  }

  methodInfoGetFlights = new grpcWeb.AbstractClientBase.MethodInfo(
    flight_def_pb.GetFlightsResponse,
    (request: flight_def_pb.GetFlightsRequest) => {
      return request.serializeBinary();
    },
    flight_def_pb.GetFlightsResponse.deserializeBinary
  );

  getFlights(
    request: flight_def_pb.GetFlightsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: flight_def_pb.GetFlightsResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.PublicQuery/GetFlights',
      request,
      metadata || {},
      this.methodInfoGetFlights,
      callback);
  }

}

