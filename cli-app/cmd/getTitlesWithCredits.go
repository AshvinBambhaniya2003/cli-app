/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"netflix/custom"
)

// getTitlesWithCreditsCmd represents the getTitlesWithCredits command
var getTitlesWithCreditsCmd = &cobra.Command{
	Use:   "getTitlesWithCredits",
	Short: "Retrieve detailed information about titles along with their associated credits.",
	Long: `The getTitlesWithCredits command allows you to fetch comprehensive details about titles, including their descriptions, release information, and genres, along with credits such as actors and their roles.
	This command provides a comprehensive overview of each title's cast and crew, facilitating a deeper understanding of the content.`,
	Run: func(cmd *cobra.Command, args []string) {
		custom.GetDetailedTitlesWithCredits()
	},
}

func init() {
	titleCmd.AddCommand(getTitlesWithCreditsCmd)
}
