package leetcode

import (
	"log"

	"github.com/vldcration/go-ressources/util"
)

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

func NewLeetcode(opt ...Opt) *Leetcode {
	l := &Leetcode{}
	for _, o := range opt {
		o(l)
	}

	return l
}

func WithProblems(problems []Problem) Opt {
	return func(l *Leetcode) {
		l.Problems = problems
	}
}

func MustLoadProblems() Opt {
	return func(l *Leetcode) {
		var pathToProblems = util.RootPath() + "/data/solution_bank.json"
		if err := util.LoadJSON(pathToProblems, l); err != nil {
			log.Printf("Error load problems: %v", err)
			panic(err)
		}
	}
}

func (l *Leetcode) AddProblem(p Problem) {
	l.Problems = append(l.Problems, p)
}
