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
var SeededAllIssuersList = new([]*db.Account)
var SeededAllInvoicesList = new([]*db.Invoice)
var NewInvoiceData = new(db.Invoice)
var SeededAllBidsList = new([]*db.Bid)
var NewBidData = new(db.Bid)

var RequestGetAllInvestors = &pb.Empty{}
var RequestGetAllIssuers = &pb.Empty{}
var RequestGetAllInvoices = &pb.Empty{}
var RequestNewInvoice = &pb.NewInvoiceRequest{}
var RequestGetAllBids = &pb.Empty{}
var RequesteNewBid = new(pb.NewBidRequest)

var ResponseGetAllInvestors = new(pb.GetAllInvestorsResponse)
var ResponseGetAllIssuers = new(pb.GetAllIssuersResponse)
var ResponseGetAllInvoices = new(pb.GetAllInvoicesResponse)
var ResponseNewInvoice = new(pb.NewInvoiceResponse)
var ResponseGetAllBids = new(pb.GetAllBidsResponse)
var ResponseNewBid = new(pb.NewBidResponse)

func init() {
	loadFixtureDataJson("seededInvestors.json", SeededAllInvestorsList)
	loadFixtureDataJson("seededIssuers.json", SeededAllIssuersList)
	loadFixtureDataJson("seededInvoices.json", SeededAllInvoicesList)
	loadFixtureDataJson("newInvoice.json", NewInvoiceData)
	loadFixtureDataJson("seededBids.json", SeededAllBidsList)
	loadFixtureDataJson("newBid.json", NewBidData)

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

	RequestNewInvoice = &pb.NewInvoiceRequest{
		IssuerAccountId: NewInvoiceData.IssuerAccountId,
		Reference:       NewInvoiceData.Reference,
		Denom:           NewInvoiceData.Denom,
		Amount:          NewInvoiceData.Amount,
		Asking:          NewInvoiceData.Asking,
	}

	RequesteNewBid = &pb.NewBidRequest{
		InvoiceId:       NewBidData.InvoiceId,
		BidderAccountId: NewBidData.BidderAccountId,
		Offer:           NewBidData.Offer,
	}

	for _, i := range *SeededAllInvestorsList {
		ResponseGetAllInvestors.Data = append(ResponseGetAllInvestors.Data, &pb.Investor{
			Id:      i.Id,
			Name:    i.Name,
			Balance: i.Balance,
		})
	}

	for _, i := range *SeededAllIssuersList {
		ResponseGetAllIssuers.Data = append(ResponseGetAllIssuers.Data, &pb.Issuer{
			Id:      i.Id,
			Name:    i.Name,
			Balance: i.Balance,
		})
	}

	for _, b := range *SeededAllBidsList {
		ResponseGetAllBids.Data = append(ResponseGetAllBids.Data, &pb.Bid{
			Id:              b.Id,
			InvoiceId:       b.InvoiceId,
			BidderAccountId: b.BidderAccountId,
			Offer:           b.Offer,
			State:           b.State,
		})
	}

	for _, i := range *SeededAllInvoicesList {
		ResponseGetAllInvoices.Data = append(ResponseGetAllInvoices.Data, &pb.Invoice{
			Id:              i.Id,
			IssuerAccountId: i.IssuerAccountId,
			Reference:       i.Reference,
			Denom:           i.Denom,
			Amount:          i.Amount,
			Asking:          i.Asking,
			State:           i.State,
		})
	}

	ResponseNewInvoice.Data = &pb.Invoice{
		Id:              NewInvoiceData.Id,
		IssuerAccountId: NewInvoiceData.IssuerAccountId,
		Reference:       NewInvoiceData.Reference,
		Denom:           NewInvoiceData.Denom,
		Amount:          NewInvoiceData.Amount,
		Asking:          NewInvoiceData.Asking,
		State:           NewInvoiceData.State,
	}

	ResponseNewBid.Data = &pb.Bid{
		Id:              NewBidData.Id,
		InvoiceId:       NewBidData.InvoiceId,
		BidderAccountId: NewBidData.BidderAccountId,
		Offer:           NewBidData.Offer,
		State:           NewBidData.State,
	}
}
