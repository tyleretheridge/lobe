package writers

import (
	"camm_extractor/internal/containers"
	"fmt"
)

type TerminalWriter struct {
}

func (w *TerminalWriter) Init() {
}

func (w *TerminalWriter) Write(packetContainer containers.PacketContainer) error {
	for _, packet := range packetContainer.Packets() {
		fmt.Printf("%d | %s\n", packet.PacketType(), packet)
	}
	return nil
}
