package businesslogic

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MarcBernstein0/match-display/src/models"
)

var (
	ErrCouldNotCreateReq error = errors.New("could not create request")
	ErrCouldNotCreateRes error = errors.New("could not create response")
	ErrResponseNotOK     error = errors.New("response not ok")
	ErrServerProblem     error = errors.New("server error")
	ErrNoData            error = errors.New("no data found")
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

		FetchMatches(tournaments *Tournaments) ([]models.TournamentMatches, error)
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
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/tournaments.json", nil)
	if err != nil {
		return nil, fmt.Errorf("%w. %s", ErrCouldNotCreateReq, http.StatusText(http.StatusInternalServerError))
	}
	q := req.URL.Query()
	q.Add("api_key", c.config.apiKey)
	q.Add("state", "in_progress")
	q.Add("created_after", date)
	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w. %s", ErrCouldNotCreateRes, http.StatusText(http.StatusInternalServerError))
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode))
	}

	return &Tournaments{}, nil
}

func (c *customClient) FetchMatches(tournaments *Tournaments) ([]models.TournamentMatches, error) {
	return nil, nil
}
