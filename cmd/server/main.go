package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	server "github.com/marcosantonastasi/arex_challenge/pkg/server"
	repos "github.com/marcosantonastasi/arex_challenge/repos"
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
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: repos.IssuersRepository{}})
	pb.RegisterInvoiceServiceServer(s, server.InvoiceServiceServer{Repo: repos.InvoiceRepository{}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
