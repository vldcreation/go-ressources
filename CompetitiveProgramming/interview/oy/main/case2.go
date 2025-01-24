package main

import (
	"fmt"
	"os"
)

func main() {
	// Open input file
	inputFile, err := os.Open("input2.in")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Open output file for writing
	outputFile, err := os.Create("input2.out")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Redirect standard input to read from the input file
	stdin := os.Stdin
	os.Stdin = inputFile
	defer func() {
		os.Stdin = stdin
	}()

	// Redirect standard output to write to the output file
	stdout := os.Stdout
	os.Stdout = outputFile
	defer func() {
		os.Stdout = stdout
		outputFile.Sync()
	}()

	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var n, m int
		fmt.Scan(&n)
		fmt.Scan(&m)

		grid := make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[j])
		}

		var w string
		fmt.Scan(&w)

		count := findWordCount(grid, w)
		fmt.Printf("Case %d: %d\n", i, count)
	}
}

func findWordCount(grid []string, word string) int {
	count := 0
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Horizontal, Vertical
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonal
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for _, dir := range directions {
				if checkWordInDirection(grid, i, j, dir, word) {
					count++
				}
			}
		}
	}
	return count
}

func checkWordInDirection(grid []string, row, col int, dir []int, word string) bool {
	for i := 0; i < len(word); i++ {
		newRow, newCol := row+i*dir[0], col+i*dir[1]
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || string(grid[newRow][newCol]) != string(word[i]) {
			return false
		}
	}
	return true
}
