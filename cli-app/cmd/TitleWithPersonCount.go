package cmd

import (
	"fmt"
	"netflix/models"
	"netflix/services"

	"github.com/spf13/cobra"
)

var skip, limit int
var searchQuery, selects, order, orderBy string

var titlefilepath = "CSV/titles.csv"
var creditfilepath = "CSV/credits.csv"

var countCmd = &cobra.Command{
	Use:   "title-with-personcount",
	Short: "Displays the count of persons associated with each title.",
	Long: `The "personcount" command retrieves the count of persons associated with each title in the database or CSV file. 
	Additionally, the command may include options for sorting and filtering the results based on specific criteria, providing users with flexibility in analyzing person-count data`,
	Run: func(cmd *cobra.Command, args []string) {

		titles, err := models.ReadTitles(titlefilepath)
		if err != nil {
			fmt.Println("Error reading titles:", err)
			return
		}

		credits, err := models.ReadCredits(creditfilepath)
		if err != nil {
			fmt.Println("Error reading credits:", err)
			return
		}

		titleWithPersonCount := services.ListTitleWithPersonCount(titles, credits, searchQuery)

		for _, titleCount := range titleWithPersonCount {
			fmt.Printf("%s,%d\n", titleCount.Title, titleCount.Count)
		}
	},
}

func init() {
	titleCmd.AddCommand(countCmd)
	countCmd.Flags().IntVar(&skip, "skip", 0, "Skip the first N records")
	countCmd.Flags().IntVar(&limit, "limit", -1, "Limit the number of records to M")
	countCmd.Flags().StringVar(&searchQuery, "search", "", "Print only specified columns")
	countCmd.Flags().StringVar(&selects, "selects", "", "Print only specified columns")
	countCmd.Flags().StringVar(&order, "order", "", "Order records by ASC | DSC")
	countCmd.Flags().StringVar(&orderBy, "order-by", "", "Define the column on which order is applied")
}
