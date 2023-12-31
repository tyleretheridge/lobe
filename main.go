package main

import (
	"bytes"
	"camm_extractor/internal/decoder"
	"log"
	"os"
)

const FFmpegPath = "C:\\Tools\\ffmpeg\\bin\\ffmpeg.exe"
const inputVideo = "C:\\Users\\tjaco\\Downloads\\video_lrv_1.mp4"
const outputPath = "C:\\Users\\tjaco\\Downloads\\camm.bin"

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

}
