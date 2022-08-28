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
	const hostname string = "localhost"
	const port int = 50051

	s := startServer(hostname, port)
	defer s.Stop()

	c := startClient(hostname, port)
	defer c.Close()

	os.Exit(m.Run())
}

func startServer(hostname string, port int) (s *grpc.Server) {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPwd := os.Getenv("POSTGRES_PASSWORD")
	pgHostName := os.Getenv("POSTGRES_HOSTNAME")
	pgDbName := os.Getenv("POSTGRES_TESTING_DB")

	if pgUser == "" || pgPwd == "" || pgHostName == "" || pgDbName == "" {
		log.Fatal("e2e is missing .env variables")
	}

	dockerPgDb := db.NewPgDb(pgUser, pgPwd, pgHostName, pgDbName)

	dbErr := dockerPgDb.Connect()
	if dbErr != nil {
		log.Fatalf("e2e test conneciton to PostgresDB failed %+v", dbErr)
	}
	log.Printf("e2e test conneciton to PostgresDB %+v", dockerPgDb)

	s = grpc.NewServer()
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{Db: dockerPgDb}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: &repos.IssuersRepository{Db: dockerPgDb}})
	pb.RegisterInvoiceServiceServer(s, &server.InvoiceServiceServer{Repo: &repos.InvoicesRepository{Db: dockerPgDb}})
	pb.RegisterBidServiceServer(s, &server.BidServiceServer{Repo: &repos.BidsRepository{Db: dockerPgDb}})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		log.Fatalf("e2e test TCP failed to listen on port %d because of error %v", port, err)
	}
	log.Printf("e2e test TCP listening on %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("e2e test server failed with error: %v", err)
			s.Stop()
		}
		log.Printf("e2e test server ready on %v", lis.Addr())
	}()

	return s
}

func startClient(hostname string, port int) (conn *grpc.ClientConn) {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("e2e test client could not not connect to the gRPC e2e test server: %v", err)
		conn.Close()
		return nil
	}
	log.Printf("e2e test client ready for %v", conn.Target())

	clientServices.investor = client.NewInvestorServiceClient(conn)
	clientServices.issuer = client.NewIssuerServiceClient(conn)
	clientServices.invoice = client.NewInvoiceServiceClient(conn)
	clientServices.bid = client.NewBidServiceClient(conn)

	return conn

}
