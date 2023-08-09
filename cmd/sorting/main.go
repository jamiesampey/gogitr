package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/sorting"
	"sort"
)

func main() {
	var dogs = []sorting.Dog{
		{"Spot", "spaniel", "bw", 12},
		{"Gracie", "mutt", "tricolor", 4},
		{"Bingo", "heeler", "white", 2},
		{"Bluey", "heeler", "merle", 3},
	}

	fmt.Println(dogs)
	sort.Sort(sorting.ByAge(dogs))
	fmt.Println(dogs)
	sort.Sort(sorting.ByName(dogs))
	fmt.Println(dogs)
}
