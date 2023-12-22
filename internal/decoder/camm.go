package decoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CAMMStream struct {
	DecodedPackets []DecodedPacket
	CountsMap      map[uint16]int
}

func (s *CAMMStream) Decode(buf *bytes.Buffer) error {
	for {
		// exit when buffer is empty
		if buf.Len() == 0 {
			break
		}

		// Always Reserved Packet that's empty
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

		packet, err := NewDecodedPacket(packetType)
		if err != nil {
			return err
		}
		err = binary.Read(buf, binary.LittleEndian, packet)
		if err != nil {
			return err
		}
		s.AddPacket(packet, packetType)
	}
	return nil
}

func (s *CAMMStream) Results() {
	for _, packet := range s.DecodedPackets {
		fmt.Printf("First Packet Type: %v\n", packet.PacketType())
		break
	}
	for packetType, count := range s.CountsMap {
		fmt.Printf("Packet Type: %d, Count: %d\n", packetType, count)
	}
}

func (s *CAMMStream) AddPacket(packet DecodedPacket, packetType uint16) {
	s.DecodedPackets = append(s.DecodedPackets, packet)
	s.CountsMap[packetType]++
}

func NewCAMMStream() *CAMMStream {
	return &CAMMStream{
		DecodedPackets: []DecodedPacket{},
		CountsMap:      map[uint16]int{},
	}
}
