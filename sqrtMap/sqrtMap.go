package sqrtMap

import (
	"fmt"
	"log"
)

func PrintSqrtMap() {
	fmt.Print("Enter an integer: ")
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatalf("Error occurred: %s", err)
	}
	for k, v := range SqrtMap(num) {
		fmt.Printf("%d->%v\n", k, v)
	}

}

func SqrtMap(num int) map[int]int {
	var result = make(map[int]int)

	for i := 1; i <= num; i++ {
		result[i] = i * i
	}

	return result
}
