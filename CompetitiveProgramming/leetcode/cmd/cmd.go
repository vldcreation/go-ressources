package cmd

import (
	"fmt"
	"io"
	"log"
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
- Usage: leetcode add-solution <difficulty> <solution_slug>
-	difficulty: easy, medium, hard
*/
func addSolution(cmd *cobra.Command, args []string) {
	// args: difficulty, solution_slug
	if len(args) < 2 {
		log.Fatal("Invalid number of arguments")
		return
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

	difficulty := args[0]
	solutionSlug := args[1]

	if solutionSlug == "" {
		log.Fatal("Invalid solution slug, it cannot be empty")
		return
	}

	// Create the solution file
	if !slices.Contains(MapSolutionLevel, difficulty) {
		log.Fatal("Invalid difficulty")
		return
	}

	// lookup the existing solution filename based on the solution slug
	// if it exists, return an error
	if !mustReplace {
		if _, err := os.Stat(RootPath + "/" + difficulty + "/" + solutionSlug + ".go"); err == nil {
			log.Fatal("The solution file already exists")
			return
		}
	}

	// open base stub file
	fileStub, err := os.Open(RootPath + "/" + difficulty + "/" + "/base.stub")
	if err != nil {
		log.Fatalf("Error opening file solution stub: %v\n", err)
		return
	}

	defer func() {
		if err := fileStub.Close(); err != nil {
			log.Fatalf("Error closing file solution stub: %v\n", err)
			return
		}
	}()

	// process the base stub file
	// proccess logic here

	// open the solution file
	solutionFile, err := os.Create(RootPath + "/" + difficulty + "/" + solutionSlug + ".go")
	if err != nil {
		log.Fatalf("Error creating file solution: %v\n", err)
		return
	}

	defer func() {
		if err := solutionFile.Close(); err != nil {
			fmt.Printf("Error closing file solution: %v\n", err)
			return
		}
	}()

	// Create the solution file
	fileStubBytes := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fileStub.Read(fileStubBytes)
		if err != nil && err != io.EOF {
			log.Fatalf("Error reading file solution stub: %v\n", err)
			return
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := solutionFile.Write(fileStubBytes[:n]); err != nil {
			log.Fatalf("Error writing file solution: %v\n", err)
			return
		}
	}

}
