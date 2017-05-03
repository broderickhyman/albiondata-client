package marketassembley

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"reflect"
	"fmt"
	"encoding/binary"
)

var marketStartIndicator = []byte{243, 3, 1, 0, 0, 42, 0, 2, 0, 121}

type MarketAssembler struct {
	processing bool
	marketJson string
	itemCount int
}

func NewMarketAssembler() *MarketAssembler {
	return &MarketAssembler{}
}

func (ma MarketAssembler) ProcessPacket(packet gopacket.Packet) {
	// We only care about UDP packets, hopefully that is all we are passed but lets verify
	udp := packet.Layer(layers.LayerTypeUDP)
	if udp != nil {
		udp := udp.(*layers.UDP)

		if udp.Length < 55 {
			return
		}

		// The 10 bytes starting at 44 are where we expect to find the marketStartIndicator.
		marketHeader := udp.Payload[44:54]

		if reflect.DeepEqual(marketHeader, marketStartIndicator) {
			// TODO: Create assembler thing
			// pass every packet to assembler
			// assembler puts pieces together
			// assembler detects new market response
			// starts parsing market data
			// when complete assembler sends market data off and resets assembler

			// Prototype stuff here
			// https://doc.photonengine.com/zh-tw/realtime/current/reference/serialization-in-photon
			// These two bytes detail the number of market entries in the response
			numItems := udp.Payload[54:56]
			fmt.Println(numItems)

			// Skipping byte 56 actually, not sure what it does

			// These two bytes seem to be a short and detail how many bytes are in the next entry
			// This is repeated before each entry
			// So should be able to take it, parse the next N bytes into a string, then repeat until the
			// end of the packet. Continue to next packet and repeat until the we have recorded the number
			// of market entries specified in the first packet.
			firstEntryLength := udp.Payload[57:59]
			fmt.Println(binary.BigEndian.Uint16(firstEntryLength))
			fmt.Println(string(udp.Payload[59:59+binary.BigEndian.Uint16(firstEntryLength)]))
		}
	}
}
