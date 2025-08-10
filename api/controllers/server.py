from typing import Union
import sys,os

sys.path.append('/media/ali/Project/ali/api/contracts')

import grpc
from contracts import microcrud_pb2
from contracts import  microcrud_pb2_grpc
from google.protobuf import struct_pb2


from fastapi import FastAPI

app = FastAPI()

channel = grpc.insecure_channel('localhost:50051')
stub = microcrud_pb2_grpc.CrudEntityStub(channel)


@app.get("/")
def read_root():
    data_struct = struct_pb2.Struct()
    data_struct.update({
        "name": "Ali",
        "email": "ali@example.com",
        "age": 30
    })
    entity = microcrud_pb2.EntityProtocol(data=data_struct)
    response = stub.Create(entity)


    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}