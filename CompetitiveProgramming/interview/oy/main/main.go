package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var a, b, k int32
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&k)

		count := (b / k) - ((a - 1) / k)
		fmt.Println(fmt.Sprintf("Case %d: %d", i, count))
	}
}
