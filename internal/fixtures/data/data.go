package data

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"runtime"

	db "github.com/marcosantonastasi/arex_challenge/internal/db"
)

var SeededAllInvestorsList = new([]*db.Account)
var FakeAllInvestorsList = new([]*db.Account)
var SeededAllIssuersList = new([]*db.Account)
var FakeAllIssuersList = new([]*db.Account)
var SeededAllInvoicesList = new([]*db.Invoice)
var FakeAllInvoicesList = new([]*db.Invoice)
var NewInvoiceData = new(db.Invoice)
var SeededAllBidsList = new([]*db.Bid)
var FakeAllBidsList = new([]*db.Bid)
var NewBidData = new(db.Bid)
var AdjudicateBidData = new(struct {
	Id     int64
	Amount int64
})

func init() {
	loadFixtureDataJson("seededInvestors.json", SeededAllInvestorsList)
	loadFixtureDataJson("fakeInvestors.json", FakeAllInvestorsList)
	loadFixtureDataJson("seededIssuers.json", SeededAllIssuersList)
	loadFixtureDataJson("fakeIssuers.json", FakeAllIssuersList)
	loadFixtureDataJson("seededInvoices.json", SeededAllInvoicesList)
	loadFixtureDataJson("fakeInvoices.json", FakeAllInvoicesList)
	loadFixtureDataJson("newInvoice.json", NewInvoiceData)
	loadFixtureDataJson("seededBids.json", SeededAllBidsList)
	loadFixtureDataJson("fakeBids.json", FakeAllBidsList)
	loadFixtureDataJson("newBid.json", NewBidData)
	loadFixtureDataJson("adjudicateBid.json", AdjudicateBidData)
}

func loadFixtureDataJson(fileName string, dataVar any) {
	_, b, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(b), fileName)

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
		panic("cannot parse (unmarshall) JSON data from " + filePath)
	}
}
