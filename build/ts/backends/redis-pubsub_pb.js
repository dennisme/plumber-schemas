// source: backends/redis-pubsub.proto
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

goog.exportSymbol('proto.protos.backends.RedisPubSubConn', null, global);
goog.exportSymbol('proto.protos.backends.RedisPubSubReadArgs', null, global);
goog.exportSymbol('proto.protos.backends.RedisPubSubWriteArgs', null, global);
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
proto.protos.backends.RedisPubSubConn = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.backends.RedisPubSubConn, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RedisPubSubConn.displayName = 'proto.protos.backends.RedisPubSubConn';
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
proto.protos.backends.RedisPubSubReadArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.protos.backends.RedisPubSubReadArgs.repeatedFields_, null);
};
goog.inherits(proto.protos.backends.RedisPubSubReadArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RedisPubSubReadArgs.displayName = 'proto.protos.backends.RedisPubSubReadArgs';
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
proto.protos.backends.RedisPubSubWriteArgs = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.protos.backends.RedisPubSubWriteArgs.repeatedFields_, null);
};
goog.inherits(proto.protos.backends.RedisPubSubWriteArgs, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.backends.RedisPubSubWriteArgs.displayName = 'proto.protos.backends.RedisPubSubWriteArgs';
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
proto.protos.backends.RedisPubSubConn.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RedisPubSubConn.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RedisPubSubConn} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RedisPubSubConn.toObject = function(includeInstance, msg) {
  var f, obj = {
    address: jspb.Message.getFieldWithDefault(msg, 1, ""),
    username: jspb.Message.getFieldWithDefault(msg, 2, ""),
    password: jspb.Message.getFieldWithDefault(msg, 3, "")
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
 * @return {!proto.protos.backends.RedisPubSubConn}
 */
proto.protos.backends.RedisPubSubConn.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RedisPubSubConn;
  return proto.protos.backends.RedisPubSubConn.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RedisPubSubConn} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RedisPubSubConn}
 */
proto.protos.backends.RedisPubSubConn.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
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
proto.protos.backends.RedisPubSubConn.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RedisPubSubConn.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RedisPubSubConn} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RedisPubSubConn.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string address = 1;
 * @return {string}
 */
proto.protos.backends.RedisPubSubConn.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RedisPubSubConn} returns this
 */
proto.protos.backends.RedisPubSubConn.prototype.setAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string username = 2;
 * @return {string}
 */
proto.protos.backends.RedisPubSubConn.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RedisPubSubConn} returns this
 */
proto.protos.backends.RedisPubSubConn.prototype.setUsername = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string password = 3;
 * @return {string}
 */
proto.protos.backends.RedisPubSubConn.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.backends.RedisPubSubConn} returns this
 */
proto.protos.backends.RedisPubSubConn.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.protos.backends.RedisPubSubReadArgs.repeatedFields_ = [2];



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
proto.protos.backends.RedisPubSubReadArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RedisPubSubReadArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RedisPubSubReadArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RedisPubSubReadArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    database: jspb.Message.getFieldWithDefault(msg, 1, 0),
    channelList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
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
 * @return {!proto.protos.backends.RedisPubSubReadArgs}
 */
proto.protos.backends.RedisPubSubReadArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RedisPubSubReadArgs;
  return proto.protos.backends.RedisPubSubReadArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RedisPubSubReadArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RedisPubSubReadArgs}
 */
proto.protos.backends.RedisPubSubReadArgs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setDatabase(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addChannel(value);
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
proto.protos.backends.RedisPubSubReadArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RedisPubSubReadArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RedisPubSubReadArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RedisPubSubReadArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabase();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getChannelList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional uint32 database = 1;
 * @return {number}
 */
proto.protos.backends.RedisPubSubReadArgs.prototype.getDatabase = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.protos.backends.RedisPubSubReadArgs} returns this
 */
proto.protos.backends.RedisPubSubReadArgs.prototype.setDatabase = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated string channel = 2;
 * @return {!Array<string>}
 */
proto.protos.backends.RedisPubSubReadArgs.prototype.getChannelList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.protos.backends.RedisPubSubReadArgs} returns this
 */
proto.protos.backends.RedisPubSubReadArgs.prototype.setChannelList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.protos.backends.RedisPubSubReadArgs} returns this
 */
proto.protos.backends.RedisPubSubReadArgs.prototype.addChannel = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.protos.backends.RedisPubSubReadArgs} returns this
 */
proto.protos.backends.RedisPubSubReadArgs.prototype.clearChannelList = function() {
  return this.setChannelList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.protos.backends.RedisPubSubWriteArgs.repeatedFields_ = [2];



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
proto.protos.backends.RedisPubSubWriteArgs.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.backends.RedisPubSubWriteArgs.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.backends.RedisPubSubWriteArgs} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RedisPubSubWriteArgs.toObject = function(includeInstance, msg) {
  var f, obj = {
    database: jspb.Message.getFieldWithDefault(msg, 1, 0),
    channelList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
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
 * @return {!proto.protos.backends.RedisPubSubWriteArgs}
 */
proto.protos.backends.RedisPubSubWriteArgs.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.backends.RedisPubSubWriteArgs;
  return proto.protos.backends.RedisPubSubWriteArgs.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.backends.RedisPubSubWriteArgs} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.backends.RedisPubSubWriteArgs}
 */
proto.protos.backends.RedisPubSubWriteArgs.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setDatabase(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addChannel(value);
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
proto.protos.backends.RedisPubSubWriteArgs.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.backends.RedisPubSubWriteArgs.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.backends.RedisPubSubWriteArgs} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.backends.RedisPubSubWriteArgs.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDatabase();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getChannelList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional uint32 database = 1;
 * @return {number}
 */
proto.protos.backends.RedisPubSubWriteArgs.prototype.getDatabase = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.protos.backends.RedisPubSubWriteArgs} returns this
 */
proto.protos.backends.RedisPubSubWriteArgs.prototype.setDatabase = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated string channel = 2;
 * @return {!Array<string>}
 */
proto.protos.backends.RedisPubSubWriteArgs.prototype.getChannelList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.protos.backends.RedisPubSubWriteArgs} returns this
 */
proto.protos.backends.RedisPubSubWriteArgs.prototype.setChannelList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.protos.backends.RedisPubSubWriteArgs} returns this
 */
proto.protos.backends.RedisPubSubWriteArgs.prototype.addChannel = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.protos.backends.RedisPubSubWriteArgs} returns this
 */
proto.protos.backends.RedisPubSubWriteArgs.prototype.clearChannelList = function() {
  return this.setChannelList([]);
};


goog.object.extend(exports, proto.protos.backends);