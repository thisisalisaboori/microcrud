import sys
sys.path.append('/media/ali/Project/ali/api/contracts')
import grpc
from concurrent import futures
import microcrud_pb2
import microcrud_pb2_grpc
from google.protobuf import struct_pb2



class CrudEntityService(microcrud_pb2_grpc.CrudEntityServicer):
    def Create(self, request, context):
        # request.data از نوع Struct هست
        print("Received data:", request.data)

        # می‌تونی دیکشنریش رو بگیری
        data_dict = dict(request.data)
        print("Converted to dict:", data_dict)

        # پاسخ بساز
        resp_data = struct_pb2.Struct()
        resp_data.update({
            "status": "created",
            "id": "12345"
        })

        return microcrud_pb2.ResultService(
            message="Entity created successfully",
            data=resp_data
        )

    def Update(self, request, context):
        return microcrud_pb2.ResultService(message="Update not implemented")

    def Delete(self, request, context):
        return microcrud_pb2.ResultService(message="Delete not implemented")

    def GetById(self, request, context):
        return microcrud_pb2.ResultService(message="GetById not implemented")

    def GetData(self, request, context):
        return microcrud_pb2.ResultService(message="GetData not implemented")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    microcrud_pb2_grpc.add_CrudEntityServicer_to_server(CrudEntityService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started on port 50051...")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
