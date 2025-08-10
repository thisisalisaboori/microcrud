package microcrud

import (
	"context"
	"fmt"


	// "log"
	// "net"

	// "google.golang.org/grpc"

	"github.com/thisisalisaboori/microcrud/microcrud"
	"google.golang.org/protobuf/types/known/structpb"
)




type crudEntityServer struct{}


func (s *crudEntityServer) Create(ctx context.Context, req *EntityProtocol) (*ResultService, error) {
	fmt.Println("Received Create request:", req.Data)

	respData, _ := structpb.NewStruct(map[string]interface{}{
		"status": "created",
		"id":     "12345",
	})

	return &ResultService{
		Message: "Entity created successfully",
		Data:    respData,
	}, nil
}

func (s *crudEntityServer) Update(ctx context.Context, req *EntityProtocol) (*ResultService, error) {
	return &ResultService{Message: "Update not implemented"}, nil
}

func (s *crudEntityServer) Delete(ctx context.Context, req *DeleteProtocol) (*ResultService, error) {
	return &ResultService{Message: "Delete not implemented"}, nil
}

func (s *crudEntityServer) GetById(ctx context.Context, req *GetProtocol) (*ResultService, error) {
	return &ResultService{Message: "GetById not implemented"}, nil
}

