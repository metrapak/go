package statements

import "fmt"

func TestFor() {
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	count := 5
	for i := range count {
		fmt.Println(i)
	}
}
