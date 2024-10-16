package leetcode

func ReverseInt(x int) int {
	ans := 0
	if x < -2147483648 || x > 2147483647 {
		return 0
	}

	for x != 0 {
		ans = ans*10 + x%10
		x = x / 10
	}

	if ans < -2147483648 || ans > 2147483647 {
		return 0
	}

	return ans
}

// ReverseInt__GPT reverses the digits of a signed 32-bit integer.
// It returns 0 if the reversed integer overflows the signed 32-bit range.
func ReverseInt__GPT(x int) int {
	const maxInt32 = 2147483647 // 2^31 - 1

	ans := 0
	sign := 1

	// Handle negative numbers
	if x < 0 {
		sign = -1
		x = -x // Make x positive for easier processing
	}

	for x != 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before updating ans
		if ans > (maxInt32-digit)/10 {
			return 0 // Overflow would occur
		}

		ans = ans*10 + digit
	}

	return ans * sign // Restore original sign
}
