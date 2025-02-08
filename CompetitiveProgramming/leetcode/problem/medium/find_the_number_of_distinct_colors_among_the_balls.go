package medium

func QueryResults(limit int, queries [][]int) []int {
	ans := []int{}
	hashMap := map[int]int{}
	mapColor := map[int]int{}

	for _, query := range queries {
		mapColor[query[1]]++

		// Check if x already exists in mapping and adjust the count\
		if color1, ok := hashMap[query[0]]; ok {
			mapColor[color1]--
			if mapColor[color1] == 0 {
				// If count of previous y becomes zero, remove it
				delete(mapColor, color1)
			}
		}

		// Update mapping for x to the new y
		hashMap[query[0]] = query[1]

		// Add the current distinct count of y's to the result
		ans = append(ans, len(mapColor))
	}

	return ans
}
