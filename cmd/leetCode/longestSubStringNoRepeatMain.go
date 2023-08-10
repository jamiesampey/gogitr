package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/leetCode"
)

func main() {
	l := leetCode.LongestSubstringWithNoRepeatChars("abciibcdebb")
	fmt.Printf("longest substring with non-repeating chars is %d\n", l)
}
