package data

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"runtime"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
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

var ResponseGetAllInvestors = new([]*pb.Investor)
var ResponseGetAllIssuers = new([]*pb.Issuer)
var ResponseGetAllInvoices = new([]*pb.Invoice)
var ResponseNewInvoice = new(pb.Invoice)
var ResponseGetAllBids = new([]*pb.Bid)
var ResponseNewBid = new(pb.Bid)
var ResponseAdjudicateBid = new(struct {
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

	makeResponseFromDataVars()

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

func makeResponseFromDataVars() {

	for _, i := range *SeededAllInvestorsList {
		*ResponseGetAllInvestors = append(*ResponseGetAllInvestors, &pb.Investor{
			Id:      i.Id,
			Name:    i.Name,
			Balance: i.Balance,
		})
	}

	for _, i := range *SeededAllIssuersList {
		*ResponseGetAllIssuers = append(*ResponseGetAllIssuers, &pb.Issuer{
			Id:      i.Id,
			Name:    i.Name,
			Balance: i.Balance,
		})
	}

	for _, b := range *SeededAllBidsList {
		*ResponseGetAllBids = append(*ResponseGetAllBids, &pb.Bid{
			Id:              b.Id,
			InvoiceId:       b.InvoiceId,
			BidderAccountId: b.BidderAccountId,
			Offer:           b.Offer,
			State:           b.State,
		})
	}

	for _, i := range *SeededAllInvoicesList {
		*ResponseGetAllInvoices = append(*ResponseGetAllInvoices, &pb.Invoice{
			Id:              i.Id,
			IssuerAccountId: i.IssuerAccountId,
			Reference:       i.Reference,
			Denom:           i.Denom,
			Amount:          i.Amount,
			Asking:          i.Asking,
			State:           i.State,
		})
	}

	ResponseNewInvoice = ResponseNewInvoice{}

	var ResponseNewInvoice = new(pb.Invoice)
	var ResponseNewBid = new(pb.Bid)
	var ResponseAdjudicateBid = new(struct {
		Id     int64
		Amount int64
	})
}
