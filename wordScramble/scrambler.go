package wordScramble

import (
	"fmt"
	"math/rand"
	"strings"
)

func ScrambleWords(sentence string) {
	var words []string
	for _, word := range strings.Split(sentence, " ") {
		if trimmedWord := strings.TrimSpace(word); len(trimmedWord) > 0 {
			words = append(words, trimmedWord)
		}
	}

	fmt.Printf("Scrambling %v ...\n", words)

	var rndWords []string
	wordsLeft := len(words)

	for wordsLeft > 0 {
		rnd := rand.Intn(wordsLeft)
		rndWords = append(rndWords, words[rnd])
		words, wordsLeft = removeIndex(words, rnd)
	}

	fmt.Printf("Scrambled %v", rndWords)
}

func removeIndex(words []string, i int) ([]string, int) {
	n := len(words)
	if i < n-1 {
		words[i] = words[n-1]
	}
	return words[:n-1], n - 1
}
