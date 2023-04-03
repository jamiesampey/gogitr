package generics

import "fmt"

type SkiTour struct {
	Miles  float32
	VertFt int
	State  string
}

func (st SkiTour) String() string {
	return fmt.Sprintf("≈≈≈ %.1f miles, %d vert ft, %s ≈≈≈", st.Miles, st.VertFt, st.State)
}
