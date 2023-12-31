package writers

import (
	"camm_extractor/internal/containers"
	"camm_extractor/internal/packets"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CSVWriter struct {
	Filename string
}

func NewCSVWriter() *CSVWriter {
	return &CSVWriter{}
}

type CSVData interface {
	CSVHeader() []string
	CSVFormat() []string
}

func (w CSVWriter) Write(packetContainer containers.PacketContainer) error {
	packetTypeCount := len(packetContainer.EnabledPackets())
	if packetTypeCount == 1 {
		return w.toSingleFile(packetContainer)
	}
	return w.toSingleFileWrappedRecords(packetContainer)
}

func (w CSVWriter) toSingleFile(container containers.PacketContainer) error {
	file, err := os.OpenFile(w.Filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("error opening output file: %w", err)
	}
	writer := csv.NewWriter(file)
	// Get Data from Container
	data := container.Packets()
	// Write Headers
	firstPacket, err := adaptPacketToCSVWritable(data[0])
	if err != nil {
		return err
	}
	err = writer.Write(firstPacket.header())
	if err != nil {
		return fmt.Errorf("error writing header %s to csv: %w", firstPacket.header(), err)
	}
	// Write Rows
	for i := range data {
		writePacket, err := adaptPacketToCSVWritable(data[i])
		if err != nil {
			return err
		}
		err = writer.Write(writePacket.values())
		if err != nil {
			return fmt.Errorf("error writing record %s to csv: %w", writePacket.values(), err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		_ = file.Close()
		return fmt.Errorf("error writing CSV: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}
	return nil
}

func (w CSVWriter) toSingleFileWrappedRecords(container containers.PacketContainer) error {
	file, err := os.OpenFile(w.Filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("error opening output file: %w", err)
	}
	writer := csv.NewWriter(file)
	// Get Data from Container
	data := container.Packets()
	// Write Headers
	headers := []string{"packet_type", "packet_headers", "packet_values"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}
	// Write Rows
	for i := range data {
		writePacket, err := adaptPacketToCSVWritable(data[i])
		if err != nil {
			return err
		}
		wrappedData := wrapPacketData(writePacket)
		err = writer.Write(wrappedData)
		if err != nil {
			return fmt.Errorf("error writing record %s to csv: %w", wrappedData, err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		_ = file.Close()
		return fmt.Errorf("error writing CSV: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}
	return nil
}

type csvWritable interface {
	header() []string
	values() []string
	packetType() int
}

func wrapPacketData(packet csvWritable) []string {
	packetType := strconv.Itoa(packet.packetType())
	headers := fmt.Sprintf("(%s)", strings.Join(packet.header(), ", "))
	values := fmt.Sprintf("(%s)", strings.Join(packet.values(), ", "))
	return []string{packetType, headers, values}
}

func adaptPacketToCSVWritable(packet packets.DecodedPacket) (csvWritable, error) {
	switch p := packet.(type) {
	case packets.TypeZero:
		return csvTypeZeroAdapter(p), nil
	case packets.TypeOne:
		return csvTypeOneAdapter(p), nil
	case packets.TypeTwo:
		return csvTypeTwoAdapter(p), nil
	case packets.TypeThree:
		return csvTypeThreeAdapter(p), nil
	case packets.TypeFour:
		return csvTypeFourAdapter(p), nil
	case packets.TypeFive:
		return csvTypeFiveAdapter(p), nil
	case packets.TypeSix:
		return csvTypeSixAdapter(p), nil
	case packets.TypeSeven:
		return csvTypeSevenAdapter(p), nil
	}
	return nil, errors.New(fmt.Sprintf("csv: unable to adapt packet %s to csv writable", packet))
}

type csvTypeZero struct {
	data []string
}

func (p csvTypeZero) header() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p csvTypeZero) values() []string {
	return p.data
}

func (p csvTypeZero) packetType() int {
	return 0
}

func csvTypeZeroAdapter(p packets.TypeZero) csvWritable {
	return csvTypeZero{[]string{
		fmt.Sprintf("%f", p.AngleAxis[0]),
		fmt.Sprintf("%f", p.AngleAxis[1]),
		fmt.Sprintf("%f", p.AngleAxis[2]),
	}}
}

type csvTypeOne struct {
	data []string
}

func (p csvTypeOne) header() []string {
	return []string{"Pixel Exposure Time (ns)", "Rolling Shutter Skew Time (ns)"}
}

func (p csvTypeOne) values() []string {
	return p.data
}

func (p csvTypeOne) packetType() int {
	return 1
}

func csvTypeOneAdapter(p packets.TypeOne) csvWritable {
	return csvTypeOne{[]string{
		fmt.Sprintf("%d", p.PixelExposureTime),
		fmt.Sprintf("%d", p.RollingShutterSkewTime),
	}}
}

type csvTypeTwo struct {
	data []string
}

func (p csvTypeTwo) header() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p csvTypeTwo) values() []string {
	return p.data
}
func (p csvTypeTwo) packetType() int {
	return 2
}

func csvTypeTwoAdapter(p packets.TypeTwo) csvWritable {
	return csvTypeTwo{[]string{
		fmt.Sprintf("%f", p.Gyro[0]),
		fmt.Sprintf("%f", p.Gyro[1]),
		fmt.Sprintf("%f", p.Gyro[2]),
	}}
}

type csvTypeThree struct {
	data []string
}

func (p csvTypeThree) header() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p csvTypeThree) values() []string {
	return p.data
}
func (p csvTypeThree) packetType() int {
	return 3
}

func csvTypeThreeAdapter(p packets.TypeThree) csvWritable {
	return csvTypeThree{[]string{
		fmt.Sprintf("%f", p.Acceleration[0]),
		fmt.Sprintf("%f", p.Acceleration[1]),
		fmt.Sprintf("%f", p.Acceleration[2]),
	}}
}

type csvTypeFour struct {
	data []string
}

func (p csvTypeFour) header() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p csvTypeFour) values() []string {
	return p.data
}

func (p csvTypeFour) packetType() int {
	return 4
}
func csvTypeFourAdapter(p packets.TypeFour) csvWritable {
	return csvTypeFour{[]string{
		fmt.Sprintf("%f", p.Position[0]),
		fmt.Sprintf("%f", p.Position[1]),
		fmt.Sprintf("%f", p.Position[2]),
	}}
}

type csvTypeFive struct {
	data []string
}

func (p csvTypeFive) header() []string {
	return []string{"Latitude", "Longitude", "Altitude"}
}

func (p csvTypeFive) values() []string {
	return p.data
}

func (p csvTypeFive) packetType() int {
	return 5
}
func csvTypeFiveAdapter(p packets.TypeFive) csvWritable {
	return csvTypeFive{[]string{
		fmt.Sprintf("%f", p.Latitude),
		fmt.Sprintf("%f", p.Longitude),
		fmt.Sprintf("%f", p.Altitude),
	}}
}

type csvTypeSix struct {
	data []string
}

func (p csvTypeSix) header() []string {
	return []string{
		"TimeGPSEpoch",
		"GPSFixType",
		"Latitude",
		"Longitude",
		"Altitude",
		"HorizontalAccuracy",
		"VerticalAccuracy",
		"VelocityEast",
		"VelocityNorth",
		"VelocityUp",
		"SpeedAccuracy",
	}
}

func (p csvTypeSix) values() []string {
	return p.data
}

func (p csvTypeSix) packetType() int {
	return 6
}

func csvTypeSixAdapter(p packets.TypeSix) csvWritable {
	return csvTypeSix{[]string{
		fmt.Sprintf("%f", p.TimeGPSEpoch),
		fmt.Sprintf("%d", p.GPSFixType),
		fmt.Sprintf("%f", p.Latitude),
		fmt.Sprintf("%f", p.Longitude),
		fmt.Sprintf("%f", p.Altitude),
		fmt.Sprintf("%f", p.HorizontalAccuracy),
		fmt.Sprintf("%f", p.VerticalAccuracy),
		fmt.Sprintf("%f", p.VelocityEast),
		fmt.Sprintf("%f", p.VelocityNorth),
		fmt.Sprintf("%f", p.VelocityUp),
		fmt.Sprintf("%f", p.SpeedAccuracy),
	}}
}

type csvTypeSeven struct {
	data []string
}

func (p csvTypeSeven) header() []string {
	return []string{"B_Field_X", "B_Field_Y", "B_Field_Z"}
}

func (p csvTypeSeven) values() []string {
	return p.data
}

func (p csvTypeSeven) packetType() int {
	return 7
}
func csvTypeSevenAdapter(p packets.TypeSeven) csvWritable {
	return csvTypeSeven{[]string{
		fmt.Sprintf("%f", p.MagneticField[0]),
		fmt.Sprintf("%f", p.MagneticField[1]),
		fmt.Sprintf("%f", p.MagneticField[2]),
	}}
}
