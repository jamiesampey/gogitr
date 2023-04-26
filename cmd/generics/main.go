package main

import (
	"jamiesampey.com/gogitr/pkg/generics"
	"jamiesampey.com/gogitr/pkg/objects/bike"
	"jamiesampey.com/gogitr/pkg/objects/ski"
)

func main() {
	skiStack := new(generics.Stack[ski.Ski])
	for _, item := range ski.TestData() {
		skiStack.Push(item)
	}
	skiStack.PopAll()

	bikeStack := new(generics.Stack[bike.Bike])
	for _, item := range bike.TestData() {
		bikeStack.Push(item)
	}
	bikeStack.PopAll()
}
