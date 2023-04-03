package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/generics"
	"log"
)

func main() {
	tourStack := new(generics.Stack)
	tourStack.Push(generics.SkiTour{Miles: 5.4, VertFt: 3000, State: "MT"})
	tourStack.Push(generics.SkiTour{Miles: 7.2, VertFt: 3500, State: "CO"})
	tourStack.Push(generics.SkiTour{Miles: 9.8, VertFt: 4000, State: "WY"})

	popped, err := tourStack.Pop()
	if err != nil {
		log.Fatal("Received error on pop")
	}
	fmt.Printf("popped tour is %s", popped.String())
}
