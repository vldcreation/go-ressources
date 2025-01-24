package oy

func Case1(a, b, k int) int {
	// var res int
	// if a%k == 0 {
	// 	res = (b / k) - ((a - 1) / k) + 1
	// } else {
	// 	res = (b / k) - ((a - 1) / k)
	// }

	return (b / k) - ((a - 1) / k)
}
