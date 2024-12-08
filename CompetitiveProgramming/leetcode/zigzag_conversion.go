package leetcode

func ZigzagConvert(s string, numRows int) string {
	// Corner Case (Only one row)
	if numRows == 1 {
		return s
	}

	len := len(s)

	var arr = make([][]byte, numRows)

	row := 0
	var down bool

	for i := range arr {
		arr[i] = make([]byte, len)
	}

	// Traverse through given string
	for i := 0; i < len; i++ {
		arr[row][i] = s[i]

		if row == len-1 {
			down = false
		} else if row == 0 {
			down = true
		}

		// If direction is down, increment, else decrement
		if down {
			row++
		} else {
			row--
		}
	}

	ans := ""
	for _, bt := range arr {
		ans += string(bt)
	}

	return ans
}
