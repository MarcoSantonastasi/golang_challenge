package client

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	"google.golang.org/grpc"
)

type InvestorServiceClient interface {
	pb.InvestorServiceClient
}

func NewInvestorServiceClient(conn *grpc.ClientConn) pb.InvestorServiceClient {
	return pb.NewInvestorServiceClient(conn)
}

type IssuerServiceClient interface {
	pb.IssuerServiceClient
}

func NewIssuerServiceClient(conn *grpc.ClientConn) pb.IssuerServiceClient {
	return pb.NewIssuerServiceClient(conn)
}

type InvoiceServiceClient interface {
	pb.InvoiceServiceClient
}

func NewInvoiceServiceClient(conn *grpc.ClientConn) pb.InvoiceServiceClient {
	return pb.NewInvoiceServiceClient(conn)
}

type BidServiceClient interface {
	pb.BidServiceClient
}

func NewBidServiceClient(conn *grpc.ClientConn) pb.BidServiceClient {
	return pb.NewBidServiceClient(conn)
}
