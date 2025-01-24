package main

import (
	"fmt"
	"os"
)

func main() {
	// Open input file
	inputFile, err := os.Open("input.in")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Open output file for writing
	outputFile, err := os.Create("input.out")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Redirect standard input to read from the input file
	stdin := os.Stdin
	os.Stdin = inputFile

	// Restore original standard input after processing
	defer func() {
		os.Stdin = stdin
	}()

	// Redirect standard output to write to the output file
	stdout := os.Stdout
	os.Stdout = outputFile

	defer func() {
		os.Stdout = stdout
		outputFile.Sync() // Flush the buffer
	}()

	// Call your original program logic here (assuming it's defined in a function)
	solveDivisibilityProblem()
}

// Assuming your original program logic is in a separate function
func solveDivisibilityProblem() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var a, b, k int32
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&k)

		count := (b / k) - ((a - 1) / k)
		fmt.Println(fmt.Sprintf("Case %d: %d", i, count))
	}
}
