package decoder

import (
	"bytes"
	"camm_extractor/internal/containers"
	"camm_extractor/internal/packets"
	"encoding/binary"
	"fmt"
)

// DecodeCAMMData is the algorithm that handles the conversion of binary camm data in a buffer to packets and writes
// them to a PacketContainer
func DecodeCAMMData(buf *bytes.Buffer, container containers.PacketContainer) error {
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
		err = container.AddPacket(packet, packetType)
		if err != nil {
			return fmt.Errorf("decoder: error writing packet data to container: %w", err)
		}
	}
	return nil
}
