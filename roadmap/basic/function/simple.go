package function

import "fmt"

// Simple Function
func SimpleFunction() {
	fmt.Printf("This is Simple function without any arguments and return type")
}

func SimpleFunctionWithArguments(a int, b int) {
	println("This is Simple Function with arguments a and b: ", a, b)
}

var SimpleFunctionWithReturnResult int = 10

func SimpleFunctionWithReturn() int {
	return SimpleFunctionWithReturnResult
}

func SimpleFunctionWithArgumentsAndReturn(a int, b int) int {
	return a + b
}
