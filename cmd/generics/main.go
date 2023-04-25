package main

import (
	"fmt"
	"jamiesampey.com/gogitr/pkg/generics"
	"jamiesampey.com/gogitr/pkg/generics/ski"
)

func main() {
	skiStack := new(generics.Stack[ski.Ski])
	skiStack.Push(ski.New("DPS", "Pagoda Tour", 179, 106, 19.0))
	skiStack.Push(ski.New("Blizzard", "Rustler 11", 180, 112, 19.1))
	skiStack.Push(ski.New("Salomon", "QST Stella", 173, 106, 16.2))

	//var err error
	for {
		item, err := skiStack.Pop()
		if err != nil {
			fmt.Printf("err is '%v'\n", err)
			break
		}
		fmt.Printf("Popped item is %s\n", item.String())
	}

}
