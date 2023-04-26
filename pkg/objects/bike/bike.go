package bike

import "fmt"

type Bike struct {
	maker, model string
	wheelSize    float32
	travelMm     int
}

func (b Bike) String() string {
	return fmt.Sprintf("[Bike | %s %s | %dmm, %d\"]", b.maker, b.model, b.wheelSize, b.travelMm)
}

func TestData() []Bike {
	return []Bike{
		{"Rocky Mountain", "Thunderbolt", 27.5, 130},
		{"Yeti", "Betti", 27.5, 120},
		{"Pivot", "Mach 4 SL", 29, 100},
	}
}
