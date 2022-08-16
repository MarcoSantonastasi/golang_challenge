package server

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
	stubRepos "github.com/marcosantonastasi/arex_challenge/internal/fixtures/stubs/stubrepos"
)

func TestInvestorServiceServer_GetAllInvestors(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.Empty
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
			s:       &InvestorServiceServer{Repo: &stubRepos.StubInvestorsRepository{}},
			args:    args{ctx: context.Background(), req: data.RequestGetAllInvestors},
			want:    data.ResponseGetAllInvestors,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllInvestors(tt.args.ctx, tt.args.req)
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
		req *pb.Empty
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
			s:       &IssuerServiceServer{Repo: &stubRepos.StubIssuersRepository{}},
			args:    args{ctx: context.Background(), req: data.RequestGetAllIssuers},
			want:    data.ResponseGetAllIssuers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllIssuers(tt.args.ctx, tt.args.req)
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
		req *pb.Empty
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
			s:       &InvoiceServiceServer{Repo: &stubRepos.StubInvoicesRepository{}},
			args:    args{ctx: context.Background(), req: data.RequestGetAllInvoices},
			want:    data.ResponseGetAllInvoices,
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllInvoices(tt.args.ctx, tt.args.req)
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
