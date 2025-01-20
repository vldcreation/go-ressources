package function

// Recursion
func Recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursion(n-1)
}
