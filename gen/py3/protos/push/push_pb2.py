# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/push/push.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from protos.common import common_pb2 as protos_dot_common_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16protos/push/push.proto\x12\x11go_interface.push\x1a\x19google/protobuf/any.proto\x1a\x1aprotos/common/common.proto\"X\n\x05\x45mail\x12\r\n\x05topic\x18\x06 \x01(\t\x12\x0c\n\x04name\x18\x07 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x08 \x01(\t\x12\x13\n\x0b\x63ontentType\x18\t \x01(\t\x12\x0f\n\x07\x63ontent\x18\n \x01(\t\"o\n\x10SendEmailRequest\x12\x32\n\x06header\x18\x01 \x01(\x0b\x32\".go_interface.common.requestHeader\x12\'\n\x05\x65mail\x18\x02 \x01(\x0b\x32\x18.go_interface.push.Email\"F\n\x11SendEmailResponse\x12\x31\n\terrorInfo\x18\x01 \x01(\x0b\x32\x1e.go_interface.common.errorInfo\"v\n\x1dSendLarkTextByUnionIdsRequest\x12\x32\n\x06header\x18\x01 \x01(\x0b\x32\".go_interface.common.requestHeader\x12\x10\n\x08unionIds\x18\x02 \x03(\t\x12\x0f\n\x07\x63ontent\x18\x03 \x01(\t\"S\n\x1eSendLarkTextByUnionIdsResponse\x12\x31\n\terrorInfo\x18\x01 \x01(\x0b\x32\x1e.go_interface.common.errorInfo\"\x8c\x01\n\x1dSendLarkPostByUnionIdsRequest\x12\x32\n\x06header\x18\x01 \x01(\x0b\x32\".go_interface.common.requestHeader\x12\x10\n\x08unionIds\x18\x02 \x03(\t\x12%\n\x07\x63ontent\x18\x03 \x01(\x0b\x32\x14.google.protobuf.Any\"S\n\x1eSendLarkPostByUnionIdsResponse\x12\x31\n\terrorInfo\x18\x01 \x01(\x0b\x32\x1e.go_interface.common.errorInfo2\xdc\x02\n\x04Push\x12V\n\tSendEmail\x12#.go_interface.push.SendEmailRequest\x1a$.go_interface.push.SendEmailResponse\x12}\n\x16SendLarkTextByUnionIds\x12\x30.go_interface.push.SendLarkTextByUnionIdsRequest\x1a\x31.go_interface.push.SendLarkTextByUnionIdsResponse\x12}\n\x16SendLarkPostByUnionIds\x12\x30.go_interface.push.SendLarkPostByUnionIdsRequest\x1a\x31.go_interface.push.SendLarkPostByUnionIdsResponseB0Z.github.com/solost23/protopb/gen/go/protos/pushb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.push.push_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z.github.com/solost23/protopb/gen/go/protos/push'
  _EMAIL._serialized_start=100
  _EMAIL._serialized_end=188
  _SENDEMAILREQUEST._serialized_start=190
  _SENDEMAILREQUEST._serialized_end=301
  _SENDEMAILRESPONSE._serialized_start=303
  _SENDEMAILRESPONSE._serialized_end=373
  _SENDLARKTEXTBYUNIONIDSREQUEST._serialized_start=375
  _SENDLARKTEXTBYUNIONIDSREQUEST._serialized_end=493
  _SENDLARKTEXTBYUNIONIDSRESPONSE._serialized_start=495
  _SENDLARKTEXTBYUNIONIDSRESPONSE._serialized_end=578
  _SENDLARKPOSTBYUNIONIDSREQUEST._serialized_start=581
  _SENDLARKPOSTBYUNIONIDSREQUEST._serialized_end=721
  _SENDLARKPOSTBYUNIONIDSRESPONSE._serialized_start=723
  _SENDLARKPOSTBYUNIONIDSRESPONSE._serialized_end=806
  _PUSH._serialized_start=809
  _PUSH._serialized_end=1157
# @@protoc_insertion_point(module_scope)