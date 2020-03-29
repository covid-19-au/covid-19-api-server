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
proto.covid19api = require('./admin.svc_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.covid19api.DataAdminClient =
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
proto.covid19api.DataAdminPromiseClient =
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
const methodDescriptor_DataAdmin_IsServiceReady = new grpc.web.MethodDescriptor(
  '/covid19api.DataAdmin/IsServiceReady',
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
const methodInfo_DataAdmin_IsServiceReady = new grpc.web.AbstractClientBase.MethodInfo(
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
proto.covid19api.DataAdminClient.prototype.isServiceReady =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.DataAdmin/IsServiceReady',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_IsServiceReady,
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
proto.covid19api.DataAdminPromiseClient.prototype.isServiceReady =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.DataAdmin/IsServiceReady',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_IsServiceReady);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.AddCasesRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodDescriptor_DataAdmin_AddCases = new grpc.web.MethodDescriptor(
  '/covid19api.DataAdmin/AddCases',
  grpc.web.MethodType.UNARY,
  case_def_pb.AddCasesRequest,
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.AddCasesRequest} request
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
 *   !proto.covid19api.AddCasesRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_DataAdmin_AddCases = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.AddCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @param {!proto.covid19api.AddCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.BoolValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.BoolValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.DataAdminClient.prototype.addCases =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.DataAdmin/AddCases',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_AddCases,
      callback);
};


/**
 * @param {!proto.covid19api.AddCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.BoolValue>}
 *     A native promise that resolves to the response
 */
proto.covid19api.DataAdminPromiseClient.prototype.addCases =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.DataAdmin/AddCases',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_AddCases);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.PutCasesRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodDescriptor_DataAdmin_PutCases = new grpc.web.MethodDescriptor(
  '/covid19api.DataAdmin/PutCases',
  grpc.web.MethodType.UNARY,
  case_def_pb.PutCasesRequest,
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.PutCasesRequest} request
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
 *   !proto.covid19api.PutCasesRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_DataAdmin_PutCases = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.PutCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @param {!proto.covid19api.PutCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.BoolValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.BoolValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.DataAdminClient.prototype.putCases =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.DataAdmin/PutCases',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_PutCases,
      callback);
};


/**
 * @param {!proto.covid19api.PutCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.BoolValue>}
 *     A native promise that resolves to the response
 */
proto.covid19api.DataAdminPromiseClient.prototype.putCases =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.DataAdmin/PutCases',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_PutCases);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.DelCasesRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodDescriptor_DataAdmin_DelCases = new grpc.web.MethodDescriptor(
  '/covid19api.DataAdmin/DelCases',
  grpc.web.MethodType.UNARY,
  case_def_pb.DelCasesRequest,
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.DelCasesRequest} request
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
 *   !proto.covid19api.DelCasesRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_DataAdmin_DelCases = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.DelCasesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @param {!proto.covid19api.DelCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.BoolValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.BoolValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.DataAdminClient.prototype.delCases =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.DataAdmin/DelCases',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_DelCases,
      callback);
};


/**
 * @param {!proto.covid19api.DelCasesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.BoolValue>}
 *     A native promise that resolves to the response
 */
proto.covid19api.DataAdminPromiseClient.prototype.delCases =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.DataAdmin/DelCases',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_DelCases);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.AddFlightsRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodDescriptor_DataAdmin_AddFlights = new grpc.web.MethodDescriptor(
  '/covid19api.DataAdmin/AddFlights',
  grpc.web.MethodType.UNARY,
  flight_def_pb.AddFlightsRequest,
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.AddFlightsRequest} request
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
 *   !proto.covid19api.AddFlightsRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_DataAdmin_AddFlights = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.AddFlightsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @param {!proto.covid19api.AddFlightsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.BoolValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.BoolValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.DataAdminClient.prototype.addFlights =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.DataAdmin/AddFlights',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_AddFlights,
      callback);
};


/**
 * @param {!proto.covid19api.AddFlightsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.BoolValue>}
 *     A native promise that resolves to the response
 */
proto.covid19api.DataAdminPromiseClient.prototype.addFlights =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.DataAdmin/AddFlights',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_AddFlights);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.covid19api.DelFlightsRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodDescriptor_DataAdmin_DelFlights = new grpc.web.MethodDescriptor(
  '/covid19api.DataAdmin/DelFlights',
  grpc.web.MethodType.UNARY,
  flight_def_pb.DelFlightsRequest,
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.DelFlightsRequest} request
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
 *   !proto.covid19api.DelFlightsRequest,
 *   !proto.google.protobuf.BoolValue>}
 */
const methodInfo_DataAdmin_DelFlights = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_wrappers_pb.BoolValue,
  /**
   * @param {!proto.covid19api.DelFlightsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_wrappers_pb.BoolValue.deserializeBinary
);


/**
 * @param {!proto.covid19api.DelFlightsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.BoolValue)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.BoolValue>|undefined}
 *     The XHR Node Readable Stream
 */
proto.covid19api.DataAdminClient.prototype.delFlights =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/covid19api.DataAdmin/DelFlights',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_DelFlights,
      callback);
};


/**
 * @param {!proto.covid19api.DelFlightsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.BoolValue>}
 *     A native promise that resolves to the response
 */
proto.covid19api.DataAdminPromiseClient.prototype.delFlights =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/covid19api.DataAdmin/DelFlights',
      request,
      metadata || {},
      methodDescriptor_DataAdmin_DelFlights);
};


module.exports = proto.covid19api;

