package containers

import "camm_extractor/internal/packets"

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

func (s *CAMMStream) Packets() []packets.DecodedPacket {
	return s.DecodedPackets
}

func (s *CAMMStream) AddPacket(packet packets.DecodedPacket, packetType uint16) error {
	s.DecodedPackets = append(s.DecodedPackets, packet)
	s.CountsMap[packetType]++
	return nil
}
