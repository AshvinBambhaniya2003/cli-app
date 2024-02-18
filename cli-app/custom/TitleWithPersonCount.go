package custom

import (
	"fmt"
	"strings"
	"netflix/models"
)

type TitleCount struct {
	Title string
	Count int
}

var titlefilepath = "CSV/titles.csv"
var creditfilepath = "CSV/credits.csv"

func GetTitleWithCount(searchQuery string) {

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

	var searchedTitles []models.Title
    for _, title := range titles {
        if searchQuery == "" || strings.Contains(strings.ToLower(title.Title), strings.ToLower(searchQuery)) {
            searchedTitles = append(searchedTitles, title)
        }
    }

	listAllTitlesWithPersonCount(searchedTitles,credits)
}


func listAllTitlesWithPersonCount(titles []models.Title, credits []models.Credit) {
	
	var titleCounts []TitleCount
	
	for _, title := range titles {
		count := countPersonsForTitle(title.ID, credits)
		title_count := TitleCount{
			Title: title.Title,
			Count: count,
		}
		titleCounts = append(titleCounts, title_count)
	}
	
	for _, titleCount := range titleCounts {
		fmt.Printf("%s,%d\n", titleCount.Title, titleCount.Count)
	}
	

}

func countPersonsForTitle(titleID string, credits []models.Credit) int {
	count := 0
	for _, credit := range credits {
		if credit.TitleID == titleID {
			count++
		}
	}
	return count
}

