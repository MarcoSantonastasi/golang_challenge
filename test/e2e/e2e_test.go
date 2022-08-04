package e2etest

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	client "github.com/marcosantonastasi/arex_challenge/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AllInvestorsList []*pb.Investor
var AllIssuersList []*pb.Issuer
var AllInvoicesList []*pb.Invoice

var InvestorServiceClient client.InvestorServiceClient
var IssuerServiceClient client.IssuerServiceClient
var InvoiceServiceClient client.InvoiceServiceClient

var (
	addr       = flag.String("addr", "localhost:50051", "the address to connect to")
	globalConn *grpc.ClientConn
)

func TestE2E_GetAllInvestors(t *testing.T) {
	testCases := []struct {
		desc    string
		client  client.InvestorServiceClient
		want    pb.GetAllInvestorsResponse
		wantErr bool
	}{
		{
			desc:    "queries the db for Investors",
			client:  InvestorServiceClient,
			want:    pb.GetAllInvestorsResponse{Data: AllInvestorsList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := tt.client.GetAllInvestors()
			if (err != nil) != tt.wantErr {
				t.Errorf("Got InvestorServiceServer.GetAllInvestors() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got InvestorServiceServer.GetAllInvestors() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestE2E_GetAllIssuers(t *testing.T) {
	testCases := []struct {
		desc    string
		client  client.IssuerServiceClient
		want    pb.GetAllIssuersResponse
		wantErr bool
	}{
		{
			desc:    "queries the db for Issuers",
			client:  IssuerServiceClient,
			want:    pb.GetAllIssuersResponse{Data: AllIssuersList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := tt.client.GetAllIssuers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Got IssuerServiceServer.GetAllIssuers() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got IssuersServiceServer.GetAllIssuers() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestE2E_GetAllInvoices(t *testing.T) {
	testCases := []struct {
		desc    string
		client  client.InvoiceServiceClient
		want    pb.GetAllInvoicesResponse
		wantErr bool
	}{
		{
			desc:    "queries the db for Invoices",
			client:  InvoiceServiceClient,
			want:    pb.GetAllInvoicesResponse{Data: AllInvoicesList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := tt.client.GetAllInvoices()
			if (err != nil) != tt.wantErr {
				t.Errorf("Got InvoicersServiceServer.GetAllInvoices() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got InvoicesServiceServer.GetAllInvoices() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}

func TestMain(m *testing.M) {
	_, b, _, _ := runtime.Caller(0)
	path := path.Dir(b)

	investorsFile, investorsFileErr := os.Open(path + "/../../data/investors.json")
	if investorsFileErr != nil {
		panic("cannot open " + path + "/../../data/investors.json")
	}
	investorsData, investorsDataErr := io.ReadAll(investorsFile)
	if investorsDataErr != nil {
		panic("cannot read " + path + "/../../data/investors.json")
	}
	investorsJsonErr := json.Unmarshal(investorsData, &AllInvestorsList)
	if investorsJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/../../data/investors.json")
	}
	defer investorsFile.Close()

	issuersFile, issuersFileErr := os.Open(path + "/../../data/issuers.json")
	if issuersFileErr != nil {
		panic("cannot open " + path + "/../../data/issuers.json")
	}
	issuersData, issuersDataErr := io.ReadAll(issuersFile)
	if issuersDataErr != nil {
		panic("cannot read " + path + "/../../data/issuers.json")
	}
	issuersJsonErr := json.Unmarshal(issuersData, &AllIssuersList)
	if issuersJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data form " + path + "/../../data/issuers.json")
	}
	defer issuersFile.Close()

	invoicesFile, invoicesFileErr := os.Open(path + "/../../data/invoices.json")
	if invoicesFileErr != nil {
		panic("cannot open " + path + "/../../data/invoices.json")
	}
	invoicesData, invoicesDataErr := io.ReadAll(invoicesFile)
	if invoicesDataErr != nil {
		panic("cannot read " + path + "/../../data/invoices.json")
	}
	invoicesJsonErr := json.Unmarshal(invoicesData, &AllInvoicesList)
	if invoicesJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/../../data/invoices.json")
	}
	defer invoicesFile.Close()

	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not not connect to the gRPC server: %v", err)
	}

	globalConn = conn
	defer conn.Close()

	InvestorServiceClient = client.NewInvestorServiceClient(conn)
	InvoiceServiceClient = client.NewInvoiceServiceClient(conn)
	IssuerServiceClient = client.NewIssuerServiceClient(conn)

	code := m.Run()

	os.Exit(code)
}
