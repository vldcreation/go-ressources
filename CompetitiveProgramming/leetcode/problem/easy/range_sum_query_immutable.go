package easy

import "fmt"

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{}

	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = nums[0]
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	n.prefixSum = prefixSum
	return n
}

func (na *NumArray) SumRange(left int, right int) int {
	fmt.Printf("numArray: %+v\n", na.prefixSum)
	return na.prefixSum[right+1] - na.prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
