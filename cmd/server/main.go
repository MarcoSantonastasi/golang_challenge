package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/db"
	server "github.com/marcosantonastasi/arex_challenge/pkg/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Db: db.ArexDb{}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Db: db.ArexDb{}})
	pb.RegisterInvoiceServiceServer(s, server.InvoiceServiceServer{Db: db.ArexDb{}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
