package e2etest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	client "github.com/marcosantonastasi/arex_challenge/internal/client"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	repos "github.com/marcosantonastasi/arex_challenge/internal/repos"
	server "github.com/marcosantonastasi/arex_challenge/internal/server"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

var expectData = struct {
	allInvestorsList []*pb.Investor
	allIssuersList   []*pb.Issuer
	allInvoicesList  []*pb.Invoice
}{}

var clientServices = struct {
	investor client.InvestorServiceClient
	issuer   client.IssuerServiceClient
	invoice  client.InvoiceServiceClient
}{}

func TestE2E_GetAllInvestors(t *testing.T) {
	testCases := []struct {
		desc    string
		client  pb.InvestorServiceClient
		want    pb.GetAllInvestorsResponse
		wantErr bool
	}{
		{
			desc:    "queries the db for Investors",
			client:  clientServices.investor,
			want:    pb.GetAllInvestorsResponse{Data: expectData.allInvestorsList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			log.Printf("running %s", tt.desc)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllInvestors(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllInvestors() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got GetAllInvestors() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestE2E_GetAllIssuers(t *testing.T) {
	testCases := []struct {
		desc    string
		client  pb.IssuerServiceClient
		want    pb.GetAllIssuersResponse
		wantErr bool
	}{
		{
			desc:    "queries the db for Issuers",
			client:  clientServices.issuer,
			want:    pb.GetAllIssuersResponse{Data: expectData.allIssuersList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			log.Printf("running %s", tt.desc)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllIssuers(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllIssuers() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got GetAllIssuers() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestE2E_GetAllInvoices(t *testing.T) {
	testCases := []struct {
		desc    string
		client  pb.InvoiceServiceClient
		want    pb.GetAllInvoicesResponse
		wantErr bool
	}{
		{
			desc:    "queries the db for Invoices",
			client:  clientServices.invoice,
			want:    pb.GetAllInvoicesResponse{Data: expectData.allInvoicesList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			log.Printf("running %s", tt.desc)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllInvoices(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllInvoices() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got GetAllInvoices() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestMain(m *testing.M) {

	log.Printf("TestMain()")

	loadTestData()

	s := startBufferedServer()
	defer s.Stop()

	conn := startBufferedClient()
	defer conn.Close()
	
	os.Exit(m.Run())
}

func loadTestData() {
	log.Printf("loadTestData()")
	
	_, b, _, _ := runtime.Caller(0)
	path := path.Dir(b)
	
	investorsFile, investorsFileErr := os.Open(path + "/../data/investors.json")
	if investorsFileErr != nil {
		panic("cannot open " + path + "/../data/investors.json")
	}
	investorsData, investorsDataErr := io.ReadAll(investorsFile)
	if investorsDataErr != nil {
		panic("cannot read " + path + "/../data/investors.json")
	}
	investorsJsonErr := json.Unmarshal(investorsData, &expectData.allInvestorsList)
	if investorsJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/../data/investors.json")
	}
	defer investorsFile.Close()
	
	issuersFile, issuersFileErr := os.Open(path + "/../data/issuers.json")
	if issuersFileErr != nil {
		panic("cannot open " + path + "/../data/issuers.json")
	}
	issuersData, issuersDataErr := io.ReadAll(issuersFile)
	if issuersDataErr != nil {
		panic("cannot read " + path + "/../data/issuers.json")
	}
	issuersJsonErr := json.Unmarshal(issuersData, &expectData.allIssuersList)
	if issuersJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data form " + path + "/../data/issuers.json")
	}
	defer issuersFile.Close()
	
	invoicesFile, invoicesFileErr := os.Open(path + "/../data/invoices.json")
	if invoicesFileErr != nil {
		panic("cannot open " + path + "/../data/invoices.json")
	}
	invoicesData, invoicesDataErr := io.ReadAll(invoicesFile)
	if invoicesDataErr != nil {
		panic("cannot read " + path + "/../data/invoices.json")
	}
	invoicesJsonErr := json.Unmarshal(invoicesData, &expectData.allInvoicesList)
	if invoicesJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/../data/invoices.json")
	}
	defer invoicesFile.Close()
	
}

func startBufferedServer() (s *grpc.Server) {
	
	log.Printf("startBufferedServer()")
	
	const port int = 50051
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s = grpc.NewServer()
	
	pb.RegisterInvestorServiceServer(s, &server.InvestorServiceServer{Repo: &repos.InvestorsRepository{Db: &db.Db{Conn: db.DockerPG.Conn}}})
	pb.RegisterIssuerServiceServer(s, &server.IssuerServiceServer{Repo: &repos.IssuersRepository{Db: &db.Db{Conn: db.DockerPG.Conn}}})
	pb.RegisterInvoiceServiceServer(s, &server.InvoiceServiceServer{Repo: &repos.InvoicesRepository{Db: &db.Db{Conn: db.DockerPG.Conn}}})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		s.Stop()
	}
	}()
	
	return s
}

func startBufferedClient() (conn *grpc.ClientConn) {
	
	log.Printf("startBufferedClient()")
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
