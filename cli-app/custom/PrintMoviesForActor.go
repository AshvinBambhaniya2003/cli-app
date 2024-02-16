package custom

import (
	"netflix/models"
	"fmt"
	"strings"
)

var titlefilepath = "CSV/titles.csv"
var creditfilepath = "CSV/credits.csv"

func ListMoviesForActor(actorName string){
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

	fmt.Printf("Movies for actor %s:\n", actorName)
	for _, credit := range credits {
		if strings.EqualFold(credit.Name, actorName) {
			title := findTitleByID(titles, credit.TitleID)
			fmt.Printf("- %s as %s\n", title.Title, credit.Character)
		}
	}
}

func findTitleByID(titles []models.Title, id string) *models.Title {
	for _, title := range titles {
		if title.ID == id {
			return &title
		}
	}
	return nil
}