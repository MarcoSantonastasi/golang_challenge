package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcosantonastasi/arex_challenge/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedInvoiceServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetAll(ctx context.Context, in *pb.Empty) (*pb.GetAllInvoicesResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.GetAllInvoicesResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInvoiceServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
