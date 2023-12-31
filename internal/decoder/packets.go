package decoder

import (
	"camm_extractor/internal/output"
	"fmt"
)

type ReservedPacket struct {
	Reserved uint16
}

type TypeHeader struct {
	Value uint16
}

type PacketTypeZero struct {
	AngleAxis [3]float32
}

func (p PacketTypeZero) PacketType() int {
	return 0
}

func (p PacketTypeZero) String() string {
	return fmt.Sprintf("%f, %f, %f", p.AngleAxis[0], p.AngleAxis[1], p.AngleAxis[2])
}

func (p PacketTypeZero) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p PacketTypeZero) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.AngleAxis[0]),
		fmt.Sprintf("%f", p.AngleAxis[1]),
		fmt.Sprintf("%f", p.AngleAxis[2]),
	}
}

type PacketTypeOne struct {
	PixelExposureTime      int32
	RollingShutterSkewTime int32
}

func (p PacketTypeOne) PacketType() int {
	return 1
}

func (p PacketTypeOne) String() string {
	return fmt.Sprintf("%d, %d", p.PixelExposureTime, p.RollingShutterSkewTime)
}

func (p PacketTypeOne) CSVHeader() []string {
	return []string{"Pixel Exposure Time (ns)", "Rolling Shutter Skew Time (ns)"}
}

func (p PacketTypeOne) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%d", p.PixelExposureTime),
		fmt.Sprintf("%d", p.RollingShutterSkewTime),
	}
}

type PacketTypeTwo struct {
	Gyro [3]float32
}

func (p PacketTypeTwo) PacketType() int {
	return 2
}

func (p PacketTypeTwo) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Gyro[0], p.Gyro[1], p.Gyro[2])
}

func (p PacketTypeTwo) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p PacketTypeTwo) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Gyro[0]),
		fmt.Sprintf("%f", p.Gyro[1]),
		fmt.Sprintf("%f", p.Gyro[2]),
	}
}

type PacketTypeThree struct {
	Acceleration [3]float32
}

func (p PacketTypeThree) PacketType() int {
	return 3
}

func (p PacketTypeThree) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Acceleration[0], p.Acceleration[1], p.Acceleration[2])
}

func (p PacketTypeThree) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p PacketTypeThree) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Acceleration[0]),
		fmt.Sprintf("%f", p.Acceleration[1]),
		fmt.Sprintf("%f", p.Acceleration[2]),
	}
}

type PacketTypeFour struct {
	Position [3]float32
}

func (p PacketTypeFour) PacketType() int {
	return 4
}

func (p PacketTypeFour) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Position[0], p.Position[1], p.Position[2])
}

func (p PacketTypeFour) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p PacketTypeFour) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Position[0]),
		fmt.Sprintf("%f", p.Position[1]),
		fmt.Sprintf("%f", p.Position[2]),
	}
}

type PacketTypeFive struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (p PacketTypeFive) PacketType() int {
	return 5
}

func (p PacketTypeFive) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Latitude, p.Longitude, p.Altitude)
}

func (p PacketTypeFive) CSVHeader() []string {
	return []string{"Latitude", "Longitude", "Altitude"}
}

func (p PacketTypeFive) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Latitude),
		fmt.Sprintf("%f", p.Longitude),
		fmt.Sprintf("%f", p.Altitude),
	}
}

type PacketTypeSix struct {
	TimeGPSEpoch       float64 //seconds
	GPSFixType         int32   // 0, 2(D), 3(D)
	Latitude           float64 // degrees
	Longitude          float64 // degrees
	Altitude           float32 // meters
	HorizontalAccuracy float32 // meters
	VerticalAccuracy   float32 // meters
	VelocityEast       float32 // meters/second
	VelocityNorth      float32 // meters/second
	VelocityUp         float32 // meters/second
	SpeedAccuracy      float32 // meters/second
}

func (p PacketTypeSix) PacketType() int {
	return 6
}

func (p PacketTypeSix) String() string {
	return fmt.Sprintf("%f, %d, %f, %f, %f, %f, %f, %f, %f, %f, %f",
		p.TimeGPSEpoch,
		p.GPSFixType,
		p.Latitude,
		p.Longitude,
		p.Altitude,
		p.HorizontalAccuracy,
		p.VerticalAccuracy,
		p.VelocityEast,
		p.VelocityNorth,
		p.VelocityUp,
		p.SpeedAccuracy)
}

func (p PacketTypeSix) CSVHeader() []string {
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

func (p PacketTypeSix) CSVFormat() []string {
	return []string{
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
	}
}

type PacketTypeSeven struct {
	MagneticField [3]float32
}

func (p PacketTypeSeven) PacketType() int {
	return 7
}

func (p PacketTypeSeven) String() string {
	return fmt.Sprintf("%f, %f, %f", p.MagneticField[0], p.MagneticField[1], p.MagneticField[2])
}

func (p PacketTypeSeven) CSVHeader() []string {
	return []string{"B_Field_X", "B_Field_Y", "B_Field_Z"}
}

func (p PacketTypeSeven) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.MagneticField[0]),
		fmt.Sprintf("%f", p.MagneticField[1]),
		fmt.Sprintf("%f", p.MagneticField[2]),
	}
}

type DecodedPacket interface {
	PacketType() int
	fmt.Stringer
	output.CSVData
}

func NewDecodedPacket(packetType uint16) (DecodedPacket, error) {
	switch packetType {
	case 0:
		return new(PacketTypeZero), nil
	case 1:
		return new(PacketTypeOne), nil
	case 2:
		return new(PacketTypeTwo), nil
	case 3:
		return new(PacketTypeThree), nil
	case 4:
		return new(PacketTypeFour), nil
	case 5:
		return new(PacketTypeFive), nil
	case 6:
		return new(PacketTypeSix), nil
	case 7:
		return new(PacketTypeSeven), nil
	}
	return nil, fmt.Errorf("unsupported packet type %v", packetType)
}
