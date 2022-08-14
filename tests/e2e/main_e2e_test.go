package e2etest

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	client "github.com/marcosantonastasi/arex_challenge/internal/client"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	repos "github.com/marcosantonastasi/arex_challenge/internal/repos"
	server "github.com/marcosantonastasi/arex_challenge/internal/server"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

var clientServices = struct {
	investor client.InvestorServiceClient
	issuer   client.IssuerServiceClient
	invoice  client.InvoiceServiceClient
	bid      client.BidServiceClient
}{}

func TestMain(m *testing.M) {

	s := startServer()
	defer s.Stop()

	c := startClient()
	defer c.Close()

	os.Exit(m.Run())
}

func startServer() (s *grpc.Server) {
	const port int = 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pgUser := os.Getenv("POSTGRES_USER")
	pgPwd := os.Getenv("POSTGRES_PASSWORD")
	pgHostName := os.Getenv("POSTGRES_HOSTNAME")
	pgDbName := os.Getenv("POSTGRES_TEST_DB")

	dockerPgDb := db.NewPgDb(pgUser, pgPwd, pgHostName, pgDbName)

	dockerPgDb.Connect()

	s = grpc.NewServer()
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{Db: dockerPgDb}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: &repos.IssuersRepository{Db: dockerPgDb}})
	pb.RegisterInvoiceServiceServer(s, &server.InvoiceServiceServer{Repo: &repos.InvoicesRepository{Db: dockerPgDb}})
	pb.RegisterBidServiceServer(s, &server.BidServiceServer{Repo: &repos.BidsRepository{Db: dockerPgDb}})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			s.Stop()
		}
	}()

	return s
}

func startClient() (conn *grpc.ClientConn) {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not not connect to the gRPC server: %v", err)
		conn.Close()
		return nil
	}

	clientServices.investor = client.NewInvestorServiceClient(conn)
	clientServices.issuer = client.NewIssuerServiceClient(conn)
	clientServices.invoice = client.NewInvoiceServiceClient(conn)
	clientServices.bid = client.NewBidServiceClient(conn)

	return conn

}
