// source: ps_write.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var common_ps_common_auth_pb = require('./common/ps_common_auth_pb.js');
goog.object.extend(proto, common_ps_common_auth_pb);
var common_ps_common_status_pb = require('./common/ps_common_status_pb.js');
goog.object.extend(proto, common_ps_common_status_pb);
var opts_ps_opts_write_pb = require('./opts/ps_opts_write_pb.js');
goog.object.extend(proto, opts_ps_opts_write_pb);
goog.exportSymbol('proto.protos.WriteRequest', null, global);
goog.exportSymbol('proto.protos.WriteResponse', null, global);
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
proto.protos.WriteRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, null);
};
goog.inherits(proto.protos.WriteRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.WriteRequest.displayName = 'proto.protos.WriteRequest';
}
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
proto.protos.WriteResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, null);
};
goog.inherits(proto.protos.WriteResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.WriteResponse.displayName = 'proto.protos.WriteResponse';
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
proto.protos.WriteRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.WriteRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.WriteRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.WriteRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    auth: (f = msg.getAuth()) && common_ps_common_auth_pb.Auth.toObject(includeInstance, f),
    opts: (f = msg.getOpts()) && opts_ps_opts_write_pb.WriteOptions.toObject(includeInstance, f)
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
 * @return {!proto.protos.WriteRequest}
 */
proto.protos.WriteRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.WriteRequest;
  return proto.protos.WriteRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.WriteRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.WriteRequest}
 */
proto.protos.WriteRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 9999:
      var value = new common_ps_common_auth_pb.Auth;
      reader.readMessage(value,common_ps_common_auth_pb.Auth.deserializeBinaryFromReader);
      msg.setAuth(value);
      break;
    case 1:
      var value = new opts_ps_opts_write_pb.WriteOptions;
      reader.readMessage(value,opts_ps_opts_write_pb.WriteOptions.deserializeBinaryFromReader);
      msg.setOpts(value);
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
proto.protos.WriteRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.WriteRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.WriteRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.WriteRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAuth();
  if (f != null) {
    writer.writeMessage(
      9999,
      f,
      common_ps_common_auth_pb.Auth.serializeBinaryToWriter
    );
  }
  f = message.getOpts();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      opts_ps_opts_write_pb.WriteOptions.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.Auth auth = 9999;
 * @return {?proto.protos.common.Auth}
 */
proto.protos.WriteRequest.prototype.getAuth = function() {
  return /** @type{?proto.protos.common.Auth} */ (
    jspb.Message.getWrapperField(this, common_ps_common_auth_pb.Auth, 9999));
};


/**
 * @param {?proto.protos.common.Auth|undefined} value
 * @return {!proto.protos.WriteRequest} returns this
*/
proto.protos.WriteRequest.prototype.setAuth = function(value) {
  return jspb.Message.setWrapperField(this, 9999, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.WriteRequest} returns this
 */
proto.protos.WriteRequest.prototype.clearAuth = function() {
  return this.setAuth(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.WriteRequest.prototype.hasAuth = function() {
  return jspb.Message.getField(this, 9999) != null;
};


/**
 * optional opts.WriteOptions opts = 1;
 * @return {?proto.protos.opts.WriteOptions}
 */
proto.protos.WriteRequest.prototype.getOpts = function() {
  return /** @type{?proto.protos.opts.WriteOptions} */ (
    jspb.Message.getWrapperField(this, opts_ps_opts_write_pb.WriteOptions, 1));
};


/**
 * @param {?proto.protos.opts.WriteOptions|undefined} value
 * @return {!proto.protos.WriteRequest} returns this
*/
proto.protos.WriteRequest.prototype.setOpts = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.WriteRequest} returns this
 */
proto.protos.WriteRequest.prototype.clearOpts = function() {
  return this.setOpts(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.WriteRequest.prototype.hasOpts = function() {
  return jspb.Message.getField(this, 1) != null;
};





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
proto.protos.WriteResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.WriteResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.WriteResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.WriteResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    status: (f = msg.getStatus()) && common_ps_common_status_pb.Status.toObject(includeInstance, f)
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
 * @return {!proto.protos.WriteResponse}
 */
proto.protos.WriteResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.WriteResponse;
  return proto.protos.WriteResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.WriteResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.WriteResponse}
 */
proto.protos.WriteResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1000:
      var value = new common_ps_common_status_pb.Status;
      reader.readMessage(value,common_ps_common_status_pb.Status.deserializeBinaryFromReader);
      msg.setStatus(value);
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
proto.protos.WriteResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.WriteResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.WriteResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.WriteResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStatus();
  if (f != null) {
    writer.writeMessage(
      1000,
      f,
      common_ps_common_status_pb.Status.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.Status status = 1000;
 * @return {?proto.protos.common.Status}
 */
proto.protos.WriteResponse.prototype.getStatus = function() {
  return /** @type{?proto.protos.common.Status} */ (
    jspb.Message.getWrapperField(this, common_ps_common_status_pb.Status, 1000));
};


/**
 * @param {?proto.protos.common.Status|undefined} value
 * @return {!proto.protos.WriteResponse} returns this
*/
proto.protos.WriteResponse.prototype.setStatus = function(value) {
  return jspb.Message.setWrapperField(this, 1000, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.WriteResponse} returns this
 */
proto.protos.WriteResponse.prototype.clearStatus = function() {
  return this.setStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.WriteResponse.prototype.hasStatus = function() {
  return jspb.Message.getField(this, 1000) != null;
};


goog.object.extend(exports, proto.protos);
