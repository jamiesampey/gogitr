package pointers

import (
	"fmt"
)

func BasicPointerTypes() {
	data := "foobar"

	var ptr *string
	ptr = &data

	fmt.Println("\nPointer to a string:")
	fmt.Printf("data is a [%T] with value [%s]\n", data, data)
	fmt.Printf("ptr is a [%T] with value [%p]\n", ptr, ptr)
	fmt.Printf("*ptr is a [%T] with value [%v]\n", *ptr, *ptr)

	fmt.Println("\nPointer to a pointer:")
	fmt.Printf("&ptr is a [%T] with value [%v]\n", &ptr, &ptr)
	fmt.Printf("*&ptr is a [%T] with value [%v]\n", *&ptr, *&ptr)
	fmt.Printf("**&ptr is a [%T] with value [%v]\n", **&ptr, **&ptr)
}
