/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	_ "fmt"
	_ "log"
	"github.com/spf13/cobra"

	"netflix/models"
)

// titlesCmd represents the titles command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		models.Gettitlewithcount()
		// count, err := cmd.Flags().GetBool("count")

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// if count {
			
		// }

		// search, err := cmd.Flags().GetString("search")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// if search != "" {
		// 	models.SearchTitles(search)
		// }


	},
}

func init() {
	titleCmd.AddCommand(countCmd)
	// titlesCmd.Flags().BoolP("count", "c", false, "Count of the Person in the specific titles")
	// titlesCmd.Flags().String("search", "", "get the list based on the title")
}
