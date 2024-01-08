package containers

import "camm_extractor/internal/packets"

// Base is the standard packet container that is used by default
type Base struct {
	enabledPackets map[int]struct{}
	storage        []packets.DecodedPacket
	packetCounts   map[uint16]int
}

// PacketCounts returns a map where the key is the integer packet type and the value is the number of
// packets of the corresponding key within the container
func (s *Base) PacketCounts() map[uint16]int {
	return s.packetCounts
}

// Packets returns all the decoded data in the container
func (s *Base) Packets() []packets.DecodedPacket {
	return s.storage
}

// AddPacket is a method for external functions to add packets to the internal Base container storage
// Checks if the given packet type is enabled in the container and add it to storage if it is
func (s *Base) AddPacket(packet packets.DecodedPacket, packetType uint16) error {
	_, exists := s.enabledPackets[packet.PacketType()]
	if !exists {
		return nil
	}
	s.storage = append(s.storage, packet)
	s.packetCounts[packetType]++
	return nil
}

// EnabledPackets returns a set of integers that represents the packets that will be stored
func (s *Base) EnabledPackets() map[int]struct{} {
	return s.enabledPackets
}

// NewBaseContainer constructs a new instance of Base container using the supplied options
func NewBaseContainer(options ContainerOptions) (*Base, error) {
	container := Base{
		enabledPackets: options.enabledPackets,
		storage:        []packets.DecodedPacket{},
		packetCounts:   map[uint16]int{},
	}
	return &container, nil
}
