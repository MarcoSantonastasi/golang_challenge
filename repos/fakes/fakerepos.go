package fakerepos

import (
	"encoding/json"
	"io"
	"os"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

var FakeAllInvestorsList []*pb.Investor

func init() {
	jsonFile, fileErr := os.Open("../../repos/fakes/fakeInvestors.json")
	if fileErr != nil {
		panic(".json file error")
	}
	jsonData, dataErr := io.ReadAll(jsonFile)
	if dataErr != nil {
		panic(".json file error")
	}
	jsonErr := json.Unmarshal(jsonData, &FakeAllInvestorsList)
	if jsonErr != nil {
		panic(".json file error")
	}
	defer jsonFile.Close()
}

type FakeInvestorsRepository struct {
}

func (repo *FakeInvestorsRepository) GetAllInvestors() ([]*pb.Investor, error) {
	return FakeAllInvestorsList, nil
}

type FakeIssuersRepository struct {
}

type FakeInvoiceRepository struct {
}
