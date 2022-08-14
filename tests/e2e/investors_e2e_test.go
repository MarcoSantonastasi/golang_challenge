package e2etest

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

func TestE2E_GetAllInvestors(t *testing.T) {
	tests := []struct {
		desc    string
		client  pb.InvestorServiceClient
		want    *pb.GetAllInvestorsResponse
		wantErr bool
	}{
		{
			desc:    "gets the list of all Investors",
			client:  clientServices.investor,
			want:    &pb.GetAllInvestorsResponse{Data: *data.SeededAllInvestorsList},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllInvestors(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllInvestors() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("Got GetAllInvestors() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}
