package fakerepos

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

	fakeInvestorsFile, fakeInvestorsFileErr := os.Open(path + "/fakeInvestors.json")
	if fakeInvestorsFileErr != nil {
		panic("cannot open " + path + "/fakeInvestors.json")
	}
	fakeInvestorsData, fakeInvestorsDataErr := io.ReadAll(fakeInvestorsFile)
	if fakeInvestorsDataErr != nil {
		panic("cannot read " + path + "/fakeInvestors.json")
	}
	fakeInvestorsJsonErr := json.Unmarshal(fakeInvestorsData, &FakeAllInvestorsList)
	if fakeInvestorsJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/fakeInvestors.json")
	}
	defer fakeInvestorsFile.Close()

	fakeIssuersFile, fakeIssuersFileErr := os.Open(path + "/fakeIssuers.json")
	if fakeIssuersFileErr != nil {
		panic("cannot open " + path + "/fakeIssuers.json")
	}
	fakeIssuersData, fakeIssuersDataErr := io.ReadAll(fakeIssuersFile)
	if fakeIssuersDataErr != nil {
		panic("cannot read " + path + "/fakeIssuers.json")
	}
	fakeIssuersJsonErr := json.Unmarshal(fakeIssuersData, &FakeAllInvestorsList)
	if fakeIssuersJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data form " + path + "/fakeIssuers.json")
	}
	defer fakeIssuersFile.Close()

	fakeInvoicesFile, fakeInvoicesFileErr := os.Open(path + "/fakeInvoices.json")
	if fakeInvoicesFileErr != nil {
		panic("cannot open " + path + "/fakeInvoices.json")
	}
	fakeInvoicesData, fakeInvoicesDataErr := io.ReadAll(fakeInvoicesFile)
	if fakeInvoicesDataErr != nil {
		panic("cannot read " + path + "/fakeInvoices.json")
	}
	fakeInvoicesJsonErr := json.Unmarshal(fakeInvoicesData, &FakeAllInvoicesList)
	if fakeInvoicesJsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from " + path + "/fakeInvoices.json")
	}
	defer fakeInvoicesFile.Close()
}

type FakeInvestorsRepository struct {
}

func (repo *FakeInvestorsRepository) GetAllInvestors() ([]*pb.Investor, error) {
	return FakeAllInvestorsList, nil
}

type FakeIssuersRepository struct {
}

func (repo *FakeIssuersRepository) GetAllIssuers() ([]*pb.Issuer, error) {
	return FakeAllIssuersList, nil
}

type FakeInvoicesRepository struct {
}

func (repo *FakeInvoicesRepository) GetAllInvoices() ([]*pb.Invoice, error) {
	return FakeAllInvoicesList, nil
}
