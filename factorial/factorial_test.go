package factorial

import "testing"

func TestFactorial(t *testing.T) {
	result, err := Factorial(7)
	if result != 5040 || err != nil {
		t.Errorf("Received unexpected result")
	}
}

func TestFactorialZero(t *testing.T) {
	result, err := Factorial(0)
	if result != 1 || err != nil {
		t.Errorf("Received unexpected result")
	}
}

func TestFactorialNegative(t *testing.T) {
	result, err := Factorial(-3)
	if result != 0 || err == nil {
		t.Errorf("Received unexpected result")
	}
}
