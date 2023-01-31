package businesslogic

import (
	"net/http"
	"time"

	"github.com/MarcBernstein0/match-display/src/models"
)

type (
	// map[tournament id number]struct{
	// 	Game Name,
	// 	Participants map[player id] player tag
	// }
	Tournaments struct {
		TournamentInfo map[int]struct {
			Game         string
			Participants map[int]string
		}
		Timestamp time.Time
	}

	FetchData interface {
		// FetchTournaments fetch all tournaments and tournament participants created after a specific date
		// GET https://api.challonge.com/v1/tournaments.{json|xml}
		FetchTournaments(date string) (*Tournaments, error)

		FetchMatches(tournaments Tournaments) ([]models.TournamentMatches, error)
	}

	customClient struct {
		baseURL string
		client  *http.Client
		config  struct {
			username string
			apiKey   string
		}
	}
)

func New(baseURL, username, apiKey string, client *http.Client) *customClient {
	return &customClient{
		baseURL: baseURL,
		client:  client,
		config: struct {
			username string
			apiKey   string
		}{
			username: username,
			apiKey:   apiKey,
		},
	}
}

func (c *customClient) FetchTournaments(date string) (*Tournaments, error) {
	return nil, nil
}
