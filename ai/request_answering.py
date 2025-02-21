from ai.tags_parcing import tagger
import grpc
from concurrent import futures
import service_pb2
import service_pb2_grpc

class Nlp(service_pb2_grpc.NlpServicer):
    def film_tags(self, request, context):
        return service_pb2.TagsReply(tags=tagger(request.description))

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_NlpServicer_to_server(Nlp(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()