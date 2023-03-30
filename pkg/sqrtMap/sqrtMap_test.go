package sqrtMap

import (
	"testing"
)

func TestSqrtMap(t *testing.T) {
	num := 5
	var result = SqrtMap(num)
	for i := 1; i <= num; i++ {
		val, exists := result[i]
		if !exists {
			t.Errorf("expected to find a sqrt value for index %d", i)
		} else if val != i*i {
			t.Errorf("incorrect sqrt value %d for index %d", val, i)
		}
	}
}
