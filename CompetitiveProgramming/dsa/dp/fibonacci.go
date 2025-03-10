package dp

import "slices"

func NaiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return NaiveFibonacci(n-1) + NaiveFibonacci(n-2)
}

func MemoizedFibonacci(n int, memo []int) int {
	if n <= 1 {
		return n
	}

	if !slices.Contains(memo, n) {
		memo[n] = MemoizedFibonacci(n-1, memo) + MemoizedFibonacci(n-2, memo)
	}

	return MemoizedFibonacci(n, memo)
}

func TabulatedFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func SpaceOptimizedFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1

	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}

	return b
}
