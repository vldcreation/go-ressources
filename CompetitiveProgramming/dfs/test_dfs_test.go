package dfs

import (
	"testing"

	"github.com/vldcration/go-ressources/util"
)

func TestGetAllPossibilities(t *testing.T) {
	tests := []struct {
		target         string
		word_bank      []string
		expectedResult [][]string
	}{
		{
			"purple",
			[]string{"purp", "p", "ur", "le", "purpl"},
			[][]string{
				{"purp", "le"},
				{"p", "ur", "p", "le"},
			},
		},
		{
			"blindman",
			[]string{"bl", "blind", "i", "nd", "man", "m", "an"},
			[][]string{
				{"bl", "i", "nd", "man"},
				{"bl", "i", "nd", "m", "an"},
				{"blind", "man"},
				{"blind", "m", "an"},
			},
		},
	}

	for _, test := range tests {
		result := getAllPossibilities(test.target, test.word_bank)
		if !util.Compare2DSlice(result, test.expectedResult) {
			t.Errorf("Expected %v, got %v", test.expectedResult, result)
		}
	}
}
