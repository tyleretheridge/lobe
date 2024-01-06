package writers

import (
	"fmt"
)

type Reserved struct {
	Reserved uint16
}

type TypeHeader struct {
	Value uint16
}

type TypeZero struct {
	AngleAxis [3]float32
}

func (p TypeZero) PacketType() int {
	return 0
}

func (p TypeZero) String() string {
	return fmt.Sprintf("%f, %f, %f", p.AngleAxis[0], p.AngleAxis[1], p.AngleAxis[2])
}

func (p TypeZero) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p TypeZero) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.AngleAxis[0]),
		fmt.Sprintf("%f", p.AngleAxis[1]),
		fmt.Sprintf("%f", p.AngleAxis[2]),
	}
}

type TypeOne struct {
	PixelExposureTime      int32
	RollingShutterSkewTime int32
}

func (p TypeOne) PacketType() int {
	return 1
}

func (p TypeOne) String() string {
	return fmt.Sprintf("%d, %d", p.PixelExposureTime, p.RollingShutterSkewTime)
}

func (p TypeOne) CSVHeader() []string {
	return []string{"Pixel Exposure Time (ns)", "Rolling Shutter Skew Time (ns)"}
}

func (p TypeOne) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%d", p.PixelExposureTime),
		fmt.Sprintf("%d", p.RollingShutterSkewTime),
	}
}

type TypeTwo struct {
	Gyro [3]float32
}

func (p TypeTwo) PacketType() int {
	return 2
}

func (p TypeTwo) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Gyro[0], p.Gyro[1], p.Gyro[2])
}

func (p TypeTwo) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p TypeTwo) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Gyro[0]),
		fmt.Sprintf("%f", p.Gyro[1]),
		fmt.Sprintf("%f", p.Gyro[2]),
	}
}

type TypeThree struct {
	Acceleration [3]float32
}

func (p TypeThree) PacketType() int {
	return 3
}

func (p TypeThree) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Acceleration[0], p.Acceleration[1], p.Acceleration[2])
}

func (p TypeThree) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p TypeThree) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Acceleration[0]),
		fmt.Sprintf("%f", p.Acceleration[1]),
		fmt.Sprintf("%f", p.Acceleration[2]),
	}
}

type TypeFour struct {
	Position [3]float32
}

func (p TypeFour) PacketType() int {
	return 4
}

func (p TypeFour) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Position[0], p.Position[1], p.Position[2])
}

func (p TypeFour) CSVHeader() []string {
	return []string{"Axis_X", "Axis_Y", "Axis_Z"}
}

func (p TypeFour) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Position[0]),
		fmt.Sprintf("%f", p.Position[1]),
		fmt.Sprintf("%f", p.Position[2]),
	}
}

type TypeFive struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (p TypeFive) PacketType() int {
	return 5
}

func (p TypeFive) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Latitude, p.Longitude, p.Altitude)
}

func (p TypeFive) CSVHeader() []string {
	return []string{"Latitude", "Longitude", "Altitude"}
}

func (p TypeFive) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.Latitude),
		fmt.Sprintf("%f", p.Longitude),
		fmt.Sprintf("%f", p.Altitude),
	}
}

type TypeSix struct {
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

func (p TypeSix) PacketType() int {
	return 6
}

func (p TypeSix) String() string {
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

func (p TypeSix) CSVHeader() []string {
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

func (p TypeSix) CSVFormat() []string {
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

type TypeSeven struct {
	MagneticField [3]float32
}

func (p TypeSeven) PacketType() int {
	return 7
}

func (p TypeSeven) String() string {
	return fmt.Sprintf("%f, %f, %f", p.MagneticField[0], p.MagneticField[1], p.MagneticField[2])
}

func (p TypeSeven) CSVHeader() []string {
	return []string{"B_Field_X", "B_Field_Y", "B_Field_Z"}
}

func (p TypeSeven) CSVFormat() []string {
	return []string{
		fmt.Sprintf("%f", p.MagneticField[0]),
		fmt.Sprintf("%f", p.MagneticField[1]),
		fmt.Sprintf("%f", p.MagneticField[2]),
	}
}
