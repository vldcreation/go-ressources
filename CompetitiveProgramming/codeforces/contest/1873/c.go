package c1873

import "fmt"

func solvec() {
	N := 10
	ans := 0
	for i := 0; i < N; i++ {
		var S string
		fmt.Scan(&S)
		for j, c := range S {
			if c == 'X' {
				ans += min((11-abs(9-2*i))/2, (11-abs(9-2*j))/2)
			}
		}
	}
	fmt.Println(ans)
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

func SolveC() {
	var tt int
	fmt.Scan(&tt)
	for tt > 0 {
		solvec()
		tt--
	}
}
