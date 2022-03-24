package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hellokvn/go-grpc-product-svc/pkg/db"
	pb "github.com/hellokvn/go-grpc-product-svc/pkg/pb"
	services "github.com/hellokvn/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	port := ":50052"
	h := db.Init("postgres://kevin@localhost:5432/product_svc")

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
