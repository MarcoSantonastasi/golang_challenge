package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

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

	pgUser := os.Getenv("POSTGRES_USER")
	pgPwd := os.Getenv("POSTGRES_PASSWORD")
	pgHostName := os.Getenv("POSTGRES_HOSTNAME")
	pgDbName := os.Getenv("POSTGRES_PROD_DB")

	dockerPgDb := db.NewPgDb(pgUser, pgPwd, pgHostName, pgDbName)

	dockerPgDb.Connect()
	defer dockerPgDb.Close()

	s := grpc.NewServer()
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{Db: dockerPgDb}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: &repos.IssuersRepository{Db: dockerPgDb}})
	pb.RegisterInvoiceServiceServer(s, &server.InvoiceServiceServer{Repo: &repos.InvoicesRepository{Db: dockerPgDb}})
	pb.RegisterBidServiceServer(s, &server.BidServiceServer{Repo: &repos.BidsRepository{Db: dockerPgDb}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
