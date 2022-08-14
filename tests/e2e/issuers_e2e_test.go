package e2etest

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

func TestE2E_GetAllIssuers(t *testing.T) {
	testCases := []struct {
		desc    string
		client  pb.IssuerServiceClient
		want    *pb.GetAllIssuersResponse
		wantErr bool
	}{
		{
			desc:    "gets the list of all Issuers",
			client:  clientServices.issuer,
			want:    &pb.GetAllIssuersResponse{Data: *data.SeededAllIssuersList},
			wantErr: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			got, err := tt.client.GetAllIssuers(ctx, &pb.Empty{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Got GetAllIssuers() error = %v, instead expected error %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data, tt.want.Data) {
				t.Errorf("Got GetAllIssuers() = %v, but wanted %v", got, tt.want)
			}

		})
	}
}
