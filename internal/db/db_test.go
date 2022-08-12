package db_test

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
)

var testDb db.IDb

func TestMain(m *testing.M) {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPwd := os.Getenv("POSTGRES_PASSWORD")
	pgHostname := os.Getenv("POSTGRES_HOSTNAME")
	pgDbname := os.Getenv("POSTGRES_DB")

	testDb = db.NewPgDb(pgUser, pgPwd, pgHostname, pgDbname)

	testDb.Connect()

	exitCode := m.Run()

	testDb.Close()

	os.Exit(exitCode)
}

func TestDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name     string
		db       db.IDb
		wantData []*pb.Investor
	}{
		{
			name:     "gets the list of all Investors from the investors view",
			db:       testDb,
			wantData: loadSeededInvestorsData(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := tt.db.GetAllInvestors(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", gotData, tt.wantData)
			}
		})
	}
}

func TestDb_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name     string
		db       db.IDb
		wantData []*pb.Issuer
	}{
		{
			name:     "gets the list of all Issuers from the issuers view",
			db:       testDb,
			wantData: loadSeededIssuersData(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := tt.db.GetAllIssuers(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", gotData, tt.wantData)
			}
		})
	}
}

func TestDb_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name     string
		db       db.IDb
		wantData []*pb.Invoice
	}{
		{
			name:     "gets the list of all Invoices from the invoices table",
			db:       testDb,
			wantData: loadSeededInvoicesData(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := tt.db.GetAllInvoices(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", gotData, tt.wantData)
			}
		})
	}
}

func loadSeededInvestorsData() (allInvestorsList []*pb.Investor) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/..", "/fixtures/data", "seededInvestors.json")

	investorsFile, investorsFileErr := os.Open(dataFile)
	if investorsFileErr != nil {
		panic("cannot open " + dataFile)
	}
	defer investorsFile.Close()

	investorsData, investorsDataErr := io.ReadAll(investorsFile)
	if investorsDataErr != nil {
		panic("cannot read " + dataFile)
	}
	investorsJsonErr := json.Unmarshal(investorsData, &allInvestorsList)
	if investorsJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + dataFile)
	}

	return
}

func loadSeededIssuersData() (allIssuersList []*pb.Issuer) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/..", "/fixtures/data", "seededIssuers.json")

	issuersFile, issuersFileErr := os.Open(dataFile)
	if issuersFileErr != nil {
		panic("cannot open " + dataFile)
	}
	defer issuersFile.Close()

	issuersData, issuersDataErr := io.ReadAll(issuersFile)
	if issuersDataErr != nil {
		panic("cannot read " + dataFile)
	}
	issuersJsonErr := json.Unmarshal(issuersData, &allIssuersList)
	if issuersJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data form " + dataFile)
	}
	return
}

func loadSeededInvoicesData() (allInvoicesList []*pb.Invoice) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/..", "/fixtures/data", "seededInvoices.json")

	invoicesFile, invoicesFileErr := os.Open(dataFile)
	if invoicesFileErr != nil {
		panic("cannot open " + dataFile)
	}
	defer invoicesFile.Close()

	invoicesData, invoicesDataErr := io.ReadAll(invoicesFile)
	if invoicesDataErr != nil {
		panic("cannot read " + dataFile)
	}
	invoicesJsonErr := json.Unmarshal(invoicesData, &allInvoicesList)
	if invoicesJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + dataFile)
	}
	return
}
