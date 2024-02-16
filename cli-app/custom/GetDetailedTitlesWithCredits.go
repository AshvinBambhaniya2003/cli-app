package custom

import (
	"netflix/models"
	"fmt"
)

var titlefilepath = "CSV/titles.csv"
var creditfilepath = "CSV/credits.csv"


type DetailedTitle struct {
    models.Title
    Credits []models.Credit
}

func GetDetailedTitlesWithCredits(){
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

	detailedTitles := JoinTitlesAndCredits(titles, credits)

	// Display detailed title information
	for _, title := range detailedTitles {
		fmt.Printf("Title: %s\n", title.Title)
		fmt.Printf("Type: %s\n", title.Type)
		fmt.Printf("Description: %s\n", title.Description)
		// Add more fields as needed
		fmt.Println("Credits:")
		for _, credit := range title.Credits {
			fmt.Printf("- %s as %s\n", credit.Name, credit.Character)
		}
		fmt.Println("------------------------------------")
	}

}

func JoinTitlesAndCredits(titles []models.Title, credits []models.Credit) []DetailedTitle {
    // Create a map to store credits by title ID
    creditMap := make(map[string][]models.Credit)
    for _, credit := range credits {
        creditMap[credit.TitleID] = append(creditMap[credit.TitleID], credit)
    }

    // Create a slice to store detailed titles
    var detailedTitles []DetailedTitle

    // Iterate over titles and populate detailedTitles
    for _, title := range titles {
        detailedTitle := DetailedTitle{Title: title, Credits: creditMap[title.ID]}
        detailedTitles = append(detailedTitles, detailedTitle)
    }

    return detailedTitles
}