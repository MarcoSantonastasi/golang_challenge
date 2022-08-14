package client

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	"google.golang.org/grpc"
)

type InvestorServiceClient interface {
	pb.InvestorServiceClient
}
type IssuerServiceClient interface {
	pb.IssuerServiceClient
}
type InvoiceServiceClient interface {
	pb.InvoiceServiceClient
}

func NewInvestorServiceClient(conn *grpc.ClientConn) pb.InvestorServiceClient {
	return pb.NewInvestorServiceClient(conn)
}

func NewIssuerServiceClient(conn *grpc.ClientConn) pb.IssuerServiceClient {
	return pb.NewIssuerServiceClient(conn)
}

func NewInvoiceServiceClient(conn *grpc.ClientConn) pb.InvoiceServiceClient {
	return pb.NewInvoiceServiceClient(conn)
}

func NewBidServiceClient(conn *grpc.ClientConn) pb.BidServiceClient {
	return pb.NewBidServiceClient(conn)
}
