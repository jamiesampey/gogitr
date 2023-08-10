package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/leetCode"
)

func main() {
	str := "[(([{}]{}))]"
	fmt.Printf("Is balanced: %v\n", leetCode.IsBalancedParenTypes(str))
}
