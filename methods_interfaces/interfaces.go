package methods_interfaces

// An interface type is defined by a set of methods.
// A value of interface type can hold any value that implements those methods.

// The empty interface, interface{}, may hold values of any type.
// Every type implements at least zero methods, so the empty interface may hold values of any type.

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func test() {
	var i I = T{"hello"}
	i.M()
}

// Type assertion
func assertions() {
	var i interface{} = "hello"

	s := i.(string) // ok, s == "hello"
	fmt.Println(s)

	s, ok := i.(string) // s == "hello", ok == true
	fmt.Println(s, ok)

	n, ok := i.(int) // n == 0, ok == false
	fmt.Println(n, ok)

	// n = i.(int) // panic: interface conversion: interface {} is string, not int
}

// Type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
