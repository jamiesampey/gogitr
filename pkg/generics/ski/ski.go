package ski

import "fmt"

type Ski struct {
	maker, model      string
	lengthCm, widthMm int
	radiusM           float32
}

func (s Ski) String() string {
	return fmt.Sprintf("[SKI | %s %s | %dcm, %dmm, %.1fm]", s.maker, s.model, s.lengthCm, s.widthMm, s.radiusM)
}

func New(maker, model string, length, width int, radius float32) Ski {
	return Ski{maker, model, length, width, radius}
}
