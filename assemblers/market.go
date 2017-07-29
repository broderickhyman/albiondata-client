package assemblers

import (
	"encoding/binary"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/regner/albionmarket-client/utils"
)

/*
19    -> Location of more packets indicator
20:24 -> Packet ID
24:28 -> ID of first packet in the bundle
28:32 -> Expected number of packets
32:36 -> Packet ID in bundle
44:   -> Data when not first packet of market response
44:54 -> Location of market response start indicator
54:56 -> Number of market items to be expected in total
56:   -> Data when first packet of market response
*/

var (
	marketStartIndicator = []byte{243, 3, 1, 0, 0, 42, 0, 2, 0, 121}
	morePacketsIndicator = byte(164)
)

type MarketAssembler struct {
	itemCount      int
	packetCount    int
	processedCount int
	itemsBuffer    []byte
	processing     bool
	config         utils.ClientConfig
	locationId     int
}

func NewMarketAssembler(config utils.ClientConfig) *MarketAssembler {
	return &MarketAssembler{
		config:     config,
		locationId: 0,
	}
}

func (ma *MarketAssembler) ProcessPacket(packet gopacket.Packet) {
	udp := packet.Layer(layers.LayerTypeUDP)
	if udp != nil {
		udp := udp.(*layers.UDP)

		if len(udp.Payload) < 56 {
			return
		}

		morePackets := udp.Payload[19]
		currentPacket := int(binary.BigEndian.Uint32(udp.Payload[32:36]))
		marketHeader := udp.Payload[44:54]

		if reflect.DeepEqual(marketHeader, marketStartIndicator) {
			// Ensure things are reset
			ma.processing = true
			ma.itemsBuffer = nil
			ma.packetCount = 0

			ma.itemCount = int(binary.BigEndian.Uint16(udp.Payload[54:56]))

			// Minus 1 because packet counts themselves start at zero
			ma.packetCount = int(binary.BigEndian.Uint32(udp.Payload[28:32])) - 1

			ma.itemsBuffer = append(ma.itemsBuffer, udp.Payload[59:]...)
		} else if !ma.processing {
			// We are not processing so may as well ignore! :D
			return
		} else if morePackets != morePacketsIndicator && currentPacket != ma.packetCount {
			// If we don't have the more packets indicator AND the current packet
			// doesn't have a packet count matching the number of packets we expect
			// move on, possibly got some other UDP packet we don't care about.
			return
		} else if morePackets != morePacketsIndicator && currentPacket == ma.packetCount {
			// There are no more packets and this is confirmed as the last packet
			ma.itemsBuffer = append(ma.itemsBuffer, udp.Payload[44:]...)

			results := extractStrings(ma.itemsBuffer)

			if !ma.config.DisableUpload {
				utils.SendMarketItems(results, ma.config.IngestUrl, ma.locationId)
			}

			if ma.config.SaveLocally {
				utils.SaveMarketItems(results)
			}

			ma.processing = false
		} else {
			ma.itemsBuffer = append(ma.itemsBuffer, udp.Payload[44:]...)
		}
	}
}

func extractStrings(payload []byte) []string {
	var results []string
	r, _ := regexp.Compile("\\{[^\\{\\}]*\\}")

	for _, match := range r.FindAllStringSubmatch(string(payload), -1) {
		results = append(results, strings.Join(match, ""))
	}

	return results
}
