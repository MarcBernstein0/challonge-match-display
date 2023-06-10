package main

import (
	"log"
	"net/http"
	"os"

	mainlogic "github.com/MarcBernstein0/challonge-match-display/backend/src/main-logic"
	"github.com/MarcBernstein0/challonge-match-display/backend/src/routes"
)

const BASE_URL = "https://api.challonge.com/v1"

func main() {
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8080"
	}

	apiKey, present := os.LookupEnv("API_KEY")
	if !present {
		log.Fatalf("api_key not provided in env")
	}

	customClient := mainlogic.New(BASE_URL, apiKey, http.DefaultClient)
	r := routes.RouteSetup(customClient)

	r.Run(":" + port)
}
