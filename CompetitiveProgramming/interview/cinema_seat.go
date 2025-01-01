package interview

import "fmt"

// col 1 = 3
// col 5 = 14
// col 18 = 18

// col = 3
// row = 5

func Solution(col int, row int) int {
	tCol := 0
	if col <= 4 {
		tCol = 4
	} else if col > 4 && col <= 14 {
		tCol = 14
	} else if col > 14 && col <= 18 {
		tCol = 18
	}

	div := tCol - col + 1
	sum := div * row
	ans := sum - 1
	return ans
}

func main() {
	col1 := 1
	row1 := 7
	expectedOutput := 27
	if got := Solution(col1, row1); got != expectedOutput {
		fmt.Printf("expected : %+v but got: %+v\n", expectedOutput, got)
	}
}
