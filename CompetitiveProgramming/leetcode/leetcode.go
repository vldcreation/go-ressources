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
	ID         int        `json:"id"` // Leetcode ID or Problem ID
	Tittle     string     `json:"tittle"`
	Difficulty string     `json:"difficulty"`
	Slug       string     `json:"slug"`
	Url        string     `json:"url" comment:"URL to the problem"`
	Solutions  []Solution `json:"solutions"`
}

type Solution struct {
	Language string `json:"language"`
	Url      string `json:"url"`
}

type Opt func(*Leetcode)

func NewLeetcode(...Opt) *Leetcode {
	return &Leetcode{}
}

func (l *Leetcode) AddProblem(p Problem) {
	l.Problems = append(l.Problems, p)
}
