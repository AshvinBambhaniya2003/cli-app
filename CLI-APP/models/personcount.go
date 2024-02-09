package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"errors"
	"sort"
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

type TitleCount struct {
	Title string
	Count int
}

var titlefilepath = "CSV/titles.csv"
var creditfilepath = "CSV/credits.csv"

func Gettitlewithcount(skip, limit int, searchQuery, selects, order, orderBy string) {
	titles, err := readTitles(titlefilepath)
	if err != nil {
		fmt.Println("Error reading titles:", err)
		return
	}

	credits, err := readCredits(creditfilepath)
	if err != nil {
		fmt.Println("Error reading credits:", err)
		return
	}

	var searchedTitles []Title
    for _, title := range titles {
        if searchQuery == "" || strings.Contains(strings.ToLower(title.Title), strings.ToLower(searchQuery)) {
            searchedTitles = append(searchedTitles, title)
        }
    }

	paginatedTitles, err := paginateTitles(searchedTitles, skip, limit)
	if err != nil {
		fmt.Println(err)
		return
	}

	listAllTitlesWithPersonCount(paginatedTitles, credits, selects, order, orderBy)
}

func paginateTitles(titles []Title, skip, limit int) ([]Title, error) {
	if skip < 0 {
		skip = 0
	}
	if limit < -1 {
		limit = -1
	}

	end := len(titles)
	
	if skip < 0 || end < 0 || skip >= len(titles) || end > len(titles) {
		return nil, errors.New("invalid pagination indices")
	}

	if limit != -1 && skip+limit < end {
		end = skip + limit
	}	

	return titles[skip:end], nil
}



func readTitles(filename string) ([]Title, error) {
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
	first := 1
	for _, record := range records {
		title := Title{
			ID:          record[0],
			Title:       record[1],
			Type:        record[2],
			Description: record[3],
			ReleaseYear: parseYear(record[4]),
		}
		if (first == 1){
			first++
		}else{
			titles = append(titles, title)
		}
		
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

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var credits []Credit
	first := 1
	for _, record := range records {
		credit := Credit{
			PersonID: parseInt(record[0]),
			TitleID:  record[1],
			Name:     record[2],
			Character: strings.Trim(record[3], `"`),
			Role:     record[4],
		}
		if (first == 1){
			first++
		}else{
			credits = append(credits, credit)
		}
		
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

func listAllTitlesWithPersonCount(titles []Title, credits []Credit, selects, order, orderBy string) {
	
	var titleCounts []TitleCount
	
	for _, title := range titles {
		count := countPersonsForTitle(title.ID, credits)
		title_count := TitleCount{
			Title: title.Title,
			Count: count,
		}
		titleCounts = append(titleCounts, title_count)
	}

	if orderBy != "" {
		switch orderBy {
		case "Title":
			if order == "DSC" {
				sort.Slice(titleCounts, func(i, j int) bool {
					return titleCounts[i].Title > titleCounts[j].Title
				})
			} else {
				sort.Slice(titleCounts, func(i, j int) bool {
					return titleCounts[i].Title < titleCounts[j].Title
				})
			}
		case "Count":
			if order == "DSC" {
				sort.Slice(titleCounts, func(i, j int) bool {
					return titleCounts[i].Count > titleCounts[j].Count
				})
			} else {
				sort.Slice(titleCounts, func(i, j int) bool {
					return titleCounts[i].Count < titleCounts[j].Count
				})
			}
		}
	}


	if selects != "" {
		selectedColumns := strings.Split(selects, ",")
		for _, titleCount := range titleCounts {
			for _, col := range selectedColumns {
				switch col {
				case "Title":
					fmt.Print(titleCount.Title, " ")
				case "Count":
					fmt.Print(titleCount.Count, " ")
				}
			}
			fmt.Println()
		}
	} else {

		for _, titleCount := range titleCounts {
			fmt.Printf("%s,%d\n", titleCount.Title, titleCount.Count)
		}
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

