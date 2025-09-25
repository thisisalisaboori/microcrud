from typing import Union
import sys,os
import uvicorn

sys.path.append('/media/ali/Project/ali/microcud/api/endpoint/contracts')

import grpc
from contracts import crud_pb2
from contracts import  crud_pb2_grpc
from google.protobuf import struct_pb2
from pydantic import BaseModel

from fastapi import FastAPI


class CrudBaseMode(BaseModel):
    data: dict


app = FastAPI()

channel = grpc.insecure_channel('localhost:50051')
stub = crud_pb2_grpc.CrudServiceStub(channel)



@app.get("/{entity}/{id}")
def getbyid(id:str , entity:str):
    
    entity = crud_pb2.GetItemRequest(id=id ,entity=entity,Bucket= "ali")
    response = stub.GetItemById(entity)
    print(response.data)
    return {"Ok":response.ok , "data": response.data}  

@app.post("/{entity}")
def create(entity:str, js:CrudBaseMode):
    data_struct = struct_pb2.Struct()
    data_struct.update(js.data)
    print(data_struct)
    data= crud_pb2.CreateItemRequest(entity= entity,data= data_struct,Bucket= "ali")
    response= stub.CreateItem(data)
    
    return {"Ok":response.ok} 


@app.put("/{entity}/{id}")
def update(entity:str, id:str,js:CrudBaseMode):
    print(js)
    data_struct = struct_pb2.Struct()
    data_struct.update(js.data)
    data= crud_pb2.UpdateItemRequest(id= id,entity= entity,data= data_struct,Bucket= "ali")
    response= stub.UpdateItem(data)
    return {"Ok":response.ok} 


@app.delete("/{entity}/{id}")
def delete(entity:str, id:str):
    data= crud_pb2.DeleteItemRequest(id= id,entity= entity)
    response= stub.DeleteItem(data)
    return {"Ok":response.ok} 

@app.get("/{entity}/")
async def getItems(entity:str,pageindex:int,pagesize:int):
    print(pageindex ,pagesize)
    data= crud_pb2.GetItemsRequest(entity=entity , pageIndex=pageindex ,pageSize= pagesize,Bucket= "ali")
    response= stub.GetItems(data)
    ls=[]
    for row in response.data:
        ls.append(row.data)
    return {"Ok":response.ok ,"data":ls} 



if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)