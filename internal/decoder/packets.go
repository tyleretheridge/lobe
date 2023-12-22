package decoder

import "fmt"

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

type PacketTypeOne struct {
	PixelExposureTime      int32
	RollingShutterSkewTime int32
}

func (p PacketTypeOne) PacketType() int {
	return 1
}

type PacketTypeTwo struct {
	Gyro [3]float32
}

func (p PacketTypeTwo) PacketType() int {
	return 2
}

type PacketTypeThree struct {
	Acceleration [3]float32
}

func (p PacketTypeThree) PacketType() int {
	return 3
}

type PacketTypeFour struct {
	Position [3]float32
}

func (p PacketTypeFour) PacketType() int {
	return 4
}

type PacketTypeFive struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (p PacketTypeFive) PacketType() int {
	return 5
}

type PacketTypeSix struct {
	TimeGPSEpoch       float64 //seconds
	GPXFixType         int32   // 0, 2(D), 3(D)
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

type PacketTypeSeven struct {
	MagneticField [3]float32
}

func (p PacketTypeSeven) PacketType() int {
	return 7
}

type DecodedPacket interface {
	PacketType() int
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
