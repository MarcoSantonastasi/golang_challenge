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
			want:    &pb.GetAllBidsResponse{Data: *data.SeededAllBidsList},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllBids(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllBids() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("Got GetAllBids() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}
