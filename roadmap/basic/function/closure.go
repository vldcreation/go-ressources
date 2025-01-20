package function

import "fmt"

// Closure
func Closure() func() {
	x := 10
	return func() {
		fmt.Println(x)
	}
}
