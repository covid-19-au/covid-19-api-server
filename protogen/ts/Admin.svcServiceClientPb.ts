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

export class DataAdminClient {
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
        '/covid19api.DataAdmin/IsServiceReady',
      request,
      metadata || {},
      this.methodInfoIsServiceReady,
      callback);
  }

  methodInfoAddCases = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_wrappers_pb.BoolValue,
    (request: case_def_pb.AddCasesRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_wrappers_pb.BoolValue.deserializeBinary
  );

  addCases(
    request: case_def_pb.AddCasesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.DataAdmin/AddCases',
      request,
      metadata || {},
      this.methodInfoAddCases,
      callback);
  }

  methodInfoPutCases = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_wrappers_pb.BoolValue,
    (request: case_def_pb.PutCasesRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_wrappers_pb.BoolValue.deserializeBinary
  );

  putCases(
    request: case_def_pb.PutCasesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.DataAdmin/PutCases',
      request,
      metadata || {},
      this.methodInfoPutCases,
      callback);
  }

  methodInfoDelCases = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_wrappers_pb.BoolValue,
    (request: case_def_pb.DelCasesRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_wrappers_pb.BoolValue.deserializeBinary
  );

  delCases(
    request: case_def_pb.DelCasesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.DataAdmin/DelCases',
      request,
      metadata || {},
      this.methodInfoDelCases,
      callback);
  }

  methodInfoAddFlights = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_wrappers_pb.BoolValue,
    (request: flight_def_pb.AddFlightsRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_wrappers_pb.BoolValue.deserializeBinary
  );

  addFlights(
    request: flight_def_pb.AddFlightsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.DataAdmin/AddFlights',
      request,
      metadata || {},
      this.methodInfoAddFlights,
      callback);
  }

  methodInfoDelFlights = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_wrappers_pb.BoolValue,
    (request: flight_def_pb.DelFlightsRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_wrappers_pb.BoolValue.deserializeBinary
  );

  delFlights(
    request: flight_def_pb.DelFlightsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_wrappers_pb.BoolValue) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/covid19api.DataAdmin/DelFlights',
      request,
      metadata || {},
      this.methodInfoDelFlights,
      callback);
  }

}

