package services

import (
	"netflix/models"
)

type DetailedTitle struct {
	models.Title
	Credits []models.Credit
}

func ListTitlesWithCredits(titles []models.Title, credits []models.Credit) []DetailedTitle {

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
