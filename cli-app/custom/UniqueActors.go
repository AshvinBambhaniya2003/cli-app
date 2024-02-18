package custom

import (
    "fmt"
    "netflix/models" 
)

var creditfilepath = "CSV/credits.csv"

func ListUniqueActors() {
    // Read credits from CSV
    credits, err := models.ReadCredits(creditfilepath)
    if err != nil {
        fmt.Println("Error reading credits:", err)
        return
    }

    // Create a map to store unique actors
    uniqueActors := make(map[string]bool)

    // Iterate over credits and add actors to the map

    for _, credit := range credits {
		if credit.Role == "ACTOR"{
			uniqueActors[credit.Name] = true
		}
		
    }


    // Extract unique actors from the map
    var actors []string
    for actor := range uniqueActors {
        actors = append(actors, actor)
    }

    // Print the list of unique actors
    fmt.Println("List of Unique Actors:")
    for _, actor := range actors {
        fmt.Println(actor)
    }
}
