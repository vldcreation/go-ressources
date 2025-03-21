package easy

func IsPalindrome(x int) bool {
	var reverse = 0
	var copy = x

	//The loop break when the copy of original number becomes zero
	//Also negative number cannot be a palindrome
	for copy > 0 {
		digit := copy % 10
		reverse = reverse*10 + digit
		copy = copy / 10
	}

	return reverse == x
}
