package services

import (
	"netflix/models"
	"strings"
)

func ListMoviesCountByReleaseYear(titles []models.Title) map[int]int {
	moviesCountByReleaseYear := make(map[int]int)
	for _, record := range titles {
		if record.Type == "MOVIE" {
			moviesCountByReleaseYear[record.ReleaseYear]++
		}
	}

	return moviesCountByReleaseYear
}

func ListMoviesCountByAgeCertificate(titles []models.Title) map[string]int {
	moviesCountByAgeCertificate := make(map[string]int)
	for _, record := range titles {
		if record.Type == "MOVIE" {
			moviesCountByAgeCertificate[record.AgeCertification]++
		}
	}

	return moviesCountByAgeCertificate
}

func ListMoviesCountPercentageByGenres(titles []models.Title) (map[string]int, int) {

	genreCount := make(map[string]int)

	// Count the occurrences of each genre
	totalTitles := 0
	for _, record := range titles {
		genres := strings.Split(record.Genres[0], ",")
		for _, genre := range genres {
			genre = strings.Trim(genre, "[]'\" ")
			genreCount[genre]++
			totalTitles++
		}
	}

	return genreCount, totalTitles
}

func ListMoviesCountPercentageByCountry(titles []models.Title) (map[string]int, int) {

	countryCount := make(map[string]int)

	// Count the occurrences of each genre
	totalTitles := 0
	for _, record := range titles {
		genres := strings.Split(record.ProductionCountries[0], ",")
		for _, genre := range genres {
			genre = strings.Trim(genre, "[]'\" ")
			countryCount[genre]++
			totalTitles++
		}
	}

	return countryCount, totalTitles
}

func ListMoviesBySeasonsCount(titles []models.Title) map[int]int {

	seasonsCounts := make(map[int]int)

	for _, record := range titles {
		if record.Seasons != 0 {
			seasonsCounts[record.Seasons]++
		}
	}

	return seasonsCounts
}

func GetTitleTypeCountsAndPercentages(titles []models.Title) (int, float64, int, float64) {

	movieCount := 0
	showCount := 0

	// Iterate through each record in the CSV file
	for _, record := range titles {
		// Check the title type (movie or show)
		if record.Type == "MOVIE" {
			movieCount++
		} else if record.Type == "SHOW" {
			showCount++
		}
	}

	// Calculate total count and percentages
	totalCount := movieCount + showCount
	moviePercentage := float64(movieCount) / float64(totalCount) * 100
	showPercentage := float64(showCount) / float64(totalCount) * 100

	return movieCount, moviePercentage, showCount, showPercentage
}
func GetMostWorkingActor(credits []models.Credit) string {
	actorCounts := make(map[string]int)

	// Iterate through each record in the CSV file
	for _, record := range credits {
		actor := record.Name

		// Increment the count for the current actor
		actorCounts[actor]++
	}

	// Find the actor with the highest count
	var mostWorkingActor string
	maxCount := 0
	for actor, count := range actorCounts {
		if count > maxCount {
			maxCount = count
			mostWorkingActor = actor
		}
	}

	return mostWorkingActor
}
