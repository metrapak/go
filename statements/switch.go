package statements

import (
	"fmt"
	"runtime"
)

func TestSwitch() {
	fmt.Println("Go runs on")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
