package services

import (
	"fmt"
	"log"
	"netflix/models"
	"strings"
)

func PrintMoviesForActor(titles []models.Title, credits []models.Credit, actorName string) {
	fmt.Printf("Movies for actor %s:\n", actorName)
	for _, credit := range credits {
		if strings.EqualFold(credit.Name, actorName) {
			title := findMovieByID(titles, credit.TitleID)
			if title != nil {
				fmt.Printf("- %s as %s\n", title.Title, credit.Character)
			} else {
				log.Fatal("No any movies worked by this Actor")
			}
		}
	}
}

func findMovieByID(titles []models.Title, id string) *models.Title {
	for _, title := range titles {
		if title.ID == id && strings.EqualFold(title.Type, "MOVIE") {
			return &title
		}
	}
	return nil
}
