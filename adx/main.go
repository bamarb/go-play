package main

import (
	"github.com/bsm/openrtb/v3"
)

func main() {
}

func mkBidResponse(req *openrtb.BidRequest) *openrtb.BidResponse {
	resp := &openrtb.BidResponse{
		ID: req.ID,
		SeatBids: []openrtb.SeatBid{
			{Bids: []openrtb.Bid{}},
		},
	}

	return resp
}

func mkNativeResp() {
}
