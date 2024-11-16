package leetcode

import (
	"fmt"
)

func ZigzagConversion(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	currentRow := 0
	goingDown := false

	for _, char := range s {
		rows[currentRow] += string(char)

		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}

	}

	ans := ""
	for _, row := range rows {
		ans += row
	}

	return ans
}

func ZigzagConversion__BEST(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)

	i := 0

	goingDown := false
	row := 0
	for i < len(s) {
		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}

		rows[row] += s[i : i+1]
		fmt.Printf("row: %d, i: %d, s[i]: %s\n", row, i, s[i:i+1])

		if goingDown {
			row++
		} else {
			row--
		}
		i++
	}

	zigzagged := ""
	for _, row := range rows {
		zigzagged += row
	}

	return zigzagged
}
