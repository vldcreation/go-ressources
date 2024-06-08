package dfs

import (
	"strings"
)

/*
eg1.
	target: "purple"
	result: ["purp", "le"],
			["p", "ur", "p", "le"]
*/

func getAllPossibilities(target string, chars []string) [][]string {
	result := [][]string{}
	dfs(target, chars, []string{}, &result)
	return result
}

func dfs(target string, chars []string, current []string, result *[][]string) {
	if target == "" {
		*result = append(*result, append([]string{}, current...))
		return
	}

	for _, char := range chars {
		if strings.HasPrefix(target, char) {
			dfs(strings.TrimPrefix(target, char), chars, append(current, char), result)
		}
	}
}
