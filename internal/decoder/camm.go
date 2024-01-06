package decoder

import (
	"bytes"
	"camm_extractor/internal/packets"
	"encoding/binary"
	"fmt"
)

type CAMMStream struct {
	DecodedPackets []packets.DecodedPacket
	CountsMap      map[uint16]int
}

func NewCAMMStream() *CAMMStream {
	return &CAMMStream{
		DecodedPackets: []packets.DecodedPacket{},
		CountsMap:      map[uint16]int{},
	}
}

func (s *CAMMStream) Decode(buf *bytes.Buffer) error {
	for {
		// exit when buffer is empty
		if buf.Len() == 0 {
			break
		}

		// Always Reserved Packet that is empty
		var reserved uint16
		err := binary.Read(buf, binary.LittleEndian, &reserved)
		if err != nil {
			return fmt.Errorf("error reading reserved bytes: %v", err)
		}
		// Read packetType for switch statement
		var packetType uint16
		err = binary.Read(buf, binary.LittleEndian, &packetType)
		if err != nil {
			return fmt.Errorf("error reading reserved bytes: %v", err)
		}

		packet, err := packets.GetEmptyPacket(packetType)
		if err != nil {
			return err
		}
		err = binary.Read(buf, binary.LittleEndian, packet)
		if err != nil {
			return err
		}
		s.addPacket(packet, packetType)
	}
	return nil
}

func (s *CAMMStream) addPacket(packet packets.DecodedPacket, packetType uint16) {
	s.DecodedPackets = append(s.DecodedPackets, packet)
	s.CountsMap[packetType]++
}
