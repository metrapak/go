package variadic_functions

import "fmt"

func Sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}
