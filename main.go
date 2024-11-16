package main

import (
	"github.com/spf13/cobra"
	leetcode_cmd "github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode/cmd"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use:   "go-ressources",
			Short: "go-ressources is a CLI tool to manage ressources",
			Long:  "go-ressources is a CLI tool to manage ressources",
		}
	)

	leetcodeCmd := leetcode_cmd.NewCMD()

	rootCmd.AddCommand(leetcodeCmd)
	rootCmd.Execute()
}
