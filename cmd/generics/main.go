package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/generics"
	"log"
)

type skiTour struct {
	miles  float32
	vertFt int
	state  string
}

func main() {
	tourStack := new(generics.Stack)
	tourStack.Push(skiTour{miles: 5.4, vertFt: 3000, state: "MT"})
	tourStack.Push(skiTour{miles: 7.2, vertFt: 3500, state: "CO"})
	tourStack.Push(skiTour{miles: 9.8, vertFt: 4000, state: "WY"})

	popped, err := tourStack.Pop()
	if err != nil {
		log.Fatal("Received error on pop")
	}
	fmt.Printf("popped tour is %v", popped)
}
