package writers

import (
	"camm_extractor/internal/packets"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type CSVWriter struct {
}

type CSVData interface {
	CSVHeader() []string
	CSVFormat() []string
}

func (w CSVWriter) ToFile(filepath string, data []CSVData) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("error opening output file: %w", err)
	}
	writer := csv.NewWriter(file)

	// Write Headers
	err = writer.Write(data[0].CSVHeader())
	if err != nil {
		return fmt.Errorf("error writing header %s to csv: %w", data[0].CSVHeader(), err)
	}

	// Write Rows
	for _, entry := range data {
		err = writer.Write(entry.CSVFormat())
		if err != nil {
			return fmt.Errorf("error writing record %s to csv: %w", entry.CSVFormat(), err)
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

func csvTypeSevenAdapter(p packets.TypeSeven) csvWritable {
	return csvTypeSeven{[]string{
		fmt.Sprintf("%f", p.MagneticField[0]),
		fmt.Sprintf("%f", p.MagneticField[1]),
		fmt.Sprintf("%f", p.MagneticField[2]),
	}}
}
