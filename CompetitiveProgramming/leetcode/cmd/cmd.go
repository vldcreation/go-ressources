package cmd

import (
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

const (
	RootPath = "CompetitiveProgramming/leetcode"
)

const (
	LevelEasy   = "easy"
	LevelMedium = "medium"
	LevelHard   = "hard"
)

var (
	MapSolutionLevel = []string{
		LevelEasy,
		LevelMedium,
		LevelHard,
	}
)

func NewCMD() *cobra.Command {
	var (
		rootCmd = &cobra.Command{
			Use:   "leetcode",
			Short: "leetcode is a CLI tool to manage ressources",
			Long:  "leetcode is a CLI tool to manage ressources",
		}
	)

	addSolutionCmd := &cobra.Command{
		Use:   "add-solution",
		Short: "Add a solution to a leetcode problem",
		Long:  "Add a solution to a leetcode problem",
		Run:   addSolution,
		ValidArgs: []string{
			"replace",
		},
	}
	addSolutionCmd.Flags().BoolP("replace", "r", false, "Replace the solution file if it exists")

	// Add subcommands
	rootCmd.AddCommand(addSolutionCmd)
	return rootCmd
}

/*
- Add a solution to a leetcode problem
- Usage: leetcode add-solution <category> <solution_slug>
-	category: easy, medium, hard
*/
func addSolution(cmd *cobra.Command, args []string) {
	// args: category, solution_slug
	if len(args) < 2 {
		panic("Not enough arguments")
	}

	// register the flag --replace
	// if the flag is present, the solution file will be replaced
	// if the flag is not present, the solution file will not be replaced
	// default: false

	mustReplace := false
	flagReplace := cmd.Flag("replace")
	if flagReplace != nil {
		mustReplace = flagReplace.Value.String() == "true"
	}

	category := args[0]
	solutionSlug := args[1]

	if solutionSlug == "" {
		panic("Invalid solution slug, it cannot be empty")
	}

	// Create the solution file
	if !slices.Contains(MapSolutionLevel, category) {
		panic("Invalid category")
	}

	// lookup the existing solution filename based on the solution slug
	// if it exists, return an error
	if !mustReplace {
		if _, err := os.Stat(RootPath + "/" + category + "/" + solutionSlug + ".go"); err == nil {
			panic("Solution already exists")
		}
	}

	// open base stub file
	fileStub, err := os.Open(RootPath + "/" + category + "/" + "/base.stub")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fileStub.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
			panic(err)
		}
	}()

	// process the base stub file
	// proccess logic here

	// open the solution file
	solutionFile, err := os.Create(RootPath + "/" + category + "/" + solutionSlug + ".go")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := solutionFile.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
			panic(err)
		}
	}()

	// Create the solution file
	fileStubBytes := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fileStub.Read(fileStubBytes)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := solutionFile.Write(fileStubBytes[:n]); err != nil {
			panic(err)
		}
	}

}
