package fakerepos

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

func TestFakeInvestorsRepository_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name    string
		repo    *FakeInvestorsRepository
		want    []*pb.Investor
		wantErr bool
	}{
		{
			name: "GetAlInvestors () returns exactly the data json file",
			repo: &FakeInvestorsRepository{},
			want: FakeAllInvestorsList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllInvestors()
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeInvestorsRepository.GetAllInvestors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeInvestorsRepository.GetAllInvestors() = %v, want %v", got, tt.want)
			}
		})
	}
}
