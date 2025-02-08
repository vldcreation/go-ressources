package medium

import "fmt"

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{}

	for i := 0; i < len(matrix); i++ {
		numMatrix.calculatePrefixSum(matrix[i])
	}
	return numMatrix
}

func (n *NumMatrix) calculatePrefixSum(row []int) {
	prefixSum := make([]int, len(row))
	prefixSum[0] = row[0]
	for i := 1; i < len(row); i++ {
		prefixSum[i] = prefixSum[i-1] + row[i]
	}
	fmt.Printf("%v\n", prefixSum)
	n.prefixSum = append(n.prefixSum, prefixSum)
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			sum += n.prefixSum[i][col2]
		} else {
			sum += n.prefixSum[i][col2] - n.prefixSum[i][col1-1]
		}
	}
	return sum
}
