package server

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	fakeRepos "github.com/marcosantonastasi/arex_challenge/test/fakes/fakerepos"
)

func TestInvestorServiceServer_GetAllInvestors(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Empty
	}

	tests := []struct {
		name    string
		s       *InvestorServiceServer
		args    args
		want    *pb.GetAllInvestorsResponse
		wantErr bool
	}{
		{
			name:    "gets all Investors on the repository (3 for newly seeded db)",
			s:       &InvestorServiceServer{Repo: &fakeRepos.FakeInvestorsRepository{}},
			args:    args{ctx: context.Background(), in: &pb.Empty{}},
			want:    &pb.GetAllInvestorsResponse{Data: fakeRepos.FakeAllInvestorsList},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllInvestors(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got InvestorServiceServer.GetAllInvestors() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got InvestorServiceServer.GetAllInvestors() = %v, but wanted %v", got, tt.want)
			}
		})
	}
}

func TestIssuerServiceServer_GetAllIssuers(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Empty
	}
	tests := []struct {
		name    string
		s       *IssuerServiceServer
		args    args
		want    *pb.GetAllIssuersResponse
		wantErr bool
	}{
		{
			name:    "gets all Issuers on the repository (3 for newly seeded db)",
			s:       &IssuerServiceServer{Repo: &fakeRepos.FakeIssuersRepository{}},
			args:    args{ctx: context.Background(), in: &pb.Empty{}},
			want:    &pb.GetAllIssuersResponse{Data: fakeRepos.FakeAllIssuersList},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllIssuers(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("IssuerServiceServer.GetAllIssuers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IssuerServiceServer.GetAllIssuers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvoiceServiceServer_GetAllInvoices(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Empty
	}
	tests := []struct {
		name    string
		s       *InvoiceServiceServer
		args    args
		want    *pb.GetAllInvoicesResponse
		wantErr bool
	}{
		{
			name:    "gets all Invoices on the repository (3 for newly seeded db)",
			s:       &InvoiceServiceServer{Repo: &fakeRepos.FakeInvoicesRepository{}},
			args:    args{ctx: context.Background(), in: &pb.Empty{}},
			want:    &pb.GetAllInvoicesResponse{Data: fakeRepos.FakeAllInvoicesList},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllInvoices(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvoiceServiceServer.GetAllinvoices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvoiceServiceServer.GetAllinvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
