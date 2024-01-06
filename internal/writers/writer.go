package writers

import (
	"camm_extractor/internal/containers"
	"fmt"
)

type WriterType int

const (
	CSV WriterType = iota
)

func GetWriter(writerType WriterType) DataWriter {
	return nil
}

type DataWriter interface {
	Init()
	Write(packetContainer containers.PacketContainer) error
}

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
