package divisibility

import (
	"bytes"
	"fmt"
	"strconv"
)

func DivisibleAndNotDivisible(div, notDiv int) []int {
	var result []int

	for i := 1; i < 100; i++ {
		if i%div == 0 && i%notDiv != 0 {
			result = append(result, i)
		}
	}

	return result
}

func PrintSlice(slice []int) {
	var buf bytes.Buffer

	for _, val := range slice {
		if buf.Len() > 0 {
			buf.WriteString(", ")
		}

		buf.WriteString(strconv.Itoa(val))
	}

	fmt.Println(buf.String())
}
