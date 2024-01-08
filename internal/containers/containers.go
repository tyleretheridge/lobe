package containers

import (
	"camm_extractor/internal/packets"
	"camm_extractor/internal/utils"
	"errors"
)

type PacketContainer interface {
	AddPacket(packet packets.DecodedPacket, packetType uint16) error
	Packets() []packets.DecodedPacket
	PacketCounts() map[uint16]int
	EnabledPackets() map[int]struct{}
}

// Enum for implemented container types to be used in GetContainer
type containerType int

const (
	BaseContainer containerType = iota
)

// ContainerOptions models the configuration for data container
type ContainerOptions struct {
	enabledPackets map[int]struct{}
}

// NewContainerOptions returns a constructed instance of ContainerOptions
func NewContainerOptions(enabledPackets []int) (*ContainerOptions, error) {
	packetSet, err := utils.IntSliceToSet(enabledPackets)
	if err != nil {
		return nil, err
	}
	return &ContainerOptions{packetSet}, nil
}

// GetContainer Returns a container configured with the given options
func GetContainer(container containerType, options ContainerOptions) (PacketContainer, error) {
	switch container {
	case BaseContainer:
		return NewBaseContainer(options)
	default:
		return nil, errors.New("invalid container type")
	}
}
