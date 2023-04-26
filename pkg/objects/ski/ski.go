package ski

import (
	"fmt"
)

type Ski struct {
	maker, model      string
	lengthCm, widthMm int
	radiusM           float32
}

func (s Ski) String() string {
	return fmt.Sprintf("[Ski | %s %s | %dcm, %dmm, %.1fm]", s.maker, s.model, s.lengthCm, s.widthMm, s.radiusM)
}

func TestData() []Ski {
	return []Ski{
		{"DPS", "Pagoda Tour", 179, 106, 19.0},
		{"Blizzard", "Rustler 11", 180, 112, 19.1},
		{"Salomon", "QST Stella", 173, 106, 16.2},
	}
}
