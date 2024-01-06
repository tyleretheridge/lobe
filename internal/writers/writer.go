package writers

import "camm_extractor/internal/decoder"

type WriterType int

const (
	CSV WriterType = iota
)

func GetWriter(writerType WriterType) DataWriter {
	return nil
}

type DataWriter interface {
	Write(stream decoder.CAMMStream) error
}
