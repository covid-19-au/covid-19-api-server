/**
 * @fileoverview gRPC-Web generated client stub for covid19api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js')

var case_def_pb = require('./case.def_pb.js')

var flight_def_pb = require('./flight.def_pb.js')
const proto = {};
proto.covid19api = require('./public.svc_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.covid19api.PublicQueryClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.covid19api.PublicQueryPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodDescriptor_PublicQuery_IsServiceReady = new grpc.web.MethodDescriptor(
  '/covid19api.PublicQuery/IsServiceReady',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_PublicQuery_IsServiceReady = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.BoolValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.BoolValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.PublicQueryClient.prototype.isServiceReady =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.PublicQuery/IsServiceReady',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_IsServiceReady,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.BoolValue>}
 *     A native promise that resolves to the response
 */
proto.covid19api.PublicQueryPromiseClient.prototype.isServiceReady =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.PublicQuery/IsServiceReady',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_IsServiceReady);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.GetCasesRequest,
 *   !proto.covid19api.GetCasesResponse>}
 */
const methodDescriptor_PublicQuery_GetCases = new grpc.web.MethodDescriptor(
  '/covid19api.PublicQuery/GetCases',
  grpc.web.MethodType.UNARY,
  case_def_pb.GetCasesRequest,
  case_def_pb.GetCasesResponse,
  /**
   * @param {!proto.covid19api.GetCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  case_def_pb.GetCasesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.covid19api.GetCasesRequest,
 *   !proto.covid19api.GetCasesResponse>}
 */
const methodInfo_PublicQuery_GetCases = new grpc.web.AbstractClientBase.MethodInfo(
  case_def_pb.GetCasesResponse,
  /**
   * @param {!proto.covid19api.GetCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  case_def_pb.GetCasesResponse.deserializeBinary
);


/**
 * @param {!proto.covid19api.GetCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.covid19api.GetCasesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.covid19api.GetCasesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.PublicQueryClient.prototype.getCases =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.PublicQuery/GetCases',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_GetCases,
      callback);
};


/**
 * @param {!proto.covid19api.GetCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.covid19api.GetCasesResponse>}
 *     A native promise that resolves to the response
 */
proto.covid19api.PublicQueryPromiseClient.prototype.getCases =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.PublicQuery/GetCases',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_GetCases);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.GetFlightsRequest,
 *   !proto.covid19api.GetFlightsResponse>}
 */
const methodDescriptor_PublicQuery_GetFlights = new grpc.web.MethodDescriptor(
  '/covid19api.PublicQuery/GetFlights',
  grpc.web.MethodType.UNARY,
  flight_def_pb.GetFlightsRequest,
  flight_def_pb.GetFlightsResponse,
  /**
   * @param {!proto.covid19api.GetFlightsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  flight_def_pb.GetFlightsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.covid19api.GetFlightsRequest,
 *   !proto.covid19api.GetFlightsResponse>}
 */
const methodInfo_PublicQuery_GetFlights = new grpc.web.AbstractClientBase.MethodInfo(
  flight_def_pb.GetFlightsResponse,
  /**
   * @param {!proto.covid19api.GetFlightsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  flight_def_pb.GetFlightsResponse.deserializeBinary
);


/**
 * @param {!proto.covid19api.GetFlightsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.covid19api.GetFlightsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.covid19api.GetFlightsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.PublicQueryClient.prototype.getFlights =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.PublicQuery/GetFlights',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_GetFlights,
      callback);
};


/**
 * @param {!proto.covid19api.GetFlightsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.covid19api.GetFlightsResponse>}
 *     A native promise that resolves to the response
 */
proto.covid19api.PublicQueryPromiseClient.prototype.getFlights =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.PublicQuery/GetFlights',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_GetFlights);
};


module.exports = proto.covid19api;

