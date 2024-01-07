package containers

import "camm_extractor/internal/packets"

type PacketContainer interface {
	AddPacket(packet packets.DecodedPacket, packetType uint16) error
	Packets() []packets.DecodedPacket
	PacketCounts() map[uint16]int
	EnabledPackets() map[uint8]struct{}
}

type Base struct {
	enabledPackets map[uint8]struct{}
	storage        []packets.DecodedPacket
	packetCounts   map[uint16]int
}

func NewBaseContainer() *Base {
	return &Base{
		storage:      []packets.DecodedPacket{},
		packetCounts: map[uint16]int{},
	}
}

func (s *Base) PacketCounts() map[uint16]int {
	return s.packetCounts
}

func (s *Base) Packets() []packets.DecodedPacket {
	return s.storage
}

func (s *Base) AddPacket(packet packets.DecodedPacket, packetType uint16) error {
	s.storage = append(s.storage, packet)
	s.packetCounts[packetType]++
	return nil
}

func (s *Base) EnabledPackets() map[uint8]struct{} {
	return s.enabledPackets
}
