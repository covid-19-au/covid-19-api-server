// source: case.def.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.covid19api.GetCaseStatsResponse');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.covid19api.GetCaseStatsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.covid19api.GetCaseStatsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.covid19api.GetCaseStatsResponse.displayName = 'proto.covid19api.GetCaseStatsResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.covid19api.GetCaseStatsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.covid19api.GetCaseStatsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.covid19api.GetCaseStatsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.covid19api.GetCaseStatsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    lastUpdTime: jspb.Message.getFieldWithDefault(msg, 1, 0),
    confirmedCount: jspb.Message.getFieldWithDefault(msg, 11, 0),
    confirmedDiffYesterday: jspb.Message.getFieldWithDefault(msg, 12, 0),
    deathsCount: jspb.Message.getFieldWithDefault(msg, 21, 0),
    deathsDiffYesterday: jspb.Message.getFieldWithDefault(msg, 22, 0),
    recoveredCount: jspb.Message.getFieldWithDefault(msg, 31, 0),
    recoveredDiffYesterday: jspb.Message.getFieldWithDefault(msg, 32, 0),
    testedCount: jspb.Message.getFieldWithDefault(msg, 41, 0),
    testedDiffYesterday: jspb.Message.getFieldWithDefault(msg, 42, 0),
    inHospitalCount: jspb.Message.getFieldWithDefault(msg, 51, 0),
    inHospitalDiffYesterday: jspb.Message.getFieldWithDefault(msg, 52, 0),
    inIcuCount: jspb.Message.getFieldWithDefault(msg, 61, 0),
    inIcuDiffYesterday: jspb.Message.getFieldWithDefault(msg, 62, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.covid19api.GetCaseStatsResponse}
 */
proto.covid19api.GetCaseStatsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.covid19api.GetCaseStatsResponse;
  return proto.covid19api.GetCaseStatsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.covid19api.GetCaseStatsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.covid19api.GetCaseStatsResponse}
 */
proto.covid19api.GetCaseStatsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLastUpdTime(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setConfirmedCount(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setConfirmedDiffYesterday(value);
      break;
    case 21:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDeathsCount(value);
      break;
    case 22:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDeathsDiffYesterday(value);
      break;
    case 31:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setRecoveredCount(value);
      break;
    case 32:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setRecoveredDiffYesterday(value);
      break;
    case 41:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTestedCount(value);
      break;
    case 42:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTestedDiffYesterday(value);
      break;
    case 51:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setInHospitalCount(value);
      break;
    case 52:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setInHospitalDiffYesterday(value);
      break;
    case 61:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setInIcuCount(value);
      break;
    case 62:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setInIcuDiffYesterday(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.covid19api.GetCaseStatsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.covid19api.GetCaseStatsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.covid19api.GetCaseStatsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.covid19api.GetCaseStatsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLastUpdTime();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getConfirmedCount();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
  f = message.getConfirmedDiffYesterday();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
  f = message.getDeathsCount();
  if (f !== 0) {
    writer.writeInt64(
      21,
      f
    );
  }
  f = message.getDeathsDiffYesterday();
  if (f !== 0) {
    writer.writeInt64(
      22,
      f
    );
  }
  f = message.getRecoveredCount();
  if (f !== 0) {
    writer.writeInt64(
      31,
      f
    );
  }
  f = message.getRecoveredDiffYesterday();
  if (f !== 0) {
    writer.writeInt64(
      32,
      f
    );
  }
  f = message.getTestedCount();
  if (f !== 0) {
    writer.writeInt64(
      41,
      f
    );
  }
  f = message.getTestedDiffYesterday();
  if (f !== 0) {
    writer.writeInt64(
      42,
      f
    );
  }
  f = message.getInHospitalCount();
  if (f !== 0) {
    writer.writeInt64(
      51,
      f
    );
  }
  f = message.getInHospitalDiffYesterday();
  if (f !== 0) {
    writer.writeInt64(
      52,
      f
    );
  }
  f = message.getInIcuCount();
  if (f !== 0) {
    writer.writeInt64(
      61,
      f
    );
  }
  f = message.getInIcuDiffYesterday();
  if (f !== 0) {
    writer.writeInt64(
      62,
      f
    );
  }
};


/**
 * optional int64 last_upd_time = 1;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getLastUpdTime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setLastUpdTime = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 confirmed_count = 11;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getConfirmedCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setConfirmedCount = function(value) {
  return jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int64 confirmed_diff_yesterday = 12;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getConfirmedDiffYesterday = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setConfirmedDiffYesterday = function(value) {
  return jspb.Message.setProto3IntField(this, 12, value);
};


/**
 * optional int64 deaths_count = 21;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getDeathsCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 21, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setDeathsCount = function(value) {
  return jspb.Message.setProto3IntField(this, 21, value);
};


/**
 * optional int64 deaths_diff_yesterday = 22;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getDeathsDiffYesterday = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 22, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setDeathsDiffYesterday = function(value) {
  return jspb.Message.setProto3IntField(this, 22, value);
};


/**
 * optional int64 recovered_count = 31;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getRecoveredCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 31, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setRecoveredCount = function(value) {
  return jspb.Message.setProto3IntField(this, 31, value);
};


/**
 * optional int64 recovered_diff_yesterday = 32;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getRecoveredDiffYesterday = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 32, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setRecoveredDiffYesterday = function(value) {
  return jspb.Message.setProto3IntField(this, 32, value);
};


/**
 * optional int64 tested_count = 41;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getTestedCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 41, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setTestedCount = function(value) {
  return jspb.Message.setProto3IntField(this, 41, value);
};


/**
 * optional int64 tested_diff_yesterday = 42;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getTestedDiffYesterday = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 42, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setTestedDiffYesterday = function(value) {
  return jspb.Message.setProto3IntField(this, 42, value);
};


/**
 * optional int64 in_hospital_count = 51;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getInHospitalCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 51, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setInHospitalCount = function(value) {
  return jspb.Message.setProto3IntField(this, 51, value);
};


/**
 * optional int64 in_hospital_diff_yesterday = 52;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getInHospitalDiffYesterday = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 52, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setInHospitalDiffYesterday = function(value) {
  return jspb.Message.setProto3IntField(this, 52, value);
};


/**
 * optional int64 in_icu_count = 61;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getInIcuCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 61, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setInIcuCount = function(value) {
  return jspb.Message.setProto3IntField(this, 61, value);
};


/**
 * optional int64 in_icu_diff_yesterday = 62;
 * @return {number}
 */
proto.covid19api.GetCaseStatsResponse.prototype.getInIcuDiffYesterday = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 62, 0));
};


/**
 * @param {number} value
 * @return {!proto.covid19api.GetCaseStatsResponse} returns this
 */
proto.covid19api.GetCaseStatsResponse.prototype.setInIcuDiffYesterday = function(value) {
  return jspb.Message.setProto3IntField(this, 62, value);
};

