# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: push/push.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0fpush/push.proto\x1a\x19google/protobuf/any.proto\"X\n\x05\x45mail\x12\r\n\x05topic\x18\x06 \x01(\t\x12\x0c\n\x04name\x18\x07 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x08 \x01(\t\x12\x13\n\x0b\x63ontentType\x18\t \x01(\t\x12\x0f\n\x07\x63ontent\x18\n \x01(\t\")\n\x10SendEmailRequest\x12\x15\n\x05\x65mail\x18\x02 \x01(\x0b\x32\x06.Email\"\x13\n\x11SendEmailResponse\"B\n\x1dSendLarkTextByUnionIdsRequest\x12\x10\n\x08unionIds\x18\x02 \x03(\t\x12\x0f\n\x07\x63ontent\x18\x03 \x01(\t\" \n\x1eSendLarkTextByUnionIdsResponse\"X\n\x1dSendLarkPostByUnionIdsRequest\x12\x10\n\x08unionIds\x18\x02 \x03(\t\x12%\n\x07\x63ontent\x18\x03 \x01(\x0b\x32\x14.google.protobuf.Any\" \n\x1eSendLarkPostByUnionIdsResponse2\xf0\x01\n\x04Push\x12\x32\n\tSendEmail\x12\x11.SendEmailRequest\x1a\x12.SendEmailResponse\x12Y\n\x16SendLarkTextByUnionIds\x12\x1e.SendLarkTextByUnionIdsRequest\x1a\x1f.SendLarkTextByUnionIdsResponse\x12Y\n\x16SendLarkPostByUnionIds\x12\x1e.SendLarkPostByUnionIdsRequest\x1a\x1f.SendLarkPostByUnionIdsResponseb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'push.push_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _EMAIL._serialized_start=46
  _EMAIL._serialized_end=134
  _SENDEMAILREQUEST._serialized_start=136
  _SENDEMAILREQUEST._serialized_end=177
  _SENDEMAILRESPONSE._serialized_start=179
  _SENDEMAILRESPONSE._serialized_end=198
  _SENDLARKTEXTBYUNIONIDSREQUEST._serialized_start=200
  _SENDLARKTEXTBYUNIONIDSREQUEST._serialized_end=266
  _SENDLARKTEXTBYUNIONIDSRESPONSE._serialized_start=268
  _SENDLARKTEXTBYUNIONIDSRESPONSE._serialized_end=300
  _SENDLARKPOSTBYUNIONIDSREQUEST._serialized_start=302
  _SENDLARKPOSTBYUNIONIDSREQUEST._serialized_end=390
  _SENDLARKPOSTBYUNIONIDSRESPONSE._serialized_start=392
  _SENDLARKPOSTBYUNIONIDSRESPONSE._serialized_end=424
  _PUSH._serialized_start=427
  _PUSH._serialized_end=667
# @@protoc_insertion_point(module_scope)
