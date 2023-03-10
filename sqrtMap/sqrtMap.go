package sqrtMap

import (
	"fmt"
	"log"
	"sort"
)

func PrintSqrtMap() {
	fmt.Print("Enter an integer: ")
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}

	var squares = SqrtMap(num)

	var keysAsc []int
	for k := range squares {
		keysAsc = append(keysAsc, k)
	}
	sort.Ints(keysAsc)

	for k := range keysAsc {
		fmt.Printf("%d->%d\n", k, squares[k])
	}
}

func SqrtMap(num int) map[int]int {
	var result = make(map[int]int)

	for i := 1; i <= num; i++ {
		result[i] = i * i
	}

	return result
}
