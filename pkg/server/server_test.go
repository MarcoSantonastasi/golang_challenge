package server

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllInvestors(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvestorServiceServer.GetAllInvestors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvestorServiceServer.GetAllInvestors() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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

func TestInvoiceServiceServer_GetAllinvoices(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllinvoices(tt.args.ctx, tt.args.in)
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
