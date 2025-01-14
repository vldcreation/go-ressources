package ifglife

/* MinimumFlips
 * Start with an initial string of zeros.
 * Choose any digit to flip. When a digit is flipped, its value and those to the right switch state between 0 and 1.
 *
 * Example:
 * target = "01011"
 * Initial = "00000"
 * Flip the 3rd digit: 00011
 * Flip the 2nd digit: 01000
 * Flip the 4th digit: 01011
 * 3 flips are required to reach the target. then return 3
 */

func MinimumFlips(input string) int {
	ans := 0

	curState := '0'

	for _, c := range input {
		if c != curState {
			ans++
			curState = c
		}
	}

	return ans
}
