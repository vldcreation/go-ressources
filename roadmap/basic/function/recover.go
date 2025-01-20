package function

import "fmt"

// Recover
func Recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	panic("This is panic")
}
