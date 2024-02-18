/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"netflix/custom"
	"github.com/spf13/cobra"
)

// printMoviesForActorCmd represents the printMoviesForActor command
var printMoviesForActorCmd = &cobra.Command{
	Use:   "printmoviesforactor",
	Short: "print a list of movies for a given actor along with the characters they played.",
	Long: `printmoviesforactor is a command-line application that accepts the name of an actor as input and prints a list of movies in which the actor worked, along with the characters they played.
	It reads the credits data from the provided CSV file, searches for movies involving the specified actor, and outputs the movie titles and corresponding character names.
	This command is useful for obtaining a comprehensive list of movies a particular actor participated in and the roles they portrayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(args) < 2 {
			custom.ListMoviesForActor(args[0])
		}else{
			fmt.Println("only one argument allow")
		}
	}, 
}

func init() {
	titleCmd.AddCommand(printMoviesForActorCmd)
}
