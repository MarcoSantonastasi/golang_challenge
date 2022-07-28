package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/db"
)

type InvestorServiceServer struct {
	pb.UnimplementedInvestorServiceServer
	Db db.IArexDb
}
type IssuerServiceServer struct {
	pb.UnimplementedIssuerServiceServer
	Db db.IArexDb
}
type InvoiceServiceServer struct {
	pb.UnimplementedInvoiceServiceServer
	Db db.IArexDb
}

func (s *InvestorServiceServer) GetAllInvestors(ctx context.Context, in *pb.Empty) (*pb.GetAllInvestorsResponse, error) {
	if s.Db == nil {
		return nil, status.Error(codes.Internal, "no database connection found")
	}
	log.Printf("Received: %v", in)
	return &pb.GetAllInvestorsResponse{}, nil
}

func (s *IssuerServiceServer) GetAllIssuers(ctx context.Context, in *pb.Empty) (*pb.GetAllIssuersResponse, error) {
	if s.Db == nil {
		return nil, status.Error(codes.Internal, "no database connection found")
	}
	log.Printf("Received: %v", in)
	return &pb.GetAllIssuersResponse{}, nil
}

func (s *InvoiceServiceServer) GetAllinvoices(ctx context.Context, in *pb.Empty) (*pb.GetAllInvoicesResponse, error) {
	if s.Db == nil {
		return nil, status.Error(codes.Internal, "no database connection found")
	}
	log.Printf("Received: %v", in)
	return &pb.GetAllInvoicesResponse{}, nil
}
