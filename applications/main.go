package main

import (
	"log"
	"net"

	"github.com/thisisalisaboori/microcrud/microcrud"
	"google.golang.org/grpc"
	//pb "github.com/thisisalisaboori/microcrud"
)
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	microcrud.RegisterCrudEntityServer(grpcServer, &microcrud.crudEntityServer{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
