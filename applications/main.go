package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	//pb "github.com/thisisalisaboori/microcrud"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterCrudEntityServer(grpcServer, &MyCrudEntityServer{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
