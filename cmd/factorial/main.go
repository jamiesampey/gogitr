package main

import (
	"flag"
	"jamiesampey.com/gogitr/pkg/factorial"
	"log"
	"strconv"
)

// Usage: go run cmd/factorial/main.go -- 9
// Output: 9! = 362880
func main() {
	flag.Parse()
	input, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	factorial.PrintFactorial(int(input))
}
