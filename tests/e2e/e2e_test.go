package e2etest

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	client "github.com/marcosantonastasi/arex_challenge/internal/client"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
	repos "github.com/marcosantonastasi/arex_challenge/internal/repos"
	server "github.com/marcosantonastasi/arex_challenge/internal/server"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

var clientServices = struct {
	investor client.InvestorServiceClient
	issuer   client.IssuerServiceClient
	invoice  client.InvoiceServiceClient
}{}

func TestE2E_GetAllInvestors(t *testing.T) {
	tests := []struct {
		desc    string
		client  pb.InvestorServiceClient
		want    *pb.GetAllInvestorsResponse
		wantErr bool
	}{
		{
			desc:    "gets the list of all Investors",
			client:  clientServices.investor,
			want:    &pb.GetAllInvestorsResponse{Data: *data.SeededAllInvestorsList},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllInvestors(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllInvestors() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("Got GetAllInvestors() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestE2E_GetAllIssuers(t *testing.T) {
	testCases := []struct {
		desc    string
		client  pb.IssuerServiceClient
		want    *pb.GetAllIssuersResponse
		wantErr bool
	}{
		{
			desc:    "gets the list of all Issuers",
			client:  clientServices.issuer,
			want:    &pb.GetAllIssuersResponse{Data: *data.SeededAllIssuersList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllIssuers(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllIssuers() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("Got GetAllIssuers() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestE2E_GetAllInvoices(t *testing.T) {
	testCases := []struct {
		desc    string
		client  pb.InvoiceServiceClient
		want    *pb.GetAllInvoicesResponse
		wantErr bool
	}{
		{
			desc:    "gets the list of all Invoices",
			client:  clientServices.invoice,
			want:    &pb.GetAllInvoicesResponse{Data: *data.SeededAllInvoicesList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllInvoices(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllInvoices() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("Got GetAllInvoices() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestMain(m *testing.M) {

	s := startServer()
	defer s.Stop()

	c := startClient()
	defer c.Close()

	fmt.Println(clientServices.investor, clientServices.issuer, clientServices.invoice)

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
	pgHostname := os.Getenv("POSTGRES_HOSTNAME")
	pgDbname := os.Getenv("POSTGRES_DB")

	dockerPgDb := db.NewPgDb(pgUser, pgPwd, pgHostname, pgDbname)

	dockerPgDb.Connect()
	defer dockerPgDb.Close()

	s = grpc.NewServer()
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{Db: dockerPgDb}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: &repos.IssuersRepository{Db: dockerPgDb}})
	pb.RegisterInvoiceServiceServer(s, &server.InvoiceServiceServer{Repo: &repos.InvoicesRepository{Db: dockerPgDb}})
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

	return conn

}
