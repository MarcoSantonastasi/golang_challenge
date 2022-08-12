package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	repos "github.com/marcosantonastasi/arex_challenge/internal/repos"
	server "github.com/marcosantonastasi/arex_challenge/internal/server"
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
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{Db: &db.PgDb{Conn: db.DockerPG.Conn}}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: &repos.IssuersRepository{Db: &db.PgDb{Conn: db.DockerPG.Conn}}})
	pb.RegisterInvoiceServiceServer(s, &server.InvoiceServiceServer{Repo: &repos.InvoicesRepository{Db: &db.PgDb{Conn: db.DockerPG.Conn}}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
