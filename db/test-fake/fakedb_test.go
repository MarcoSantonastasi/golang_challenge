package fakedb

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

func TestFakeDb(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestFakeArexDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name string
		db   *FakeArexDb
		want []*pb.Investor
	}{
		{
			name: "GetAlInvestors () returns exactly the data json file",
			db:   &FakeArexDb{},
			want: FakeAllInvestorsList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllInvestors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got FakeArexDb.GetAllInvestors() = %v, but wanted %v", got, tt.want)
			}
		})
	}
}
