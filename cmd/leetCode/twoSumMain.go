package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/leetCode"
)

func main() {
	indices := leetCode.TwoSumWithMap([]int{3, 5, 6, 3, 2, 1}, 41)

	if len(indices) > 0 {
		fmt.Printf("Found sum at {%d, %d}\n", indices[0], indices[1])
	} else {
		fmt.Println("Did not find the sum")
	}
}
