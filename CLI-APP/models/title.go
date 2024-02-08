package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Title struct {
	ID          string
	Title       string
	Type        string
	Description string
	ReleaseYear int
}

type Credit struct {
	PersonID int
	TitleID  string
	Name     string
	Character string
	Role     string
}

func Gettitlewithcount() {
	titles, err := readTitles("titles.csv")
	if err != nil {
		fmt.Println("Error reading titles:", err)
		return
	}

	credits, err := readCredits("credits.csv")
	if err != nil {
		fmt.Println("Error reading credits:", err)
		return
	}

	listAllTitlesWithPersonCount(titles, credits)
}

func SearchTitles(search string) {

	titles, err := readTitles("titles.csv")
	if err != nil {
		fmt.Println("Error reading titles:", err)
		return
	}

	if search == "" {
		fmt.Println(titles)
	}

	// var filteredTitles []Title
	for _, title := range titles {
		if strings.Contains(strings.ToLower(title.Title), strings.ToLower(search)) {
			// filteredTitles = append(filteredTitles, title)
			fmt.Println(title)
		}
	}

}

func readTitles(filename string) ([]Title, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// reader.FieldsPerRecord = -1 // Allow variable number of fields

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var titles []Title
	for _, record := range records {
		title := Title{
			ID:          record[0],
			Title:       record[1],
			Type:        record[2],
			Description: record[3],
			ReleaseYear: parseYear(record[4]),
		}
		titles = append(titles, title)
	}

	return titles, nil
}

func readCredits(filename string) ([]Credit, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var credits []Credit
	for _, record := range records {
		credit := Credit{
			PersonID: parseInt(record[0]),
			TitleID:  record[1],
			Name:     record[2],
			Character: strings.Trim(record[3], `"`),
			Role:     record[4],
		}
		credits = append(credits, credit)
	}

	return credits, nil
}

func parseYear(yearStr string) int {
	var year int
	_, err := fmt.Sscanf(yearStr, "%d", &year)
	if err != nil {
		return 0
	}
	return year
}

func parseInt(str string) int {
	var num int
	_, err := fmt.Sscanf(str, "%d", &num)
	if err != nil {
		return 0
	}
	return num
}

func listAllTitlesWithPersonCount(titles []Title, credits []Credit) {
	fmt.Println("Title,Person(s) Count")
	for _, title := range titles {
		count := countPersonsForTitle(title.ID, credits)
		fmt.Printf("%s,%d\n", title.Title, count)
	}
}

func countPersonsForTitle(titleID string, credits []Credit) int {
	count := 0
	for _, credit := range credits {
		if credit.TitleID == titleID {
			count++
		}
	}
	return count
}

