package divisibility

import (
	"testing"
)

func TestDivisibleBasic(t *testing.T) {
	var div, notDiv = 5, 7
	for _, val := range DivisibleAndNotDivisible(div, notDiv) {
		if val%div != 0 || val%notDiv == 0 {
			t.Errorf("Found bad number: %d", val)
		}
	}
}

func TestDivisibleNotOverMax(t *testing.T) {
	var div, notDiv = 2, 3
	for _, val := range DivisibleAndNotDivisible(div, notDiv) {
		if val > 100 {
			t.Errorf("Found a number over 100: %d", val)
		}
	}
}
