package client

import (
	"github.com/regner/albiondata-client/log"
)

type operationRealEstateBidOnAuction struct {
}

func (op operationRealEstateBidOnAuction) Process(state *albionState, uploader iuploader) {
	log.Debug("Got RealEstateBidOnAuction operation...")
}

type operationRealEstateBidOnAuctionResponse struct {
}

func (op operationRealEstateBidOnAuctionResponse) Process(state *albionState, uploader iuploader) {
	log.Debug("Got response to RealEstateBidOnAuction operation...")
}
