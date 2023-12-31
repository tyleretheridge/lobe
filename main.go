package main

import (
	"bytes"
	"camm_extractor/internal/decoder"
	"camm_extractor/internal/output"
	"fmt"
	"log"
	"os"
	"time"
)

const FFmpegPath = "C:\\Tools\\ffmpeg\\bin\\ffmpeg.exe"
const inputVideo = "C:\\Users\\tjaco\\Downloads\\video_lrv_1.mp4"
const outputPath = "C:\\Users\\tjaco\\Downloads\\camm.bin"

const targetPacket = 6
const csvFn = `C:\Users\tjaco\GolandProjects\camm_extractor\output_data.csv`

func main() {
	startTime := time.Now()
	cammStream := decoder.NewCAMMStream()
	data, err := os.ReadFile(outputPath)
	if err != nil {
		log.Fatalf("Unable to read file: %s", outputPath)
	}
	buf := bytes.NewBuffer(data)
	err = cammStream.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}

	var packets []output.CSVData
	for _, packet := range cammStream.DecodedPackets {
		if packet.PacketType() != targetPacket {
			continue
		}
		packets = append(packets, packet)
	}

	w := output.CSVWriter{}
	err = w.ToFile(csvFn, packets)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(startTime)
	fmt.Printf("Completed in %f seconds.\n", duration.Seconds())
}
