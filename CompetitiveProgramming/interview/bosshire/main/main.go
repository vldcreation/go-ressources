package main

import (
	"fmt"
	"os"
	"strconv"
)

// n = 5
// 54321
// 5432
// 543
// 543
// 54
// 5
// n = 6
// 654321
// 65432
// 6543
// 654
// 65
// 6
func main() {
	if len(os.Args) < 1 {
		panic("Please pass n number")
	}
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	for i := int64(1); i <= n; i++ {
		if n%2 == 1 {
			if i == (n/2)+1 {
				for j := n; j >= i; j-- {
					fmt.Printf("%d", j)
				}
				println()
			}
		}
		for j := n; j >= i; j-- {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}
