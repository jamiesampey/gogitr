package factorial

import "fmt"

func Factorial(num int) (uint64, error) {
	if num < 0 {
		return 0, fmt.Errorf("cannot perform factorial on a negative int %d", num)
	}

	if num == 0 {
		return 1, nil
	}

	factorial := uint64(num)
	for i := num - 1; i > 1; i-- {
		factorial *= uint64(i)
	}

	return factorial, nil
}

func PrintFactorial(num int) {
	factorial, err := Factorial(num)

	if err != nil {
		fmt.Printf("Received error: %s", err)
	} else {
		fmt.Printf("%d! = %d\n", num, factorial)
	}
}
