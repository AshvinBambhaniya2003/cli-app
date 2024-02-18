package cmd

import (
	"fmt"
	"netflix/models"
	"netflix/services"

	"github.com/spf13/cobra"
)

var titlefilepath = "CSV/titles.csv"
var creditfilepath = "CSV/credits.csv"

var moviesCountByReleaseYear, moviesCountByAgeCertificate, genresAnalysis, countryAnalysis, seasonAnalysis, titleTypeAnalysis, mostWorkingActor bool

// printMoviesForActorCmd represents the printMoviesForActor command
var AnalyticsCmd = &cobra.Command{
	Use:   "movie-analytics",
	Short: "Perform analytics on a dataset of movies and shows.",
	Long: `The movie-analytics command allows users to analyze a dataset of movies and shows.
	It offers various analytical tasks such as counting movies by release year, age certification, runtime, genres, country, IMDB score, and more.
	Users can specify which analyses they want to perform using command-line options. The results of the analysis are displayed to the user, providing insights into the characteristics and distribution of movies and shows in the dataset.`,
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

		if moviesCountByReleaseYear {
			countRecords := services.ListMoviesCountByReleaseYear(titles)

			fmt.Println("Movies count with the release year:")
			for year, count := range countRecords {
				fmt.Printf("%d: %d\n", year, count)
			}
		}

		if moviesCountByAgeCertificate {
			countRecords := services.ListMoviesCountByAgeCertificate(titles)

			fmt.Println("Movies count with the release year:")
			for year, count := range countRecords {
				fmt.Printf("%s: %d\n", year, count)
			}
		}

		if genresAnalysis {
			genreCount, totalTitles := services.ListMoviesCountPercentageByGenres(titles)
			for genre, count := range genreCount {
				percentage := float64(count) / float64(totalTitles) * 100
				fmt.Printf("%s: %d titles (%.2f%%)\n", genre, count, percentage)
			}
		}

		if countryAnalysis {
			countryCount, totalTitles := services.ListMoviesCountPercentageByCountry(titles)
			for country, count := range countryCount {
				percentage := float64(count) / float64(totalTitles) * 100
				fmt.Printf("%s: %d titles (%.2f%%)\n", country, count, percentage)
			}
		}

		if seasonAnalysis {
			seasonsCounts := services.ListMoviesBySeasonsCount(titles)
			for season, count := range seasonsCounts {
				fmt.Printf("Number of titles with %d sessions: %d\n", season, count)
			}
		}

		if titleTypeAnalysis {
			movieCount, moviePercentage, showCount, showPercentage := services.GetTitleTypeCountsAndPercentages(titles)
			fmt.Printf("Movie count: %d (%.2f%%)\n", movieCount, moviePercentage)
			fmt.Printf("Show count: %d (%.2f%%)\n", showCount, showPercentage)
		}

		if mostWorkingActor {
			mostWorkingActor := services.GetMostWorkingActor(credits)

			fmt.Println("Most working actor:", mostWorkingActor)
		}
	},
}

func init() {
	titleCmd.AddCommand(AnalyticsCmd)
	AnalyticsCmd.Flags().BoolVar(&moviesCountByReleaseYear, "movies-count-by-release-year", false, "Count movies with release year")
	AnalyticsCmd.Flags().BoolVar(&moviesCountByAgeCertificate, "movies-count-by-age-certificates", false, "Count movies by age certificate")
	AnalyticsCmd.Flags().BoolVar(&genresAnalysis, "genres-analysis", false, "Perform genres-wise count and percentage analysis")
	AnalyticsCmd.Flags().BoolVar(&countryAnalysis, "country-analysis", false, "Perform country-wise count and percentage analysis")
	AnalyticsCmd.Flags().BoolVar(&seasonAnalysis, "season-analysis", false, "Perform session count analysis")
	AnalyticsCmd.Flags().BoolVar(&titleTypeAnalysis, "title-type-analysis", false, "Perform movie and show count analysis")
	AnalyticsCmd.Flags().BoolVar(&mostWorkingActor, "most-working-actor", false, "Find the most working actor")
}
