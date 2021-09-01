// source: backends/rabbit-streams.proto
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

goog.exportSymbol('proto.protos.backends.RabbitStreamsConn', null, global);
goog.exportSymbol('proto.protos.backends.RabbitStreamsOffsetOptions', null, global);
goog.exportSymbol('proto.protos.backends.RabbitStreamsReadArgs', null, global);
goog.exportSymbol('proto.protos.backends.RabbitStreamsWriteArgs', null, global);
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
proto.protos.backends.RabbitStreamsConn = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.backends.RabbitStreamsConn, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RabbitStreamsConn.displayName = 'proto.protos.backends.RabbitStreamsConn';
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
proto.protos.backends.RabbitStreamsOffsetOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.backends.RabbitStreamsOffsetOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RabbitStreamsOffsetOptions.displayName = 'proto.protos.backends.RabbitStreamsOffsetOptions';
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
proto.protos.backends.RabbitStreamsReadArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.backends.RabbitStreamsReadArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RabbitStreamsReadArgs.displayName = 'proto.protos.backends.RabbitStreamsReadArgs';
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
proto.protos.backends.RabbitStreamsWriteArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.backends.RabbitStreamsWriteArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RabbitStreamsWriteArgs.displayName = 'proto.protos.backends.RabbitStreamsWriteArgs';
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
proto.protos.backends.RabbitStreamsConn.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RabbitStreamsConn.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RabbitStreamsConn} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsConn.toObject = function(includeInstance, msg) {
  var f, obj = {
    dsn: jspb.Message.getFieldWithDefault(msg, 1, ""),
    useTls: jspb.Message.getBooleanFieldWithDefault(msg, 2, false),
    insecureTls: jspb.Message.getBooleanFieldWithDefault(msg, 3, false),
    username: jspb.Message.getFieldWithDefault(msg, 4, ""),
    password: jspb.Message.getFieldWithDefault(msg, 5, ""),
    clientName: jspb.Message.getFieldWithDefault(msg, 6, "")
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
 * @return {!proto.protos.backends.RabbitStreamsConn}
 */
proto.protos.backends.RabbitStreamsConn.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RabbitStreamsConn;
  return proto.protos.backends.RabbitStreamsConn.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RabbitStreamsConn} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RabbitStreamsConn}
 */
proto.protos.backends.RabbitStreamsConn.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setDsn(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUseTls(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setInsecureTls(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientName(value);
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
proto.protos.backends.RabbitStreamsConn.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RabbitStreamsConn.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RabbitStreamsConn} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsConn.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDsn();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUseTls();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getInsecureTls();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getClientName();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional string dsn = 1;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsConn.prototype.getDsn = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsConn} returns this
 */
proto.protos.backends.RabbitStreamsConn.prototype.setDsn = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool use_tls = 2;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsConn.prototype.getUseTls = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsConn} returns this
 */
proto.protos.backends.RabbitStreamsConn.prototype.setUseTls = function(value) {
  return jspb.Message.setProto3BooleanField(this, 2, value);
};


/**
 * optional bool insecure_tls = 3;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsConn.prototype.getInsecureTls = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsConn} returns this
 */
proto.protos.backends.RabbitStreamsConn.prototype.setInsecureTls = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional string username = 4;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsConn.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsConn} returns this
 */
proto.protos.backends.RabbitStreamsConn.prototype.setUsername = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string password = 5;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsConn.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsConn} returns this
 */
proto.protos.backends.RabbitStreamsConn.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string client_name = 6;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsConn.prototype.getClientName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsConn} returns this
 */
proto.protos.backends.RabbitStreamsConn.prototype.setClientName = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
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
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RabbitStreamsOffsetOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RabbitStreamsOffsetOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsOffsetOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    specificOffset: jspb.Message.getFieldWithDefault(msg, 1, 0),
    lastOffset: jspb.Message.getFieldWithDefault(msg, 2, 0),
    lastConsumed: jspb.Message.getBooleanFieldWithDefault(msg, 3, false),
    firstOffset: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    nextOffset: jspb.Message.getBooleanFieldWithDefault(msg, 5, false)
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
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RabbitStreamsOffsetOptions;
  return proto.protos.backends.RabbitStreamsOffsetOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RabbitStreamsOffsetOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setSpecificOffset(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLastOffset(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setLastConsumed(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setFirstOffset(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setNextOffset(value);
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
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RabbitStreamsOffsetOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RabbitStreamsOffsetOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsOffsetOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSpecificOffset();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getLastOffset();
  if (f !== 0) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = message.getLastConsumed();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getFirstOffset();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getNextOffset();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
};


/**
 * optional uint64 specific_offset = 1;
 * @return {number}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.getSpecificOffset = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions} returns this
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.setSpecificOffset = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional uint64 last_offset = 2;
 * @return {number}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.getLastOffset = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions} returns this
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.setLastOffset = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional bool last_consumed = 3;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.getLastConsumed = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions} returns this
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.setLastConsumed = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional bool first_offset = 4;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.getFirstOffset = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions} returns this
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.setFirstOffset = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional bool next_offset = 5;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.getNextOffset = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsOffsetOptions} returns this
 */
proto.protos.backends.RabbitStreamsOffsetOptions.prototype.setNextOffset = function(value) {
  return jspb.Message.setProto3BooleanField(this, 5, value);
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
proto.protos.backends.RabbitStreamsReadArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RabbitStreamsReadArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RabbitStreamsReadArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsReadArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    stream: jspb.Message.getFieldWithDefault(msg, 1, ""),
    declareStream: jspb.Message.getBooleanFieldWithDefault(msg, 2, false),
    declareStreamSize: jspb.Message.getFieldWithDefault(msg, 3, ""),
    offsetOptions: (f = msg.getOffsetOptions()) && proto.protos.backends.RabbitStreamsOffsetOptions.toObject(includeInstance, f)
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
 * @return {!proto.protos.backends.RabbitStreamsReadArgs}
 */
proto.protos.backends.RabbitStreamsReadArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RabbitStreamsReadArgs;
  return proto.protos.backends.RabbitStreamsReadArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RabbitStreamsReadArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RabbitStreamsReadArgs}
 */
proto.protos.backends.RabbitStreamsReadArgs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setStream(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDeclareStream(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDeclareStreamSize(value);
      break;
    case 4:
      var value = new proto.protos.backends.RabbitStreamsOffsetOptions;
      reader.readMessage(value,proto.protos.backends.RabbitStreamsOffsetOptions.deserializeBinaryFromReader);
      msg.setOffsetOptions(value);
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
proto.protos.backends.RabbitStreamsReadArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RabbitStreamsReadArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RabbitStreamsReadArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsReadArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStream();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDeclareStream();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getDeclareStreamSize();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getOffsetOptions();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.protos.backends.RabbitStreamsOffsetOptions.serializeBinaryToWriter
    );
  }
};


/**
 * optional string stream = 1;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.getStream = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsReadArgs} returns this
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.setStream = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool declare_stream = 2;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.getDeclareStream = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsReadArgs} returns this
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.setDeclareStream = function(value) {
  return jspb.Message.setProto3BooleanField(this, 2, value);
};


/**
 * optional string declare_stream_size = 3;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.getDeclareStreamSize = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsReadArgs} returns this
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.setDeclareStreamSize = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional RabbitStreamsOffsetOptions offset_options = 4;
 * @return {?proto.protos.backends.RabbitStreamsOffsetOptions}
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.getOffsetOptions = function() {
  return /** @type{?proto.protos.backends.RabbitStreamsOffsetOptions} */ (
    jspb.Message.getWrapperField(this, proto.protos.backends.RabbitStreamsOffsetOptions, 4));
};


/**
 * @param {?proto.protos.backends.RabbitStreamsOffsetOptions|undefined} value
 * @return {!proto.protos.backends.RabbitStreamsReadArgs} returns this
*/
proto.protos.backends.RabbitStreamsReadArgs.prototype.setOffsetOptions = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protos.backends.RabbitStreamsReadArgs} returns this
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.clearOffsetOptions = function() {
  return this.setOffsetOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsReadArgs.prototype.hasOffsetOptions = function() {
  return jspb.Message.getField(this, 4) != null;
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
proto.protos.backends.RabbitStreamsWriteArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RabbitStreamsWriteArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RabbitStreamsWriteArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsWriteArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    stream: jspb.Message.getFieldWithDefault(msg, 1, ""),
    declareStream: jspb.Message.getBooleanFieldWithDefault(msg, 2, false),
    declareStreamSize: jspb.Message.getFieldWithDefault(msg, 3, "")
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
 * @return {!proto.protos.backends.RabbitStreamsWriteArgs}
 */
proto.protos.backends.RabbitStreamsWriteArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RabbitStreamsWriteArgs;
  return proto.protos.backends.RabbitStreamsWriteArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RabbitStreamsWriteArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RabbitStreamsWriteArgs}
 */
proto.protos.backends.RabbitStreamsWriteArgs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setStream(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDeclareStream(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDeclareStreamSize(value);
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
proto.protos.backends.RabbitStreamsWriteArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RabbitStreamsWriteArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RabbitStreamsWriteArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RabbitStreamsWriteArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStream();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDeclareStream();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getDeclareStreamSize();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string stream = 1;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsWriteArgs.prototype.getStream = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsWriteArgs} returns this
 */
proto.protos.backends.RabbitStreamsWriteArgs.prototype.setStream = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool declare_stream = 2;
 * @return {boolean}
 */
proto.protos.backends.RabbitStreamsWriteArgs.prototype.getDeclareStream = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protos.backends.RabbitStreamsWriteArgs} returns this
 */
proto.protos.backends.RabbitStreamsWriteArgs.prototype.setDeclareStream = function(value) {
  return jspb.Message.setProto3BooleanField(this, 2, value);
};


/**
 * optional string declare_stream_size = 3;
 * @return {string}
 */
proto.protos.backends.RabbitStreamsWriteArgs.prototype.getDeclareStreamSize = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RabbitStreamsWriteArgs} returns this
 */
proto.protos.backends.RabbitStreamsWriteArgs.prototype.setDeclareStreamSize = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


goog.object.extend(exports, proto.protos.backends);