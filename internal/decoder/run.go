package decoder

import (
	"bytes"
	"fmt"
	"os"
)

func Run(binaryFile string) (*CAMMStream, error) {
	cammStream := NewCAMMStream()
	binaryData, err := os.ReadFile(binaryFile)
	if err != nil {
		return nil, fmt.Errorf("error opening binary file: %w", err)
	}

	buf := bytes.NewBuffer(binaryData)
	err = cammStream.Decode(buf)
	if err != nil {
		return nil, fmt.Errorf("error decoding binary stream: %w", err)
	}
	return cammStream, nil
}

//
//func main2() {
//	startTime := time.Now()
//	cammStream := NewCAMMStream()
//	data, err := os.ReadFile(outputPath)
//	if err != nil {
//		log.Fatalf("Unable to read file: %s", outputPath)
//	}
//	buf := bytes.NewBuffer(data)
//	err = cammStream.Decode(buf)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var packets []output.CSVData
//	for _, packet := range cammStream.DecodedPackets {
//		if packet.PacketType() != targetPacket {
//			continue
//		}
//		packets = append(packets, packet)
//	}
//
//	w := output.CSVWriter{}
//	err = w.ToFile(csvFn, packets)
//	if err != nil {
//		log.Fatal(err)
//	}
//	duration := time.Since(startTime)
//	fmt.Printf("Completed in %f seconds.\n", duration.Seconds())
//}
