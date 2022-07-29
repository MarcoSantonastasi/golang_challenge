package client

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type Client struct {
}

func (client *Client) GetAllInvestors() ([]*pb.Investor, error)
