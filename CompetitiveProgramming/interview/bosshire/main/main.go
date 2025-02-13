package main

import "fmt"

// 54321
// 5432
// 543
// 543
// 54
// 5
func main() {
	n := 6
	for i := 1; i <= n; i++ {
		if n%2 == 1 {
			// fmt.Printf("I: %d, n/2L %d ", i, n/2)
			if i == (n/2)+1 {
				for j := n; j >= i; j-- {
					fmt.Printf("%d", j)
				}
				println()
				for j := n; j >= i; j-- {
					fmt.Printf("%d", j)
				}
			} else {
				for j := n; j >= i; j-- {
					fmt.Printf("%d", j)
				}
			}
		} else {
			for j := n; j >= i; j-- {
				fmt.Printf("%d", j)
			}
		}
		fmt.Println()
	}
}
