package e2etest

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

func TestE2E_GetAllBids(t *testing.T) {
	tests := []struct {
		desc    string
		client  pb.BidServiceClient
		want    *pb.GetAllBidsResponse
		wantErr bool
	}{
		{
			desc:    "gets the list of all Bids",
			client:  clientServices.bid,
			want:    data.ResponseGetAllBids,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllBids(ctx, data.RequestGetAllBids)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllBids() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got GetAllBids() = \n%+v,\nbut wanted \n%+v", got, tt.want)
			}

		})
	}
}
func TestE2E_NewBid(t *testing.T) {
	tests := []struct {
		desc    string
		client  pb.BidServiceClient
		want    *pb.NewBidResponse
		wantErr bool
	}{
		{
			desc:    "bids on an invoice",
			client:  clientServices.bid,
			want:    data.ResponseNewBid,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.NewBid(ctx, data.RequesteNewBid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got Bid() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}

			// trick to not go crazy with indexes and ids
			if got != nil {
				tt.want.Data.Id = got.Data.Id
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got Bid() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}
