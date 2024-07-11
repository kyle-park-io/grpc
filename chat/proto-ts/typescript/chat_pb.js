// source: chat.proto
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
var global =
  (typeof globalThis !== 'undefined' && globalThis) ||
  (typeof window !== 'undefined' && window) ||
  (typeof global !== 'undefined' && global) ||
  (typeof self !== 'undefined' && self) ||
  function () {
    return this;
  }.call(null) ||
  Function('return this')();

var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
goog.object.extend(proto, google_protobuf_wrappers_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.kyle_chat.ChatMsg', null, global);
goog.exportSymbol('proto.kyle_chat.GetChatMsg', null, global);
goog.exportSymbol('proto.kyle_chat.GetTestMsg', null, global);
goog.exportSymbol('proto.kyle_chat.TestMsg', null, global);
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
proto.kyle_chat.TestMsg = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.kyle_chat.TestMsg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.kyle_chat.TestMsg.displayName = 'proto.kyle_chat.TestMsg';
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
proto.kyle_chat.ChatMsg = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.kyle_chat.ChatMsg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.kyle_chat.ChatMsg.displayName = 'proto.kyle_chat.ChatMsg';
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
proto.kyle_chat.GetTestMsg = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.kyle_chat.GetTestMsg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.kyle_chat.GetTestMsg.displayName = 'proto.kyle_chat.GetTestMsg';
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
proto.kyle_chat.GetChatMsg = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.kyle_chat.GetChatMsg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.kyle_chat.GetChatMsg.displayName = 'proto.kyle_chat.GetChatMsg';
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
  proto.kyle_chat.TestMsg.prototype.toObject = function (opt_includeInstance) {
    return proto.kyle_chat.TestMsg.toObject(opt_includeInstance, this);
  };

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Deprecated. Whether to include
   *     the JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.kyle_chat.TestMsg} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.kyle_chat.TestMsg.toObject = function (includeInstance, msg) {
    var f,
      obj = {
        id: jspb.Message.getFieldWithDefault(msg, 1, 0),
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
 * @return {!proto.kyle_chat.TestMsg}
 */
proto.kyle_chat.TestMsg.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.kyle_chat.TestMsg();
  return proto.kyle_chat.TestMsg.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.kyle_chat.TestMsg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.kyle_chat.TestMsg}
 */
proto.kyle_chat.TestMsg.deserializeBinaryFromReader = function (msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = /** @type {number} */ (reader.readUint64());
        msg.setId(value);
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
proto.kyle_chat.TestMsg.prototype.serializeBinary = function () {
  var writer = new jspb.BinaryWriter();
  proto.kyle_chat.TestMsg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.kyle_chat.TestMsg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.kyle_chat.TestMsg.serializeBinaryToWriter = function (message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(1, f);
  }
};

/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.kyle_chat.TestMsg.prototype.getId = function () {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};

/**
 * @param {number} value
 * @return {!proto.kyle_chat.TestMsg} returns this
 */
proto.kyle_chat.TestMsg.prototype.setId = function (value) {
  return jspb.Message.setProto3IntField(this, 1, value);
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
  proto.kyle_chat.ChatMsg.prototype.toObject = function (opt_includeInstance) {
    return proto.kyle_chat.ChatMsg.toObject(opt_includeInstance, this);
  };

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Deprecated. Whether to include
   *     the JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.kyle_chat.ChatMsg} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.kyle_chat.ChatMsg.toObject = function (includeInstance, msg) {
    var f,
      obj = {
        id: jspb.Message.getFieldWithDefault(msg, 1, 0),
        userId: jspb.Message.getFieldWithDefault(msg, 2, ''),
        text: jspb.Message.getFieldWithDefault(msg, 3, ''),
        eventTime:
          (f = msg.getEventTime()) &&
          google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
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
 * @return {!proto.kyle_chat.ChatMsg}
 */
proto.kyle_chat.ChatMsg.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.kyle_chat.ChatMsg();
  return proto.kyle_chat.ChatMsg.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.kyle_chat.ChatMsg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.kyle_chat.ChatMsg}
 */
proto.kyle_chat.ChatMsg.deserializeBinaryFromReader = function (msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = /** @type {number} */ (reader.readUint64());
        msg.setId(value);
        break;
      case 2:
        var value = /** @type {string} */ (reader.readString());
        msg.setUserId(value);
        break;
      case 3:
        var value = /** @type {string} */ (reader.readString());
        msg.setText(value);
        break;
      case 4:
        var value = new google_protobuf_timestamp_pb.Timestamp();
        reader.readMessage(
          value,
          google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader,
        );
        msg.setEventTime(value);
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
proto.kyle_chat.ChatMsg.prototype.serializeBinary = function () {
  var writer = new jspb.BinaryWriter();
  proto.kyle_chat.ChatMsg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.kyle_chat.ChatMsg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.kyle_chat.ChatMsg.serializeBinaryToWriter = function (message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(1, f);
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(2, f);
  }
  f = message.getText();
  if (f.length > 0) {
    writer.writeString(3, f);
  }
  f = message.getEventTime();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter,
    );
  }
};

/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.kyle_chat.ChatMsg.prototype.getId = function () {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};

/**
 * @param {number} value
 * @return {!proto.kyle_chat.ChatMsg} returns this
 */
proto.kyle_chat.ChatMsg.prototype.setId = function (value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};

/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.kyle_chat.ChatMsg.prototype.getUserId = function () {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ''));
};

/**
 * @param {string} value
 * @return {!proto.kyle_chat.ChatMsg} returns this
 */
proto.kyle_chat.ChatMsg.prototype.setUserId = function (value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};

/**
 * optional string text = 3;
 * @return {string}
 */
proto.kyle_chat.ChatMsg.prototype.getText = function () {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ''));
};

/**
 * @param {string} value
 * @return {!proto.kyle_chat.ChatMsg} returns this
 */
proto.kyle_chat.ChatMsg.prototype.setText = function (value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};

/**
 * optional google.protobuf.Timestamp event_time = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.kyle_chat.ChatMsg.prototype.getEventTime = function () {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(
      this,
      google_protobuf_timestamp_pb.Timestamp,
      4,
    )
  );
};

/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.kyle_chat.ChatMsg} returns this
 */
proto.kyle_chat.ChatMsg.prototype.setEventTime = function (value) {
  return jspb.Message.setWrapperField(this, 4, value);
};

/**
 * Clears the message field making it undefined.
 * @return {!proto.kyle_chat.ChatMsg} returns this
 */
proto.kyle_chat.ChatMsg.prototype.clearEventTime = function () {
  return this.setEventTime(undefined);
};

/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.kyle_chat.ChatMsg.prototype.hasEventTime = function () {
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
  proto.kyle_chat.GetTestMsg.prototype.toObject = function (
    opt_includeInstance,
  ) {
    return proto.kyle_chat.GetTestMsg.toObject(opt_includeInstance, this);
  };

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Deprecated. Whether to include
   *     the JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.kyle_chat.GetTestMsg} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.kyle_chat.GetTestMsg.toObject = function (includeInstance, msg) {
    var f,
      obj = {
        id: jspb.Message.getFieldWithDefault(msg, 1, 0),
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
 * @return {!proto.kyle_chat.GetTestMsg}
 */
proto.kyle_chat.GetTestMsg.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.kyle_chat.GetTestMsg();
  return proto.kyle_chat.GetTestMsg.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.kyle_chat.GetTestMsg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.kyle_chat.GetTestMsg}
 */
proto.kyle_chat.GetTestMsg.deserializeBinaryFromReader = function (
  msg,
  reader,
) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = /** @type {number} */ (reader.readUint64());
        msg.setId(value);
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
proto.kyle_chat.GetTestMsg.prototype.serializeBinary = function () {
  var writer = new jspb.BinaryWriter();
  proto.kyle_chat.GetTestMsg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.kyle_chat.GetTestMsg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.kyle_chat.GetTestMsg.serializeBinaryToWriter = function (
  message,
  writer,
) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(1, f);
  }
};

/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.kyle_chat.GetTestMsg.prototype.getId = function () {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};

/**
 * @param {number} value
 * @return {!proto.kyle_chat.GetTestMsg} returns this
 */
proto.kyle_chat.GetTestMsg.prototype.setId = function (value) {
  return jspb.Message.setProto3IntField(this, 1, value);
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
  proto.kyle_chat.GetChatMsg.prototype.toObject = function (
    opt_includeInstance,
  ) {
    return proto.kyle_chat.GetChatMsg.toObject(opt_includeInstance, this);
  };

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Deprecated. Whether to include
   *     the JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.kyle_chat.GetChatMsg} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.kyle_chat.GetChatMsg.toObject = function (includeInstance, msg) {
    var f,
      obj = {
        id: jspb.Message.getFieldWithDefault(msg, 1, 0),
        userId: jspb.Message.getFieldWithDefault(msg, 2, ''),
        eventTime:
          (f = msg.getEventTime()) &&
          google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
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
 * @return {!proto.kyle_chat.GetChatMsg}
 */
proto.kyle_chat.GetChatMsg.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.kyle_chat.GetChatMsg();
  return proto.kyle_chat.GetChatMsg.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.kyle_chat.GetChatMsg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.kyle_chat.GetChatMsg}
 */
proto.kyle_chat.GetChatMsg.deserializeBinaryFromReader = function (
  msg,
  reader,
) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = /** @type {number} */ (reader.readUint64());
        msg.setId(value);
        break;
      case 2:
        var value = /** @type {string} */ (reader.readString());
        msg.setUserId(value);
        break;
      case 4:
        var value = new google_protobuf_timestamp_pb.Timestamp();
        reader.readMessage(
          value,
          google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader,
        );
        msg.setEventTime(value);
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
proto.kyle_chat.GetChatMsg.prototype.serializeBinary = function () {
  var writer = new jspb.BinaryWriter();
  proto.kyle_chat.GetChatMsg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.kyle_chat.GetChatMsg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.kyle_chat.GetChatMsg.serializeBinaryToWriter = function (
  message,
  writer,
) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(1, f);
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(2, f);
  }
  f = message.getEventTime();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter,
    );
  }
};

/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.kyle_chat.GetChatMsg.prototype.getId = function () {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};

/**
 * @param {number} value
 * @return {!proto.kyle_chat.GetChatMsg} returns this
 */
proto.kyle_chat.GetChatMsg.prototype.setId = function (value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};

/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.kyle_chat.GetChatMsg.prototype.getUserId = function () {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ''));
};

/**
 * @param {string} value
 * @return {!proto.kyle_chat.GetChatMsg} returns this
 */
proto.kyle_chat.GetChatMsg.prototype.setUserId = function (value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};

/**
 * optional google.protobuf.Timestamp event_time = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.kyle_chat.GetChatMsg.prototype.getEventTime = function () {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(
      this,
      google_protobuf_timestamp_pb.Timestamp,
      4,
    )
  );
};

/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.kyle_chat.GetChatMsg} returns this
 */
proto.kyle_chat.GetChatMsg.prototype.setEventTime = function (value) {
  return jspb.Message.setWrapperField(this, 4, value);
};

/**
 * Clears the message field making it undefined.
 * @return {!proto.kyle_chat.GetChatMsg} returns this
 */
proto.kyle_chat.GetChatMsg.prototype.clearEventTime = function () {
  return this.setEventTime(undefined);
};

/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.kyle_chat.GetChatMsg.prototype.hasEventTime = function () {
  return jspb.Message.getField(this, 4) != null;
};

goog.object.extend(exports, proto.kyle_chat);
