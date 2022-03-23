package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hellokvn/go-grpc-product-svc/pkg/db"
	handlers "github.com/hellokvn/go-grpc-product-svc/pkg/handlers"
	pb "github.com/hellokvn/go-grpc-product-svc/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	DB := db.Init()
	h := handlers.New(DB)

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Printf("Run Product Svc on %v", port)

	s := handlers.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
