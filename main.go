package main

import (
	"github.com/spf13/cobra"
	leetcode_cmd "github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode/cmd"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use:   "app",
			Short: "app is a CLI tool to manage ressources",
			Long:  "app is a CLI tool to manage ressources",
		}
	)

	leetcodeCmd := leetcode_cmd.NewCMD()

	rootCmd.AddCommand(leetcodeCmd)
	rootCmd.Execute()
}
