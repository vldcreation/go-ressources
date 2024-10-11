package leetcode

/*
// @Author: Vicktor Desrony
// @filename: leetcode.go
// WIP: Still working on this
*/

type Leetcode struct {
	Problems []Problem `json:"problems"`
}

type Problem struct {
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	Path       string   `json:"Path"`
	Difficulty string   `json:"difficulty"`
	Solution   Solution `json:"solution"`
}

type Solution struct {
	Language string `json:"language"`
	Path     string `json:"path"`
}

func NewLeetcode() *Leetcode {
	return &Leetcode{}
}

func (l *Leetcode) AddProblem(p Problem) {
	l.Problems = append(l.Problems, p)
}
