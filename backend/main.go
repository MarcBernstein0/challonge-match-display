package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	businesslogic "github.com/MarcBernstein0/match-display/src/business-logic"
	"github.com/MarcBernstein0/match-display/src/routes"
)

const BASE_URL = "https://api.challonge.com/v1"

func main() {
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8080"
	}

	username, present := os.LookupEnv("USER_NAME")
	if !present {
		log.Fatalf("username not provided in env")
	}

	apiKey, present := os.LookupEnv("API_KEY")
	if !present {
		log.Fatalf("api_key not provided in env")
	}

	fmt.Println(port, username, apiKey)
	customClient := businesslogic.NewClient(BASE_URL, username, apiKey, http.DefaultClient)
	var tournamentCache *businesslogic.Tournaments = businesslogic.NewTournament()

	r := routes.RouteSetup(customClient, tournamentCache)

	r.Run(":" + port)
}
