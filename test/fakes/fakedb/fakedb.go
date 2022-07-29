package fakedb

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"runtime"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

var FakeAllInvestorsList []*pb.Investor
var FakeAllIssuersList []*pb.Issuer
var FakeAllInvoicesList []*pb.Invoice

func init() {
	_, b, _, _ := runtime.Caller(0)
	path := path.Dir(b)

	fakeInvestorsFile, fakeInvestorsFileErr := os.Open(path + "/../../data/fakeInvestors.json")
	if fakeInvestorsFileErr != nil {
		panic("cannot open " + path + "/../../data/fakeInvestors.json")
	}
	fakeInvestorsData, fakeInvestorsDataErr := io.ReadAll(fakeInvestorsFile)
	if fakeInvestorsDataErr != nil {
		panic("cannot read " + path + "/../../data/fakeInvestors.json")
	}
	fakeInvestorsJsonErr := json.Unmarshal(fakeInvestorsData, &FakeAllInvestorsList)
	if fakeInvestorsJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/../../data/fakeInvestors.json")
	}
	defer fakeInvestorsFile.Close()

	fakeIssuersFile, fakeIssuersFileErr := os.Open(path + "/../../data/fakeIssuers.json")
	if fakeIssuersFileErr != nil {
		panic("cannot open " + path + "/../../data/fakeIssuers.json")
	}
	fakeIssuersData, fakeIssuersDataErr := io.ReadAll(fakeIssuersFile)
	if fakeIssuersDataErr != nil {
		panic("cannot read " + path + "/../../data/fakeIssuers.json")
	}
	fakeIssuersJsonErr := json.Unmarshal(fakeIssuersData, &FakeAllInvestorsList)
	if fakeIssuersJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data form " + path + "/../../data/fakeIssuers.json")
	}
	defer fakeIssuersFile.Close()

	fakeInvoicesFile, fakeInvoicesFileErr := os.Open(path + "/../../data/fakeInvoices.json")
	if fakeInvoicesFileErr != nil {
		panic("cannot open " + path + "/../../data/fakeInvoices.json")
	}
	fakeInvoicesData, fakeInvoicesDataErr := io.ReadAll(fakeInvoicesFile)
	if fakeInvoicesDataErr != nil {
		panic("cannot read " + path + "/../../data/fakeInvoices.json")
	}
	fakeInvoicesJsonErr := json.Unmarshal(fakeInvoicesData, &FakeAllInvoicesList)
	if fakeInvoicesJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/../../data/fakeInvoices.json")
	}
	defer fakeInvoicesFile.Close()
}

type FakeDb struct {
}

func (db *FakeDb) GetAllInvestors() []*pb.Investor {
	return FakeAllInvestorsList
}
func (db *FakeDb) GetAllIssuers() []*pb.Issuer {
	return FakeAllIssuersList
}
func (db *FakeDb) GetAllInvoices() []*pb.Invoice {
	return FakeAllInvoicesList
}
