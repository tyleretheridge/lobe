package containers

import "camm_extractor/internal/packets"

type PacketContainer interface {
	AddPacket(packet packets.DecodedPacket, packetType uint16) error
	Packets() []packets.DecodedPacket
}
