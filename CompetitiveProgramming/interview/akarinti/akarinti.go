package akarinti

import (
	"strconv"
	"strings"
)

func StringChallenge(str string) string {
	r := []rune(str)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// Convert back to string
	return string(r)
}

func StringChallenge2(str string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	r := []rune(str)

	for i := 0; i < len(r); i++ {
		if (r[i] >= 'a' && r[i] <= 'z') || (r[i] >= 'A' && r[i] <= 'Z') {

			char := toLower(r[i])

			// Get next letter (wrapping z to a)
			nextChar := char + 1
			if nextChar > 'z' {
				nextChar = 'a'
			}

			// Check if next letter is a vowel
			if vowels[nextChar] {
				r[i] = toUpper(nextChar)
			} else {
				r[i] = nextChar
			}
		}
	}

	return string(r)
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func StringChallenge3(str string) string {
	r := []rune(str)
	for i := 0; i < len(str); i++ {
		r[i] = swap(r[i])
	}
	return string(r)

}

func swap(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	} else if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func MathChallenge(str string) string {
	res := 0

	// Iterate through each character from left to right
	for _, digit := range str {
		// Shift left (multiply by 2) and add current digit
		res = res*2 + int(digit-'0')
	}

	return strconv.Itoa(res)
}

// convert to string
func toString(res int) string {
	var digits []byte
	num := res
	for num > 0 {
		digit := byte(num%10) + '0'
		digits = append(digits, digit)
		num /= 10
	}

	// Reverse the digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}

func SearchingChallengeOld(str string) string {
	var res string = "-1"
	var maxCount int = 0
	for _, word := range strings.Split(str, " ") {
		mp := make(map[rune]int)
		mx := 0
		for _, char := range word {
			mp[char]++
			if mp[char] > 1 {
				mx = max(mx, mp[char])
			}
		}
		if mx > maxCount {
			maxCount = mx
			res = word
		}
	}
	return res
}

func SearchingChallenge(str string) string {
	if len(str) == 0 {
		return "-1"
	}

	var bestWord string = "-1"
	var maxRepeat int = 0
	var start int = 0
	var counts [128]int // Fixed-size array for ASCII characters

	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			// Process the word
			if i > start {
				currentMax := 0
				wordLen := i - start

				// Skip single-letter words
				if wordLen > 1 {
					// Count characters
					for j := start; j < i; j++ {
						char := str[j]
						counts[char]++
					}

					// Find max repetition
					for j := start; j < i; j++ {
						char := str[j]
						if counts[char] > currentMax {
							currentMax = counts[char]
						}
					}

					// Update result if this word has more repeating characters
					if currentMax > 1 && currentMax > maxRepeat {
						maxRepeat = currentMax
						bestWord = str[start:i]
					}

					// Clear counts array (only used positions)
					for j := start; j < i; j++ {
						counts[str[j]] = 0
					}
				}
			}
			start = i + 1
		}
	}

	return bestWord
}
