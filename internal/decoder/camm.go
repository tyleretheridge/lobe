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

func NewCAMMStream() *CAMMStream {
	return &CAMMStream{
		DecodedPackets: []DecodedPacket{},
		CountsMap:      map[uint16]int{},
	}
}

func (s *CAMMStream) Decode(buf *bytes.Buffer) error {
	var cycleTimes []float64
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

		packet, err := NewDecodedPacket(packetType)
		if err != nil {
			return err
		}
		err = binary.Read(buf, binary.LittleEndian, packet)
		if err != nil {
			return err
		}
		p, ok := packet.(*PacketTypeSix)
		if ok {
			cycleTimes = append(cycleTimes, p.TimeGPSEpoch)
		}
		s.AddPacket(packet, packetType)
	}
	return nil
}

func (s *CAMMStream) AddPacket(packet DecodedPacket, packetType uint16) {
	s.DecodedPackets = append(s.DecodedPackets, packet)
	s.CountsMap[packetType]++
}
