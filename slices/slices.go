package slices

import "fmt"

func TestSlice() {
	s := []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s)) // Output: [1 2 3] 3 3

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	fmt.Println(slice, len(slice), cap(slice)) // Output: [2 3 4] 3 4(5-1)

	a := make([]int, 3, 10)
	fmt.Println(a, len(a), cap(a)) // Output: [0 0 0] 3 10

	b := append(a, 1)
	c := append(a, 2)

	fmt.Println(b) // Output: [0 0 0 2]
	fmt.Println(c) // Output: [0 0 0 2]

	z := []int{2, 3, 5, 7, 11, 13}

	z = z[1:4]
	fmt.Println(z) // Output: [3 5 7]

	z = z[:2]
	fmt.Println(z) // Output: [3 5]

	z = z[1:]
	fmt.Println(z) // Output: [5]
}
