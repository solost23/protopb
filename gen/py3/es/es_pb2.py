# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: es/es.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common import common_pb2 as common_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0b\x65s/es.proto\x12\x02\x65s\x1a\x13\x63ommon/common.proto\")\n\tTermQuery\x12\r\n\x05\x66ield\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"*\n\nMatchQuery\x12\r\n\x05\x66ield\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"/\n\x0fMultiMatchQuery\x12\r\n\x05\x66ield\x18\x01 \x03(\t\x12\r\n\x05value\x18\x02 \x01(\t\"M\n\nRangeQuery\x12\r\n\x05\x66ield\x18\x01 \x01(\t\x12\n\n\x02gt\x18\x02 \x01(\t\x12\n\n\x02lt\x18\x03 \x01(\t\x12\x0b\n\x03gte\x18\x04 \x01(\t\x12\x0b\n\x03lte\x18\x05 \x01(\t\"\xa7\x01\n\x05Query\x12\"\n\x0btermQueries\x18\x01 \x03(\x0b\x32\r.es.TermQuery\x12$\n\x0cmatchQueries\x18\x02 \x03(\x0b\x32\x0e.es.MatchQuery\x12.\n\x11multiMatchQueries\x18\x03 \x03(\x0b\x32\x13.es.MultiMatchQuery\x12$\n\x0crangeQueries\x18\x04 \x03(\x0b\x32\x0e.es.RangeQuery\"(\n\x04Sort\x12\r\n\x05\x66ield\x18\x01 \x01(\t\x12\x11\n\tascending\x18\x02 \x01(\x08\"\xea\x01\n\rSearchRequest\x12%\n\x06header\x18\x01 \x01(\x0b\x32\x15.common.RequestHeader\x12\x1c\n\tmustQuery\x18\x02 \x01(\x0b\x32\t.es.Query\x12\x1f\n\x0cmustNotQuery\x18\x03 \x01(\x0b\x32\t.es.Query\x12\x1e\n\x0bshouldQuery\x18\x04 \x01(\x0b\x32\t.es.Query\x12\x0f\n\x07indices\x18\x05 \x03(\t\x12\x16\n\x04sort\x18\x06 \x03(\x0b\x32\x08.es.Sort\x12\x0c\n\x04page\x18\x07 \x01(\x11\x12\x0c\n\x04size\x18\x08 \x01(\x11\x12\x0e\n\x06pretty\x18\t \x01(\x08\"G\n\x08PageList\x12\x0c\n\x04size\x18\x01 \x01(\x05\x12\r\n\x05pages\x18\x02 \x01(\x03\x12\r\n\x05total\x18\x03 \x01(\x03\x12\x0f\n\x07\x63urrent\x18\x04 \x01(\x05\"g\n\x0eSearchResponse\x12$\n\terrorInfo\x18\x01 \x01(\x0b\x32\x11.common.ErrorInfo\x12\x1e\n\x08pageList\x18\x02 \x01(\x0b\x32\x0c.es.PageList\x12\x0f\n\x07records\x18\x03 \x03(\t\"k\n\rCreateRequest\x12%\n\x06header\x18\x01 \x01(\x0b\x32\x15.common.RequestHeader\x12\x12\n\ndocumentId\x18\x02 \x01(\t\x12\r\n\x05index\x18\x03 \x01(\t\x12\x10\n\x08\x64ocument\x18\x04 \x01(\t\"6\n\x0e\x43reateResponse\x12$\n\terrorInfo\x18\x01 \x01(\x0b\x32\x11.common.ErrorInfo\"Y\n\rDeleteRequest\x12%\n\x06header\x18\x01 \x01(\x0b\x32\x15.common.RequestHeader\x12\r\n\x05index\x18\x02 \x01(\t\x12\x12\n\ndocumentId\x18\x03 \x01(\t\"6\n\x0e\x44\x65leteResponse\x12$\n\terrorInfo\x18\x01 \x01(\x0b\x32\x11.common.ErrorInfo2\xa2\x01\n\rSearchService\x12/\n\x06Search\x12\x11.es.SearchRequest\x1a\x12.es.SearchResponse\x12/\n\x06\x43reate\x12\x11.es.CreateRequest\x1a\x12.es.CreateResponse\x12/\n\x06\x44\x65lete\x12\x11.es.DeleteRequest\x1a\x12.es.DeleteResponseB\'Z%github.com/solost23/protopb/gen/go/esb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'es.es_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z%github.com/solost23/protopb/gen/go/es'
  _TERMQUERY._serialized_start=40
  _TERMQUERY._serialized_end=81
  _MATCHQUERY._serialized_start=83
  _MATCHQUERY._serialized_end=125
  _MULTIMATCHQUERY._serialized_start=127
  _MULTIMATCHQUERY._serialized_end=174
  _RANGEQUERY._serialized_start=176
  _RANGEQUERY._serialized_end=253
  _QUERY._serialized_start=256
  _QUERY._serialized_end=423
  _SORT._serialized_start=425
  _SORT._serialized_end=465
  _SEARCHREQUEST._serialized_start=468
  _SEARCHREQUEST._serialized_end=702
  _PAGELIST._serialized_start=704
  _PAGELIST._serialized_end=775
  _SEARCHRESPONSE._serialized_start=777
  _SEARCHRESPONSE._serialized_end=880
  _CREATEREQUEST._serialized_start=882
  _CREATEREQUEST._serialized_end=989
  _CREATERESPONSE._serialized_start=991
  _CREATERESPONSE._serialized_end=1045
  _DELETEREQUEST._serialized_start=1047
  _DELETEREQUEST._serialized_end=1136
  _DELETERESPONSE._serialized_start=1138
  _DELETERESPONSE._serialized_end=1192
  _SEARCHSERVICE._serialized_start=1195
  _SEARCHSERVICE._serialized_end=1357
# @@protoc_insertion_point(module_scope)