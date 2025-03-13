package c1873

import (
	"fmt"
)

func solvea() {
	var s string
	fmt.Scan(&s)
	if s != "bca" && s != "cab" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func a() {
	var tt int
	fmt.Scan(&tt)
	for tt > 0 {
		solvea()
		tt--
	}
}
