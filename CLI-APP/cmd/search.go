/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	_ "fmt"
	"log"
	"github.com/spf13/cobra"

	"netflix/models"
)

// titlesCmd represents the titles command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if(len(args) < 1){
			log.Fatal("pass only one arguments please")
		}else if len(args) >1 {
			log.Fatal("pass only one arguments please")
		}else{
			models.SearchTitles(args[0])
		}
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
	titleCmd.AddCommand(searchCmd)
	// titlesCmd.Flags().BoolP("count", "c", false, "Count of the Person in the specific titles")
	// titlesCmd.Flags().String("search", "", "get the list based on the title")
}
