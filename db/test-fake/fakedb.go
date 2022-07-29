package fakedb

import (
	"encoding/json"
	"io"
	"os"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type FakeArexDb struct {
}

var FakeAllInvestorsList []*pb.Investor

func init() {
	jsonFile, fileErr := os.Open("../../db/test-fake/fakeInvestors.json")
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

func (db *FakeArexDb) GetAllInvestors() []*pb.Investor {
	return FakeAllInvestorsList
}
