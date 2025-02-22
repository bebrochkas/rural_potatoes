from nlp.tag import tagger
from genproto import tagger_pb2, tagger_pb2_grpc


import grpc
from concurrent import futures


class Tagger(tagger_pb2_grpc.TaggerServicer):
    def film_tags(self, request, context):
        return tagger_pb2.TagsReply(tags=tagger(request.description))


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    tagger_pb2_grpc.add_TaggerServicer_to_server(Tagger(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
