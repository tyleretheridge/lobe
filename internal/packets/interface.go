package packets

import (
	"fmt"
)

type DecodedPacket interface {
	PacketType() int
}

// GetEmptyPacket returns a newly instantiated packet based on the packet Type value passed
// Into the function. See CAMM Spec sheet for a reference to this mapping.
func GetEmptyPacket(packetType uint16) (DecodedPacket, error) {
	switch packetType {
	case 0:
		return new(TypeZero), nil
	case 1:
		return new(TypeOne), nil
	case 2:
		return new(TypeTwo), nil
	case 3:
		return new(TypeThree), nil
	case 4:
		return new(TypeFour), nil
	case 5:
		return new(TypeFive), nil
	case 6:
		return new(TypeSix), nil
	case 7:
		return new(TypeSeven), nil
	}
	return nil, fmt.Errorf("unsupported packet type %v", packetType)
}
