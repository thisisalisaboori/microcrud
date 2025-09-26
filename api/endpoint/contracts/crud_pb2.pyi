from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BaseResponse(_message.Message):
    __slots__ = ("ok",)
    OK_FIELD_NUMBER: _ClassVar[int]
    ok: bool
    def __init__(self, ok: bool = ...) -> None: ...

class InitRequst(_message.Message):
    __slots__ = ("Bucket", "Collection", "CreateIndex")
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    COLLECTION_FIELD_NUMBER: _ClassVar[int]
    CREATEINDEX_FIELD_NUMBER: _ClassVar[int]
    Bucket: str
    Collection: str
    CreateIndex: bool
    def __init__(self, Bucket: _Optional[str] = ..., Collection: _Optional[str] = ..., CreateIndex: bool = ...) -> None: ...

class GetByIdResponse(_message.Message):
    __slots__ = ("ok", "data")
    OK_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    ok: bool
    data: _struct_pb2.Struct
    def __init__(self, ok: bool = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class GetItemsResponse(_message.Message):
    __slots__ = ("ok", "data")
    OK_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    ok: bool
    data: _containers.RepeatedCompositeFieldContainer[GetByIdResponse]
    def __init__(self, ok: bool = ..., data: _Optional[_Iterable[_Union[GetByIdResponse, _Mapping]]] = ...) -> None: ...

class CreateItemRequest(_message.Message):
    __slots__ = ("Bucket", "entity", "data")
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    Bucket: str
    entity: str
    data: _struct_pb2.Struct
    def __init__(self, Bucket: _Optional[str] = ..., entity: _Optional[str] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class UpdateItemRequest(_message.Message):
    __slots__ = ("Bucket", "id", "entity", "data")
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    Bucket: str
    id: str
    entity: str
    data: _struct_pb2.Struct
    def __init__(self, Bucket: _Optional[str] = ..., id: _Optional[str] = ..., entity: _Optional[str] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class DeleteItemRequest(_message.Message):
    __slots__ = ("Bucket", "id", "entity")
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    Bucket: str
    id: str
    entity: str
    def __init__(self, Bucket: _Optional[str] = ..., id: _Optional[str] = ..., entity: _Optional[str] = ...) -> None: ...

class GetItemRequest(_message.Message):
    __slots__ = ("Bucket", "id", "entity")
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    Bucket: str
    id: str
    entity: str
    def __init__(self, Bucket: _Optional[str] = ..., id: _Optional[str] = ..., entity: _Optional[str] = ...) -> None: ...

class GetItemsRequest(_message.Message):
    __slots__ = ("Bucket", "entity", "pageIndex", "pageSize")
    BUCKET_FIELD_NUMBER: _ClassVar[int]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    PAGEINDEX_FIELD_NUMBER: _ClassVar[int]
    PAGESIZE_FIELD_NUMBER: _ClassVar[int]
    Bucket: str
    entity: str
    pageIndex: int
    pageSize: int
    def __init__(self, Bucket: _Optional[str] = ..., entity: _Optional[str] = ..., pageIndex: _Optional[int] = ..., pageSize: _Optional[int] = ...) -> None: ...
