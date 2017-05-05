package marketassembley

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"reflect"
	"fmt"
	"regexp"
	"encoding/binary"
	"strings"
)

/*
19    -> Indicator of if there is more packets
20:24 -> Packet ID
24:28 -> ID of first packet in the bundle
28:32 -> Expected number of packets
32:36 -> Packet ID in bundle
44:54 -> Location of market response start indicator
44:   -> Data when not first packet of market response
55:   -> Data when first packet of market response
 */


var marketStartIndicator = []byte{243, 3, 1, 0, 0, 42, 0, 2, 0, 121}
var morePacketsIndicator = byte(164)

type MarketAssembler struct {
	processing bool
	itemCount int
	itemsBuffer []byte

}

func NewMarketAssembler() *MarketAssembler {
	return &MarketAssembler{}
}

func (ma *MarketAssembler) ProcessPacket(packet gopacket.Packet) {
	udp := packet.Layer(layers.LayerTypeUDP)
	if udp != nil {
		udp := udp.(*layers.UDP)

		if len(udp.Payload) < 56 {
			return
		}

		marketHeader := udp.Payload[44:54]

		if !ma.processing && reflect.DeepEqual(marketHeader, marketStartIndicator) {
			ma.itemCount = int(binary.BigEndian.Uint16(udp.Payload[54:56]))


			ma.processing = true
			ma.itemsBuffer = nil
			ma.itemsBuffer = append(ma.itemsBuffer, udp.Payload[59:]...)

		} else if ma.processing {
			ma.itemsBuffer = append(ma.itemsBuffer, udp.Payload[44:]...)

			if udp.Payload[19] != morePacketsIndicator {
				ma.processing = false
				results := extractStrings(ma.itemsBuffer)
				fmt.Println(results)
				fmt.Println(len(results))
				fmt.Println(ma.itemCount)

				ma.itemsBuffer = nil
			}
		}
	}
}

func extractStrings(payload []byte) []string {
	var results []string
	r, _ := regexp.Compile("\\{[^\\{\\}]*\\}")

	for _, match :=  range r.FindAllStringSubmatch(string(payload), -1) {
		results = append(results, strings.Join(match, ""))
	}

	return results
}