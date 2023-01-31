package businesslogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/MarcBernstein0/match-display/src/models"
)

var (
	ErrCouldNotCreateReq error = errors.New("could not create request")
	ErrCouldNotCreateRes error = errors.New("could not create response")
	ErrCouldNotUnMarshal error = errors.New("could not unmarshal challonge data")
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
		// FetchTournaments fetch all tournaments created after a specific date
		// GET https://api.challonge.com/v1/tournaments.{json|xml}
		FetchTournaments(date string, client *customClient) error

		FetchMatches(client *customClient) ([]models.TournamentMatches, error)
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

func NewTournament() *Tournaments {
	t := Tournaments{
		TournamentInfo: make(map[int]struct {
			Game         string
			Participants map[int]string
		}),
	}

	return &t
}

func NewClient(baseURL, username, apiKey string, client *http.Client) *customClient {
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

func (t *Tournaments) FetchTournaments(date string, client *customClient) error {
	req, err := http.NewRequest(http.MethodGet, client.baseURL+"/tournaments.json", nil)
	if err != nil {
		return fmt.Errorf("%w. %s", ErrCouldNotCreateReq, http.StatusText(http.StatusInternalServerError))
	}
	q := req.URL.Query()
	q.Add("api_key", client.config.apiKey)
	q.Add("state", "in_progress")
	q.Add("created_after", date)
	req.URL.RawQuery = q.Encode()

	res, err := client.client.Do(req)
	if err != nil {
		return fmt.Errorf("%w. %s", ErrCouldNotCreateRes, http.StatusText(http.StatusInternalServerError))
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode))
	}

	var challongeTournaments models.Tournaments
	err = json.NewDecoder(res.Body).Decode(&challongeTournaments)
	if err != nil {
		return fmt.Errorf("%w, %s", ErrCouldNotUnMarshal, http.StatusText(http.StatusInternalServerError))
	}

	if len(challongeTournaments) == 0 {
		return nil
	}

	fmt.Printf("%+v, %v\n", challongeTournaments, len(challongeTournaments))
	// var tournaments Tournaments

	for _, tournament := range challongeTournaments {
		if _, ok := t.TournamentInfo[tournament.Tournament.ID]; !ok {
			t.TournamentInfo[tournament.Tournament.ID] = struct {
				Game         string
				Participants map[int]string
			}{
				Game: tournament.Tournament.GameName,
			}
		}

	}

	fmt.Printf("%+v, %v\n", t, len(t.TournamentInfo))

	return nil
}

func (t *Tournaments) FetchMatches(client *customClient) ([]models.TournamentMatches, error) {
	return nil, nil
}
