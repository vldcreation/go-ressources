package bv

import (
	"fmt"
	"math"
	"sort"
)

// testcase 1
// capacity: 1, 2, 3, 4
// numServer: 4
// expected: 3

// testcase 2
// capacity: 4, 2, 1
// numServer: 1, 1, 1
// expected: 0

func GetMaximumEfficiency(capacity []int32, numServer []int32) int64 {
	// Sort the capacity array in descending order
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})
	fmt.Printf("capacity: %v\n", capacity)

	// Sort the numServer array in descending order
	sort.Slice(numServer, func(i, j int) bool {
		return numServer[i] > numServer[j]
	})
	fmt.Printf("numServer: %v\n", numServer)

	var ans int64 = 0
	var mapUsed = make(map[int32]int)

	for _, num := range numServer {
		cur := math.MinInt32
		for _, cap := range capacity {
			if _, ok := mapUsed[cap]; ok {
				continue
			}

			res := int(cap - num)
			fmt.Println("res: ", res)
			cur = max(cur, res)
			// cur = max(cur, res)
			// if res >= 0 {
			// 	mapUsed[cap] = j
			// }
		}

		ans = int64(cur)
	}

	if ans < 0 {
		return ans * -1
	}

	return ans
}
