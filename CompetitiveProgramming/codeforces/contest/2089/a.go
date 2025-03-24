package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func solvea() {
	var n int
	fmt.Scan(&n)
	perm := make([]int, n)
	if n >= 1 {
		perm[0] = 2
	}
	if n >= 2 {
		perm[1] = 1
	}
	for i := 2; i < n; i++ {
		perm[i] = i + 1
	}
	for i, v := range perm {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		solvea()
		t--
	}
}
