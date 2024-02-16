package models

import (
	"encoding/csv"
	_ "fmt"
	"os"
	"strconv"
)


type Title struct {
    ID                string   
    Title             string   
    Type              string   
    Description       string   
    ReleaseYear       int      
    AgeCertification  string   
    Runtime           int      
    Genres            []string 
    ProductionCountries []string 
    Seasons           int      
    IMDbID            string   
    IMDbScore         float64  
    IMDbVotes         int      
    TmdbPopularity    float64  
    TmdbScore         float64  
}


func ReadTitles(filename string) ([]Title, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var titles []Title
	first := true

	for _, record := range records {
		// Skip the header row
		if first {
			first = false
			continue
		}

		title := Title{
			ID:                record[0],
			Title:             record[1],
			Type:              record[2],
			Description:       record[3],
			ReleaseYear:       parseYear(record[4]),
			AgeCertification:  record[5],
			Runtime:           parseYear(record[6]),
			Genres:            []string{record[7]},
			ProductionCountries: []string{record[8]},
			Seasons:           parseYear(record[9]),
			IMDbID:            record[10],
			IMDbScore:         parseFloat(record[11]),
			IMDbVotes:         parseInt(record[12]),
			TmdbPopularity:    parseFloat(record[13]),
			TmdbScore:         parseFloat(record[14]),
		}

		titles = append(titles, title)
	}

	return titles, nil
}

func parseYear(yearStr string) int {
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0
	}
	return year
}

func parseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return result
}

func parseInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}