/**
 * @fileoverview gRPC-Web generated client stub for covid19api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


goog.provide('proto.covid19api.PublicQueryClient');
goog.provide('proto.covid19api.PublicQueryPromiseClient');

goog.require('grpc.web.GrpcWebClientBase');
goog.require('grpc.web.AbstractClientBase');
goog.require('grpc.web.ClientReadableStream');
goog.require('grpc.web.Error');
goog.require('grpc.web.MethodDescriptor');
goog.require('grpc.web.MethodType');
goog.require('proto.covid19api.GetCaseStatsRequest');
goog.require('proto.covid19api.GetCaseStatsResponse');
goog.require('proto.covid19api.GetCasesRequest');
goog.require('proto.covid19api.GetCasesResponse');
goog.require('proto.covid19api.GetFlightsRequest');
goog.require('proto.covid19api.GetFlightsResponse');
goog.require('proto.google.protobuf.BoolValue');
goog.require('proto.google.protobuf.Empty');



goog.scope(function() {

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
  proto.google.protobuf.Empty,
  proto.google.protobuf.BoolValue,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.google.protobuf.BoolValue.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_PublicQuery_IsServiceReady = new grpc.web.AbstractClientBase.MethodInfo(
  proto.google.protobuf.BoolValue,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.google.protobuf.BoolValue.deserializeBinary
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
  proto.covid19api.GetCasesRequest,
  proto.covid19api.GetCasesResponse,
  /**
   * @param {!proto.covid19api.GetCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.covid19api.GetCasesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.covid19api.GetCasesRequest,
 *   !proto.covid19api.GetCasesResponse>}
 */
const methodInfo_PublicQuery_GetCases = new grpc.web.AbstractClientBase.MethodInfo(
  proto.covid19api.GetCasesResponse,
  /**
   * @param {!proto.covid19api.GetCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.covid19api.GetCasesResponse.deserializeBinary
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
 *   !proto.covid19api.GetCaseStatsRequest,
 *   !proto.covid19api.GetCaseStatsResponse>}
 */
const methodDescriptor_PublicQuery_GetCaseStats = new grpc.web.MethodDescriptor(
  '/covid19api.PublicQuery/GetCaseStats',
  grpc.web.MethodType.UNARY,
  proto.covid19api.GetCaseStatsRequest,
  proto.covid19api.GetCaseStatsResponse,
  /**
   * @param {!proto.covid19api.GetCaseStatsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.covid19api.GetCaseStatsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.covid19api.GetCaseStatsRequest,
 *   !proto.covid19api.GetCaseStatsResponse>}
 */
const methodInfo_PublicQuery_GetCaseStats = new grpc.web.AbstractClientBase.MethodInfo(
  proto.covid19api.GetCaseStatsResponse,
  /**
   * @param {!proto.covid19api.GetCaseStatsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.covid19api.GetCaseStatsResponse.deserializeBinary
);


/**
 * @param {!proto.covid19api.GetCaseStatsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.covid19api.GetCaseStatsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.covid19api.GetCaseStatsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.PublicQueryClient.prototype.getCaseStats =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.PublicQuery/GetCaseStats',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_GetCaseStats,
      callback);
};


/**
 * @param {!proto.covid19api.GetCaseStatsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.covid19api.GetCaseStatsResponse>}
 *     A native promise that resolves to the response
 */
proto.covid19api.PublicQueryPromiseClient.prototype.getCaseStats =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.PublicQuery/GetCaseStats',
      request,
      metadata || {},
      methodDescriptor_PublicQuery_GetCaseStats);
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
  proto.covid19api.GetFlightsRequest,
  proto.covid19api.GetFlightsResponse,
  /**
   * @param {!proto.covid19api.GetFlightsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.covid19api.GetFlightsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.covid19api.GetFlightsRequest,
 *   !proto.covid19api.GetFlightsResponse>}
 */
const methodInfo_PublicQuery_GetFlights = new grpc.web.AbstractClientBase.MethodInfo(
  proto.covid19api.GetFlightsResponse,
  /**
   * @param {!proto.covid19api.GetFlightsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.covid19api.GetFlightsResponse.deserializeBinary
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


}); // goog.scope

