package sqrtMap

import (
	"fmt"
	"log"
	"sort"
)

func RunSqrtMap() {
	var predicate = func(k int) bool {
		return k >= 5 && k%2 == 1
	}
	PrintSqrtMap(predicate)
}

func PrintSqrtMap(printFilter func(int) bool) {
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

	for _, k := range keysAsc {
		if printFilter(k) {
			fmt.Printf("%d->%d\n", k, squares[k])
		}
	}
}

func SqrtMap(num int) map[int]int {
	var result = make(map[int]int)

	for i := 1; i <= num; i++ {
		result[i] = i * i
	}

	return result
}
