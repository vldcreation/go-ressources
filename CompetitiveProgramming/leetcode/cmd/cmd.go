package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vldcreation/go-ressources/util"
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
	addSolutionCmd.Flags().StringP("difficulty", "d", "", "Difficulty of the problem")
	addSolutionCmd.Flags().StringP("solution", "s", "", "Solution slug")

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
	// register the flag --replace
	// if the flag is present, the solution file will be replaced
	// if the flag is not present, the solution file will not be replaced
	// default: false
	mustReplace := false
	flagReplace := cmd.Flag("replace")
	if flagReplace != nil {
		mustReplace = flagReplace.Value.String() == "true"
	}

	// register --difficulty flag
	// if the difficulty is not in the list of valid difficulties, return an error
	// if the solution file already exists, return an error
	// if the solution file does not exist, create the solution file
	difficulty := ""
	if difficulty == "" {
		difficultyFlag := cmd.Flag("difficulty")
		if difficultyFlag != nil {
			difficulty = difficultyFlag.Value.String()
		}
	}

	// register the solution slug
	// if the solution slug is empty, return an error
	// if the solution slug is not empty, create the solution file
	solutionSlug := ""
	if solutionSlug == "" {
		solutionFlag := cmd.Flag("solution")
		if solutionFlag != nil {
			solutionSlug = solutionFlag.Value.String()
		}
	}

	if solutionSlug == "" {
		log.Fatal("Invalid solution slug, it cannot be empty")
		return
	}

	// Create the solution file
	if !util.CheckStringInSlice(difficulty, MapSolutionLevel) {
		log.Fatal("Invalid difficulty")
		return
	}

	// lookup the existing solution filename based on the solution slug
	// if it exists, return an error
	if !mustReplace {
		if _, err := os.Stat(RootPath + "/problem" + "/" + difficulty + "/" + solutionSlug + ".go"); err == nil {
			log.Fatal("The solution file already exists")
			return
		}
	}

	// open base stub file
	fileStub, err := os.Open(RootPath + "/problem" + "/" + difficulty + "/base.stub")
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
	solutionFile, err := os.Create(RootPath + "/problem" + "/" + difficulty + "/" + solutionSlug + ".go")
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
