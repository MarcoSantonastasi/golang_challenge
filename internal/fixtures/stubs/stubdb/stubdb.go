package stubdb

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"runtime"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type StubDb struct {
}

func (db *StubDb) Connect() {}

func (db *StubDb) Close() {}

func (db *StubDb) GetAllInvestors() []*pb.Investor {
	return loadFakeInvestorsData()
}

func (db *StubDb) GetAllIssuers() []*pb.Issuer {
	return loadFakeIssuersData()
}

func (db *StubDb) GetAllInvoices() []*pb.Invoice {
	return loadFakeInvoicesData()
}

func (db *StubDb) NewInvoice(*pb.Invoice) *pb.Invoice {
	return loadFakeNewInvoiceData()
}

func loadFakeInvestorsData() (allInvestorsList []*pb.Investor) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/../../..", "/fixtures/data", "fakeInvestors.json")

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

func loadFakeIssuersData() (allIssuersList []*pb.Issuer) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/../../..", "/fixtures/data", "fakeIssuers.json")

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

func loadFakeInvoicesData() (allInvoicesList []*pb.Invoice) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/../../..", "/fixtures/data", "fakeInvoices.json")

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

func loadFakeNewInvoiceData() (newInvoiceData *pb.Invoice) {
	_, runner, _, _ := runtime.Caller(0)
	dataFile := path.Join(path.Dir(runner), "/../../..", "/fixtures/data", "newInvoice.json")

	invoiceFile, invoiceFileErr := os.Open(dataFile)
	if invoiceFileErr != nil {
		panic("cannot open " + dataFile)
	}
	defer invoiceFile.Close()

	invoiceData, invoiceDataErr := io.ReadAll(invoiceFile)
	if invoiceDataErr != nil {
		panic("cannot read " + dataFile)
	}

	invoiceJsonErr := json.Unmarshal(invoiceData, &newInvoiceData)
	if invoiceJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + dataFile)
	}
	return
}
