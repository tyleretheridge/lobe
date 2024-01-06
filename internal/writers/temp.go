package writers

import (
	"fmt"
)

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
