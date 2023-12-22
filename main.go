package main

import (
	"bytes"
	"camm_extractor/internal/decoder"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const FFmpegPath = "C:\\Tools\\ffmpeg\\bin\\ffmpeg.exe"
const inputVideo = "C:\\Users\\tjaco\\Downloads\\video_lrv_1.mp4"
const outputPath = "C:\\Users\\tjaco\\Downloads\\camm.bin"

func NewExtractionCommand(ffmpegPath string, inputPath string, outputPath string) *exec.Cmd {
	return exec.Command(ffmpegPath, "-y", "-i", inputPath, "-map_metadata", "-1", "-an", "-vn", "-sn", "-map", "0:d", "-f", "rawvideo", outputPath)
}

func main() {

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
	cammStream.Results()
	for _, packet := range cammStream.DecodedPackets {
		fmt.Printf("%+v\n", packet)
	}
}
