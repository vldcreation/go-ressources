package main

import "fmt"

func sum(v []int, ch chan int) {
	var total int = 0
	for _, num := range v {
		total += num
	}

	ch <- total
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{6, 7, 8, 9, 10}

	ans := make(chan int)
	go sum(numbers1, ans)
	go sum(numbers2, ans)

	sum1, sum2 := <-ans, <-ans
	fmt.Println(sum1, sum2, sum1+sum2)
}
