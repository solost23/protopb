# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from es import es_pb2 as es_dot_es__pb2


class SearchServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Search = channel.unary_unary(
                '/es.SearchService/Search',
                request_serializer=es_dot_es__pb2.SearchRequest.SerializeToString,
                response_deserializer=es_dot_es__pb2.SearchResponse.FromString,
                )
        self.Create = channel.unary_unary(
                '/es.SearchService/Create',
                request_serializer=es_dot_es__pb2.CreateRequest.SerializeToString,
                response_deserializer=es_dot_es__pb2.CreateResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/es.SearchService/Delete',
                request_serializer=es_dot_es__pb2.DeleteRequest.SerializeToString,
                response_deserializer=es_dot_es__pb2.DeleteResponse.FromString,
                )


class SearchServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Search(self, request, context):
        """搜索
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Create(self, request, context):
        """文档创建(自动创建index)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """删除
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SearchServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Search': grpc.unary_unary_rpc_method_handler(
                    servicer.Search,
                    request_deserializer=es_dot_es__pb2.SearchRequest.FromString,
                    response_serializer=es_dot_es__pb2.SearchResponse.SerializeToString,
            ),
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=es_dot_es__pb2.CreateRequest.FromString,
                    response_serializer=es_dot_es__pb2.CreateResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=es_dot_es__pb2.DeleteRequest.FromString,
                    response_serializer=es_dot_es__pb2.DeleteResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'es.SearchService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SearchService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Search(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/es.SearchService/Search',
            es_dot_es__pb2.SearchRequest.SerializeToString,
            es_dot_es__pb2.SearchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/es.SearchService/Create',
            es_dot_es__pb2.CreateRequest.SerializeToString,
            es_dot_es__pb2.CreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/es.SearchService/Delete',
            es_dot_es__pb2.DeleteRequest.SerializeToString,
            es_dot_es__pb2.DeleteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)