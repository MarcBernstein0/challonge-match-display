package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MarcBernstein0/challonge-match-display/backend/src/business-logic/cache"
	mainlogic "github.com/MarcBernstein0/challonge-match-display/backend/src/business-logic/main-logic"
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

	cacheTimerString, present := os.LookupEnv("CACHE_TIMER")
	if !present {
		cacheTimerString = "5"
	}
	cacheTimer, err := strconv.Atoi(cacheTimerString)
	if err != nil {
		log.Fatalf("cacheTimer could not be read properly\n%s", err)
	}

	cache := cache.NewCache(time.Duration(cacheTimer) * time.Minute)
	customClient := mainlogic.New(BASE_URL, apiKey, http.DefaultClient)
	r := routes.RouteSetup(customClient, cache)

	r.Run(":" + port)
}
