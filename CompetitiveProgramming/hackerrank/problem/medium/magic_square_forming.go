package medium

func FormingMagicSquare(s [][]int32) int32 {
	// All possible 3x3 magic squares
	possibleSquares := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	minCost := int32(81) // Maximum possible cost

	// Try each possible magic square
	for _, square := range possibleSquares {
		cost := int32(0)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cost += abs(s[i][j] - square[i][j])
			}
		}
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
