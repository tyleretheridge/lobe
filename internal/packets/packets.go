package packets

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

type TypeTwo struct {
	Gyro [3]float32
}

func (p TypeTwo) PacketType() int {
	return 2
}

func (p TypeTwo) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Gyro[0], p.Gyro[1], p.Gyro[2])
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

type TypeFour struct {
	Position [3]float32
}

func (p TypeFour) PacketType() int {
	return 4
}

func (p TypeFour) String() string {
	return fmt.Sprintf("%f, %f, %f", p.Position[0], p.Position[1], p.Position[2])
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

type TypeSeven struct {
	MagneticField [3]float32
}

func (p TypeSeven) PacketType() int {
	return 7
}

func (p TypeSeven) String() string {
	return fmt.Sprintf("%f, %f, %f", p.MagneticField[0], p.MagneticField[1], p.MagneticField[2])
}
