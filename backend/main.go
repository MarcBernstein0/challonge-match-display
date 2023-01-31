package main

import (
	"fmt"
	"log"
	"os"
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
	// customClient := mainlogic.New(BASE_URL, username, apiKey, http.DefaultClient)
	// r := routes.RouteSetup(customClient)

	// r.Run(":" + port)
}
