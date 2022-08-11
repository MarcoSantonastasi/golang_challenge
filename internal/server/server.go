package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	repos "github.com/marcosantonastasi/arex_challenge/internal/repos"
)

type InvestorServiceServer struct {
	pb.UnimplementedInvestorServiceServer
	Repo repos.IInvestorsRepository
}
type IssuerServiceServer struct {
	pb.UnimplementedIssuerServiceServer
	Repo repos.IIssuersRepository
}
type InvoiceServiceServer struct {
	pb.UnimplementedInvoiceServiceServer
	Repo repos.IInvoicesRepository
}

func (s *InvestorServiceServer) GetAllInvestors(ctx context.Context, in *pb.Empty) (*pb.GetAllInvestorsResponse, error) {
	if s.Repo == nil {
		return nil, status.Error(codes.Internal, "no repository found for Investors")
	}
	res, err := s.Repo.GetAllInvestors()
	if err != nil {
		return nil, fmt.Errorf("database error: %q", err)
	}

	return &pb.GetAllInvestorsResponse{Data: res}, nil
}

func (s *IssuerServiceServer) GetAllIssuers(ctx context.Context, in *pb.Empty) (*pb.GetAllIssuersResponse, error) {
	if s.Repo == nil {
		return nil, status.Error(codes.Internal, "no repository found for Issuers")
	}
	res, err := s.Repo.GetAllIssuers()
	if err != nil {
		return nil, fmt.Errorf("database error: %q", err)
	}
	return &pb.GetAllIssuersResponse{Data: res}, nil
}

func (s *InvoiceServiceServer) GetAllInvoices(ctx context.Context, in *pb.Empty) (*pb.GetAllInvoicesResponse, error) {
	if s.Repo == nil {
		return nil, status.Error(codes.Internal, "no repository found for Invoices")
	}
	res, err := s.Repo.GetAllInvoices()
	if err != nil {
		return nil, fmt.Errorf("database error: %q", err)
	}
	return &pb.GetAllInvoicesResponse{Data: res}, nil
}
