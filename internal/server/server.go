package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
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

type BidServiceServer struct {
	pb.UnimplementedBidServiceServer
	Repo repos.IBidsRepository
}

func (s *InvestorServiceServer) GetAllInvestors(ctx context.Context, req *pb.Empty) (*pb.GetAllInvestorsResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Investors")
	}

	res, err := s.Repo.GetAllInvestors()
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetAllInvestors: %+v", err)
	}

	data := []*pb.Investor{}

	for _, i := range *res {
		data = append(data, &pb.Investor{
			Id:      i.Id,
			Name:    i.Name,
			Balance: i.Balance,
		})
	}

	return &pb.GetAllInvestorsResponse{Data: data}, nil
}

func (s *IssuerServiceServer) GetAllIssuers(ctx context.Context, req *pb.Empty) (*pb.GetAllIssuersResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Issuers")
	}

	res, err := s.Repo.GetAllIssuers()
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetAllIssuers: %+v", err)
	}

	data := []*pb.Issuer{}

	for _, i := range *res {
		data = append(data, &pb.Issuer{
			Id:      i.Id,
			Name:    i.Name,
			Balance: i.Balance,
		})
	}

	return &pb.GetAllIssuersResponse{Data: data}, nil
}

func (s *InvoiceServiceServer) GetAllInvoices(ctx context.Context, req *pb.Empty) (*pb.GetAllInvoicesResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Invoices")
	}
	res, err := s.Repo.GetAllInvoices()
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetAllInvoices: %+v", err)
	}

	data := []*pb.Invoice{}

	for _, i := range *res {
		data = append(data, &pb.Invoice{
			Id:              i.Id,
			IssuerAccountId: i.IssuerAccountId,
			Reference:       i.Reference,
			Denom:           i.Denom,
			Amount:          i.Amount,
			Asking:          i.Asking,
			State:           i.State,
		})
	}

	return &pb.GetAllInvoicesResponse{Data: data}, nil
}

func (s *InvoiceServiceServer) GetInvoiceById(ctx context.Context, req *pb.GetInvoiceByIdRequest) (*pb.GetInvoiceByIdResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Invoices")
	}

	byInvoiceId := req.InvoiceId

	res, err := s.Repo.GetInvoiceById(byInvoiceId)
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetInvoiceById: %+v", err)
	}

	data := pb.Invoice{
		Id:              res.Id,
		IssuerAccountId: res.IssuerAccountId,
		Reference:       res.Reference,
		Denom:           res.Denom,
		Amount:          res.Amount,
		Asking:          res.Asking,
		State:           res.State,
	}

	return &pb.GetInvoiceByIdResponse{Data: &data}, nil
}

func (s *InvoiceServiceServer) NewInvoice(ctx context.Context, req *pb.NewInvoiceRequest) (*pb.NewInvoiceResponse, error) {
	if s.Repo == nil {
		return nil, status.Error(codes.Internal, "no repository found for Invoices")
	}
	newInvoiceData := &db.Invoice{
		IssuerAccountId: req.IssuerAccountId,
		Reference:       req.Reference,
		Denom:           req.Denom,
		Amount:          req.Amount,
		Asking:          req.Asking,
	}
	res, err := s.Repo.NewInvoice(newInvoiceData)
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for NewInvoice: %+v", err)
	}

	data := pb.Invoice{
		Id:              res.Id,
		IssuerAccountId: res.IssuerAccountId,
		Reference:       res.Reference,
		Denom:           res.Denom,
		Amount:          res.Amount,
		Asking:          res.Asking,
		State:           res.State,
	}

	return &pb.NewInvoiceResponse{Data: &data}, nil
}

func (s *BidServiceServer) GetAllBids(ctx context.Context, req *pb.Empty) (*pb.GetAllBidsResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Bids")
	}

	res, err := s.Repo.GetAllBids()
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetAllBids: %+v", err)
	}

	data := []*pb.Bid{}

	for _, b := range *res {
		data = append(data, &pb.Bid{
			Id:              b.Id,
			InvoiceId:       b.InvoiceId,
			BidderAccountId: b.BidderAccountId,
			Offer:           b.Offer,
			State:           b.State,
		})
	}

	return &pb.GetAllBidsResponse{Data: data}, nil
}

func (s *BidServiceServer) GetBidById(ctx context.Context, req *pb.GetBidByIdRequest) (*pb.GetBidByIdResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Bids")
	}

	byBidId := req.BidId

	res, err := s.Repo.GetBidById(byBidId)
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetBidById: %+v", err)
	}

	data := pb.Bid{
		Id:              res.Id,
		InvoiceId:       res.InvoiceId,
		BidderAccountId: res.BidderAccountId,
		Offer:           res.Offer,
		State:           res.State,
	}

	return &pb.GetBidByIdResponse{Data: &data}, nil
}

func (s *BidServiceServer) GetBidsByInvoiceId(ctx context.Context, req *pb.GetBidsByInvoiceIdRequest) (*pb.GetBidsByInvoiceIdResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Bids")
	}

	byInvoiceId := req.InvoiceId

	res, err := s.Repo.GetBidsByInvoiceId(byInvoiceId)
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetBidsByInvoiceId: %+v", err)
	}

	data := []*pb.Bid{}

	for _, b := range *res {
		data = append(data, &pb.Bid{
			Id:              b.Id,
			InvoiceId:       b.InvoiceId,
			BidderAccountId: b.BidderAccountId,
			Offer:           b.Offer,
			State:           b.State,
		})
	}

	return &pb.GetBidsByInvoiceIdResponse{Data: data}, nil
}

func (s *BidServiceServer) GetBidsByInvestorId(ctx context.Context, req *pb.GetBidsByInvestorIdRequest) (*pb.GetBidsByInvestorIdResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Bids")
	}
	byInvestorId := req.InvestorId
	res, err := s.Repo.GetBidsByInvestorId(byInvestorId)
	if err != nil {
		return nil, fmt.Errorf("error in response from repository for GetBidsByInvestorId: %+v", err)
	}

	data := []*pb.Bid{}

	for _, b := range *res {
		data = append(data, &pb.Bid{
			Id:              b.Id,
			InvoiceId:       b.InvoiceId,
			BidderAccountId: b.BidderAccountId,
			Offer:           b.Offer,
			State:           b.State,
		})
	}

	return &pb.GetBidsByInvestorIdResponse{Data: data}, nil
}

func (s *BidServiceServer) NewBid(ctx context.Context, req *pb.NewBidRequest) (*pb.NewBidResponse, error) {
	if s.Repo == nil {
		return nil, fmt.Errorf("no repository found for Bids")
	}

	newBidData := &db.Bid{
		InvoiceId:       req.InvoiceId,
		BidderAccountId: req.BidderAccountId,
		Offer:           req.Offer,
	}

	resNewBid, errNewBid := s.Repo.NewBid(newBidData)
	if errNewBid != nil {
		return nil, fmt.Errorf("error in response from repository for NewBid: %+v", errNewBid)
	}

	resFullBids, errFullBids := s.Repo.GetFulfillingBids(resNewBid.InvoiceId)
	if errFullBids != nil {
		return nil, fmt.Errorf("error in response from repository for GetWinningBids: %+v", errFullBids)
	}

	// If GetFulfillingBids is pre-sorted by the DB, we can just take the first one
	if len(*resFullBids) > 0 {
		winnigBidId := (*resFullBids)[0].Id
		// in case we have a winner AdjudicateBid and reset all others
		_, errAdjBid := s.Repo.AdjudicateBid(winnigBidId)
		if errAdjBid != nil {
			return nil, fmt.Errorf("error in response from repository for AdjudicateBid: %+v", errAdjBid)
		}
		_, errBidsToLost := s.Repo.AllRunningBidsToLost(resNewBid.InvoiceId)
		if errBidsToLost != nil {
			return nil, fmt.Errorf("error in response from repository for AllRunningBidsToLost: %+v", errBidsToLost)
		}
		// we need to query again for the latest status update on the Bid
		resUpdBid, errUpdBid := s.Repo.GetBidById(newBidData.Id)
		if errUpdBid != nil {
			return nil, fmt.Errorf("error in response from repository for GetBidById: %+v", errUpdBid)
		}
		resNewBid = resUpdBid
	}

	data := pb.Bid{
		Id:              resNewBid.Id,
		InvoiceId:       resNewBid.InvoiceId,
		BidderAccountId: resNewBid.BidderAccountId,
		Offer:           resNewBid.Offer,
		State:           resNewBid.State,
	}

	return &pb.NewBidResponse{Data: &data}, nil
}
