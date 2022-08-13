package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

var SeededAllInvestorsList = new([]*pb.Investor)
var FakeAllInvestorsList = new([]*pb.Investor)
var SeededAllIssuersList = new([]*pb.Issuer)
var FakeAllIssuersList = new([]*pb.Issuer)
var SeededAllInvoicesList = new([]*pb.Invoice)
var FakeAllInvoicesList = new([]*pb.Invoice)
var NewInvoiceData = new(pb.Invoice)

func init() {
	loadFixtureDataJson("seededInvestors.json", SeededAllInvestorsList)
	loadFixtureDataJson("fakeInvestors.json", FakeAllInvestorsList)
	loadFixtureDataJson("seededIssuers.json", SeededAllIssuersList)
	loadFixtureDataJson("fakeIssuers.json", FakeAllIssuersList)
	loadFixtureDataJson("seededInvoices.json", SeededAllInvoicesList)
	loadFixtureDataJson("fakeInvoices.json", FakeAllInvoicesList)
	loadFixtureDataJson("newInvoice.json", NewInvoiceData)
}

func loadFixtureDataJson(fileName string, dataVar any) {
	_, b, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(b), fileName)

	fmt.Println("fileName, filePath", fileName, filePath)

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		panic("cannot open " + filePath)
	}
	defer file.Close()

	data, dataErr := io.ReadAll(file)
	if dataErr != nil {
		panic("cannot read " + filePath)
	}

	jsonErr := json.Unmarshal(data, dataVar)
	if jsonErr != nil {
		panic("cannot parse (unmarshall) JSON data from newInvoice.json")
	}
}
