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

	fmt.Println("\nPointer to a pointer to a string:")
	fmt.Printf("&ptr is a [%T] with value [%v]\n", &ptr, &ptr)
	fmt.Printf("*&ptr is a [%T] with value [%v]\n", *&ptr, *&ptr)
	fmt.Printf("**&ptr is a [%T] with value [%v]\n", **&ptr, **&ptr)

	fmt.Println("\nChanging the pointed to value:")
	*ptr = "hello world"
	fmt.Printf("data now has value [%s]\n", data)
	data = "hello mars"
	fmt.Printf("*ptr now has value [%v]\n", *ptr)
}
