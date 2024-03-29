# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from push import push_pb2 as push_dot_push__pb2


class PushServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendEmail = channel.unary_unary(
                '/push.PushService/SendEmail',
                request_serializer=push_dot_push__pb2.SendEmailRequest.SerializeToString,
                response_deserializer=push_dot_push__pb2.SendEmailResponse.FromString,
                )
        self.SendLarkTextByUnionIds = channel.unary_unary(
                '/push.PushService/SendLarkTextByUnionIds',
                request_serializer=push_dot_push__pb2.SendLarkTextByUnionIdsRequest.SerializeToString,
                response_deserializer=push_dot_push__pb2.SendLarkTextByUnionIdsResponse.FromString,
                )
        self.SendLarkPostByUnionIds = channel.unary_unary(
                '/push.PushService/SendLarkPostByUnionIds',
                request_serializer=push_dot_push__pb2.SendLarkPostByUnionIdsRequest.SerializeToString,
                response_deserializer=push_dot_push__pb2.SendLarkPostByUnionIdsResponse.FromString,
                )


class PushServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendEmail(self, request, context):
        """向用户发送邮件消息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendLarkTextByUnionIds(self, request, context):
        """发送飞书消息(text)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendLarkPostByUnionIds(self, request, context):
        """发送飞书消息(富文本)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PushServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendEmail': grpc.unary_unary_rpc_method_handler(
                    servicer.SendEmail,
                    request_deserializer=push_dot_push__pb2.SendEmailRequest.FromString,
                    response_serializer=push_dot_push__pb2.SendEmailResponse.SerializeToString,
            ),
            'SendLarkTextByUnionIds': grpc.unary_unary_rpc_method_handler(
                    servicer.SendLarkTextByUnionIds,
                    request_deserializer=push_dot_push__pb2.SendLarkTextByUnionIdsRequest.FromString,
                    response_serializer=push_dot_push__pb2.SendLarkTextByUnionIdsResponse.SerializeToString,
            ),
            'SendLarkPostByUnionIds': grpc.unary_unary_rpc_method_handler(
                    servicer.SendLarkPostByUnionIds,
                    request_deserializer=push_dot_push__pb2.SendLarkPostByUnionIdsRequest.FromString,
                    response_serializer=push_dot_push__pb2.SendLarkPostByUnionIdsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'push.PushService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PushService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendEmail(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/push.PushService/SendEmail',
            push_dot_push__pb2.SendEmailRequest.SerializeToString,
            push_dot_push__pb2.SendEmailResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SendLarkTextByUnionIds(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/push.PushService/SendLarkTextByUnionIds',
            push_dot_push__pb2.SendLarkTextByUnionIdsRequest.SerializeToString,
            push_dot_push__pb2.SendLarkTextByUnionIdsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SendLarkPostByUnionIds(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/push.PushService/SendLarkPostByUnionIds',
            push_dot_push__pb2.SendLarkPostByUnionIdsRequest.SerializeToString,
            push_dot_push__pb2.SendLarkPostByUnionIdsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
