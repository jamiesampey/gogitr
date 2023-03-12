package wordScramble

import (
	"fmt"
	"math/rand"
	"strings"
)

func ScrambleWords(sentence string) []string {
	words := TrimSentence(sentence)
	fmt.Printf("\nScrambling %v ...\n", words)

	var rndWords []string
	wordsLeft := len(words)

	for wordsLeft > 0 {
		rnd := rand.Intn(wordsLeft)
		rndWords = append(rndWords, words[rnd])
		words, wordsLeft = removeIndex(words, rnd)
	}

	fmt.Printf("Scrambled %v\n", rndWords)
	return rndWords
}

func TrimSentence(rawSentence string) []string {
	var trimmedWords []string
	for _, word := range strings.Split(rawSentence, " ") {
		if trimmedWord := strings.TrimSpace(word); len(trimmedWord) > 0 {
			trimmedWords = append(trimmedWords, trimmedWord)
		}
	}

	return trimmedWords
}

func removeIndex(words []string, i int) ([]string, int) {
	n := len(words)
	if i < n-1 {
		words[i] = words[n-1]
	}
	return words[:n-1], n - 1
}
