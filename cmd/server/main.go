package main

import (
	"fmt"
	"log"
	"net"

	app "github.com/thisisalisaboori/microcrud/internal/app"
	pb "github.com/thisisalisaboori/microcrud/api/proto/microcrudproto"
	"google.golang.org/grpc"
)
func main(){
	lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterCrudServiceServer(s, &app.Server{})

    log.Println("gRPC server running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
	fmt.Println(("hello"))
}