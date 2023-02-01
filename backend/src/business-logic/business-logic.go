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

func (c *customClient) fetchData(params map[string]string, path string, result any) error {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return fmt.Errorf("%w. %s", ErrCouldNotCreateReq, http.StatusText(http.StatusInternalServerError))
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("%w. %s", ErrCouldNotCreateRes, http.StatusText(http.StatusInternalServerError))
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode))
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("%w, %s", ErrCouldNotUnMarshal, http.StatusText(http.StatusInternalServerError))
	}

	return nil
}

func (t *Tournaments) fetchTournaments(date string, client *customClient) error {

	// get tournament info
	var challongeTournaments models.Tournaments

	params := map[string]string{
		"api_key":       client.config.apiKey,
		"state":         "in_progress",
		"created_after": date,
	}

	err := client.fetchData(params, "/tournaments.json", &challongeTournaments)
	if err != nil {
		return err
	}

	// fmt.Printf("%+v, %v\n", challongeTournaments, len(challongeTournaments))

	if len(challongeTournaments) == 0 {
		return nil
	}

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

	return nil
}

func (t *Tournaments) fetchParticipants(client *customClient) error {
	params := map[string]string{
		"api_key": client.config.apiKey,
	}

	for k, v := range t.TournamentInfo {
		if len(v.Participants) == 0 {
			v.Participants = make(map[int]string)
			path := fmt.Sprintf("/tournaments/%v/participants.json", k)
			var participants models.Participants
			err := client.fetchData(params, path, &participants)
			if err != nil {
				return err
			}
			if len(participants) == 0 {
				return nil
			}
			for _, elem := range participants {
				// fmt.Println(elem.Participant.ID, elem.Participant.Name)
				v.Participants[elem.Participant.ID] = elem.Participant.Name
			}
			t.TournamentInfo[k] = v
		}
	}

	return nil
}

func (t *Tournaments) FetchTournaments(date string, client *customClient) error {
	err := t.fetchTournaments(date, client)
	if err != nil {
		return err
	}

	err = t.fetchParticipants(client)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tournaments) FetchMatches(client *customClient) ([]models.TournamentMatches, error) {
	return nil, nil
}
