package main

import (
	"fmt"
	"main/statements"
)

var c, python, java bool

var i, j int = 1, 2

func add(x int, y int) int {
	return x + y
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(add(2, 3))
	fmt.Println(split(17))
	statements.TestFor()
}
