package function

import "fmt"

// Defer
func Defer() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
