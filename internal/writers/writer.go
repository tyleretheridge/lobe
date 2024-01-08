package writers

import (
	"camm_extractor/internal/containers"
)

type WriterType int

const (
	Standard WriterType = iota
	CSV
)

func GetWriter(writerType WriterType) DataWriter {
	return nil
}

type DataWriter interface {
	Init()
	Write(packetContainer containers.PacketContainer) error
}

type WriterOptions struct {
	filename string
}
