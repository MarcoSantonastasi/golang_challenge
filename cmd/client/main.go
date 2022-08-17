package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	client "github.com/marcosantonastasi/arex_challenge/internal/client"
	"github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("\ncould not not connect to the gRPC server: %+v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	investorServiceClient := client.NewInvestorServiceClient(conn)
	issuerServiceClient := client.NewIssuerServiceClient(conn)
	invoiceServiceClient := client.NewInvoiceServiceClient(conn)
	bidServiceClient := client.NewBidServiceClient(conn)

	resGetAllInvestors, errGetAllInvestors := investorServiceClient.GetAllInvestors(ctx, &pb.Empty{})
	if errGetAllInvestors != nil {
		log.Fatalf("\ncould not get Investors: %+v", errGetAllInvestors)
	}
	log.Printf("\nAll Investors: %+v", resGetAllInvestors.GetData())

	resGetAllIssuers, errGetAllIssuers := issuerServiceClient.GetAllIssuers(ctx, &pb.Empty{})
	if errGetAllIssuers != nil {
		log.Fatalf("\ncould not get Issuers: %+v", errGetAllIssuers)
	}
	log.Printf("\nAll Issuers: %+v", resGetAllIssuers.GetData())

	resGetAllInvoices, errGetAllInvoices := invoiceServiceClient.GetAllInvoices(ctx, &pb.Empty{})
	if errGetAllInvoices != nil {
		log.Fatalf("\ncould not get Invoices: %+v", errGetAllInvoices)
	}
	log.Printf("\nAll Invoices: %+v", resGetAllInvoices.GetData())

	resNewInvoice, errNewInvoice := invoiceServiceClient.NewInvoice(ctx, &pb.NewInvoiceRequest{
		IssuerAccountId: data.NewInvoiceData.IssuerAccountId,
		Reference:       data.NewInvoiceData.Reference,
		Denom:           data.NewInvoiceData.Denom,
		Amount:          data.NewInvoiceData.Amount,
		Asking:          data.NewInvoiceData.Asking,
	})
	if errNewInvoice != nil {
		log.Fatalf("\ncould not create new Invoice: %+v", errNewInvoice)
	}
	log.Printf("\nNew Invoice: %+v", resNewInvoice.GetData())

	resGetAllBids, errGetAllBids := bidServiceClient.GetAllBids(ctx, &pb.Empty{})
	if errGetAllBids != nil {
		log.Fatalf("\ncould not get Bids: %+v", errGetAllBids)
	}
	log.Printf("\nAll Bids: %+v", resGetAllBids.GetData())

	resGetAllBidsWithInvoice, errGetAllBidsWithInvoice := bidServiceClient.GetBidWithInvoiceById(
		ctx,
		&pb.GetBidWithInvoiceByIdRequest{
			BidId: resGetAllBids.GetData()[0].Id})
	if errGetAllBidsWithInvoice != nil {
		log.Fatalf("\ncould not get BidsWithInvoices: %+v", errGetAllBidsWithInvoice)
	}
	log.Printf("\nAll Bids: %+v", resGetAllBidsWithInvoice.GetData())

	resBid, errBid := bidServiceClient.NewBid(ctx, &pb.NewBidRequest{
		InvoiceId:       data.NewBidData.InvoiceId,
		BidderAccountId: data.NewBidData.BidderAccountId,
		Offer:           data.NewBidData.Offer,
	})
	if errBid != nil {
		log.Fatalf("\ncould not Bid: %+v", errBid)
	}
	log.Printf("\nBid: %+v", resBid.GetData())

	resAdj, errAdj := bidServiceClient.AdjudicateBid(ctx, &pb.AdjudicateBidRequest{
		BidId: 1,
	})
	if errAdj != nil {
		log.Fatalf("\ncould not AdjudicateBid: %+v", errBid)
	}
	log.Printf("\nAdjudicated: %+v", resAdj.GetAmount())
}
