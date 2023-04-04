package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/generics"
	"jamiesampey.com/gogitr/pkg/generics/ski"
	"log"
)

func main() {
	tourStack := new(generics.Stack[ski.Ski])
	tourStack.Push(ski.New("DPS", "Pagoda Tour", 179, 106, 19.0))
	tourStack.Push(ski.New("Blizzard", "Rustler 11", 180, 112, 19.1))
	tourStack.Push(ski.New("Salomon", "QST Stella", 173, 106, 16.2))

	item, err := tourStack.Pop()
	if err != nil {
		log.Fatal("Received error on pop")
	}
	fmt.Printf("Popped item is %s", item.String())
}
