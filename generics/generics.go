package generics

import "fmt"

// A generic function is a function that can be used with different types.

func index[T comparable](s []T, v T) int {
	for i, t := range s {
		if t == v {
			return i
		}
	}
	return -1
}

func testGenerics() {
	fmt.Println(index([]string{"a", "b", "c"}, "b"))
	fmt.Println(index([]int{1, 2, 3}, 2))
}
