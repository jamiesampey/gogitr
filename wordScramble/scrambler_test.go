package wordScramble

import "testing"

func TestScrambledLen(t *testing.T) {
	tests := []string{
		"I'm gonna live    forever, I'm gonna cross      that river",
		"may the wind take your   troubles away",
		"I've got a \tpeaceful easy feeling",
		"mary 	had a   little      lamb",
	}

	for _, test := range tests {
		trimmedSentence := TrimSentence(test)
		scrambledWords := ScrambleWords(test)
		if len(scrambledWords) != len(trimmedSentence) {
			t.Errorf("Received a scrambled string of unexpected length for %v", trimmedSentence)
		}
	}
}
