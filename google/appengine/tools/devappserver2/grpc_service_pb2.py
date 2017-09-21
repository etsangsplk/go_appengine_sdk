#!/usr/bin/env python
#
# Copyright 2007 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: apphosting/tools/devappserver2/grpc_service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
import google
from google.net.proto2.python.public import descriptor as _descriptor
from google.net.proto2.python.public import message as _message
from google.net.proto2.python.public import reflection as _reflection
from google.net.proto2.python.public import symbol_database as _symbol_database
from google.net.proto2.proto import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='apphosting/tools/devappserver2/grpc_service.proto',
  package='apphosting.tools.devappserver2',
  syntax='proto3',
  serialized_pb=_b('\n1apphosting/tools/devappserver2/grpc_service.proto\x12\x1e\x61pphosting.tools.devappserver2\"T\n\x07Request\x12\x14\n\x0cservice_name\x18\x02 \x01(\t\x12\x0e\n\x06method\x18\x03 \x01(\t\x12\x0f\n\x07request\x18\x04 \x01(\x0c\x12\x12\n\nrequest_id\x18\x05 \x01(\t\"\xd1\x01\n\x08Response\x12\x10\n\x08response\x18\x01 \x01(\x0c\x12\x11\n\texception\x18\x02 \x01(\x0c\x12K\n\x11\x61pplication_error\x18\x03 \x01(\x0b\x32\x30.apphosting.tools.devappserver2.ApplicationError\x12\x16\n\x0ejava_exception\x18\x04 \x01(\x0c\x12;\n\trpc_error\x18\x05 \x01(\x0b\x32(.apphosting.tools.devappserver2.RpcError\"0\n\x10\x41pplicationError\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0e\n\x06\x64\x65tail\x18\x02 \x01(\t\"\xb7\x02\n\x08RpcError\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0e\n\x06\x64\x65tail\x18\x02 \x01(\t\"\x8c\x02\n\tErrorCode\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x12\n\x0e\x43\x41LL_NOT_FOUND\x10\x01\x12\x0f\n\x0bPARSE_ERROR\x10\x02\x12\x16\n\x12SECURITY_VIOLATION\x10\x03\x12\x0e\n\nOVER_QUOTA\x10\x04\x12\x15\n\x11REQUEST_TOO_LARGE\x10\x05\x12\x17\n\x13\x43\x41PABILITY_DISABLED\x10\x06\x12\x14\n\x10\x46\x45\x41TURE_DISABLED\x10\x07\x12\x0f\n\x0b\x42\x41\x44_REQUEST\x10\x08\x12\x16\n\x12RESPONSE_TOO_LARGE\x10\t\x12\r\n\tCANCELLED\x10\n\x12\x10\n\x0cREPLAY_ERROR\x10\x0b\x12\x15\n\x11\x44\x45\x41\x44LINE_EXCEEDED\x10\x0c\x32p\n\x0b\x43\x61llHandler\x12\x61\n\nHandleCall\x12\'.apphosting.tools.devappserver2.Request\x1a(.apphosting.tools.devappserver2.Response\"\x00\x62\x06proto3')
)



_RPCERROR_ERRORCODE = _descriptor.EnumDescriptor(
  name='ErrorCode',
  full_name='apphosting.tools.devappserver2.RpcError.ErrorCode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CALL_NOT_FOUND', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PARSE_ERROR', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SECURITY_VIOLATION', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OVER_QUOTA', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REQUEST_TOO_LARGE', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CAPABILITY_DISABLED', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FEATURE_DISABLED', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BAD_REQUEST', index=8, number=8,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RESPONSE_TOO_LARGE', index=9, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CANCELLED', index=10, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REPLAY_ERROR', index=11, number=11,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEADLINE_EXCEEDED', index=12, number=12,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=477,
  serialized_end=745,
)
_sym_db.RegisterEnumDescriptor(_RPCERROR_ERRORCODE)


_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='apphosting.tools.devappserver2.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='service_name', full_name='apphosting.tools.devappserver2.Request.service_name', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='method', full_name='apphosting.tools.devappserver2.Request.method', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='request', full_name='apphosting.tools.devappserver2.Request.request', index=2,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='request_id', full_name='apphosting.tools.devappserver2.Request.request_id', index=3,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=85,
  serialized_end=169,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='apphosting.tools.devappserver2.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='response', full_name='apphosting.tools.devappserver2.Response.response', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='exception', full_name='apphosting.tools.devappserver2.Response.exception', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='application_error', full_name='apphosting.tools.devappserver2.Response.application_error', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='java_exception', full_name='apphosting.tools.devappserver2.Response.java_exception', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rpc_error', full_name='apphosting.tools.devappserver2.Response.rpc_error', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=172,
  serialized_end=381,
)


_APPLICATIONERROR = _descriptor.Descriptor(
  name='ApplicationError',
  full_name='apphosting.tools.devappserver2.ApplicationError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='apphosting.tools.devappserver2.ApplicationError.code', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='detail', full_name='apphosting.tools.devappserver2.ApplicationError.detail', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=383,
  serialized_end=431,
)


_RPCERROR = _descriptor.Descriptor(
  name='RpcError',
  full_name='apphosting.tools.devappserver2.RpcError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='apphosting.tools.devappserver2.RpcError.code', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='detail', full_name='apphosting.tools.devappserver2.RpcError.detail', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _RPCERROR_ERRORCODE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=434,
  serialized_end=745,
)

_RESPONSE.fields_by_name['application_error'].message_type = _APPLICATIONERROR
_RESPONSE.fields_by_name['rpc_error'].message_type = _RPCERROR
_RPCERROR_ERRORCODE.containing_type = _RPCERROR
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
DESCRIPTOR.message_types_by_name['ApplicationError'] = _APPLICATIONERROR
DESCRIPTOR.message_types_by_name['RpcError'] = _RPCERROR
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), dict(
  DESCRIPTOR = _REQUEST,
  __module__ = 'google.appengine.tools.devappserver2.grpc_service_pb2'
  # @@protoc_insertion_point(class_scope:apphosting.tools.devappserver2.Request)
  ))
_sym_db.RegisterMessage(Request)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), dict(
  DESCRIPTOR = _RESPONSE,
  __module__ = 'google.appengine.tools.devappserver2.grpc_service_pb2'
  # @@protoc_insertion_point(class_scope:apphosting.tools.devappserver2.Response)
  ))
_sym_db.RegisterMessage(Response)

ApplicationError = _reflection.GeneratedProtocolMessageType('ApplicationError', (_message.Message,), dict(
  DESCRIPTOR = _APPLICATIONERROR,
  __module__ = 'google.appengine.tools.devappserver2.grpc_service_pb2'
  # @@protoc_insertion_point(class_scope:apphosting.tools.devappserver2.ApplicationError)
  ))
_sym_db.RegisterMessage(ApplicationError)

RpcError = _reflection.GeneratedProtocolMessageType('RpcError', (_message.Message,), dict(
  DESCRIPTOR = _RPCERROR,
  __module__ = 'google.appengine.tools.devappserver2.grpc_service_pb2'
  # @@protoc_insertion_point(class_scope:apphosting.tools.devappserver2.RpcError)
  ))
_sym_db.RegisterMessage(RpcError)



_CALLHANDLER = _descriptor.ServiceDescriptor(
  name='CallHandler',
  full_name='apphosting.tools.devappserver2.CallHandler',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=747,
  serialized_end=859,
  methods=[
  _descriptor.MethodDescriptor(
    name='HandleCall',
    full_name='apphosting.tools.devappserver2.CallHandler.HandleCall',
    index=0,
    containing_service=None,
    input_type=_REQUEST,
    output_type=_RESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CALLHANDLER)

DESCRIPTOR.services_by_name['CallHandler'] = _CALLHANDLER

try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class CallHandlerStub(object):
    # missing associated documentation comment in .proto file
    pass

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.HandleCall = channel.unary_unary(
          '/apphosting.tools.devappserver2.CallHandler/HandleCall',
          request_serializer=Request.SerializeToString,
          response_deserializer=Response.FromString,
          )


  class CallHandlerServicer(object):
    # missing associated documentation comment in .proto file
    pass

    def HandleCall(self, request, context):
      """Handles remote api call over gRPC.
      """
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_CallHandlerServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'HandleCall': grpc.unary_unary_rpc_method_handler(
            servicer.HandleCall,
            request_deserializer=Request.FromString,
            response_serializer=Response.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'apphosting.tools.devappserver2.CallHandler', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaCallHandlerServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def HandleCall(self, request, context):
      """Handles remote api call over gRPC.
      """
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaCallHandlerStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def HandleCall(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      """Handles remote api call over gRPC.
      """
      raise NotImplementedError()
    HandleCall.future = None


  def beta_create_CallHandler_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('apphosting.tools.devappserver2.CallHandler', 'HandleCall'): Request.FromString,
    }
    response_serializers = {
      ('apphosting.tools.devappserver2.CallHandler', 'HandleCall'): Response.SerializeToString,
    }
    method_implementations = {
      ('apphosting.tools.devappserver2.CallHandler', 'HandleCall'): face_utilities.unary_unary_inline(servicer.HandleCall),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_CallHandler_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('apphosting.tools.devappserver2.CallHandler', 'HandleCall'): Request.SerializeToString,
    }
    response_deserializers = {
      ('apphosting.tools.devappserver2.CallHandler', 'HandleCall'): Response.FromString,
    }
    cardinalities = {
      'HandleCall': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'apphosting.tools.devappserver2.CallHandler', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
