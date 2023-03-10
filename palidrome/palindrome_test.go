package palidrome

import "testing"

func TestIsPalidrome(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"madam", true},
		{"nurses run", true},
		{"jamie", false},
		{"I did, did I", true},
		{"Eva, can I see bees in a cave", true},
		{"Eva, can I sea bees in a cave", false},
	}

	for _, test := range tests {
		if IsPalidrome(test.str) != test.expected {
			t.Errorf("IsPalindrome('%s') returned %t", test.str, !test.expected)
		}
	}
}
