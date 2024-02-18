/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"netflix/custom"
)

// uniqueActorsCmd represents the uniqueActors command
var uniqueActorsCmd = &cobra.Command{
	Use:   "uniqueactors",
	Short: "Extract and display a list of unique actors from credits data.",
	Long: `uniqueactors command extracts and displays a list of unique actors from the credits data stored in the provided CSV file.
	It reads the credits data, iterates over the actors, and stores them in a map to ensure uniqueness.
	Then, it extracts the unique actor names and prints them out.
	This command is useful for obtaining a comprehensive list of actors involved in the production of titles.`,
	Run: func(cmd *cobra.Command, args []string) {
		custom.ListUniqueActors()
	},
}

func init() {
	titleCmd.AddCommand(uniqueActorsCmd)
}
