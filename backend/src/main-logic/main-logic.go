package mainlogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"sync"

	"github.com/MarcBernstein0/challonge-match-display/backend/src/models"
)

var (
	ErrResponseNotOK error = errors.New("response not ok")
	ErrServerProblem error = errors.New("server error")
	ErrNoData        error = errors.New("no data found")
)

type (
	participantResult struct {
		tournamentParticipant *models.TournamentParticipants
		error                 error
	}
	matchResult struct {
		tournamentMatches *models.TournamentMatches
		error             error
	}
	FetchData interface {
		// FetchTournaments fetch all tournaments created after a specific date
		// GET https://api.challonge.com/v1/tournaments.{json|xml}
		FetchTournaments(date string) ([]models.Tournament, error)
		// FetchParticipants of a given tournament
		// GET https://api.challonge.com/v1/tournaments/{tournament}/participants.{json|xml}
		FetchParticipants(tournaments []models.Tournament) ([]models.TournamentParticipants, error)
		// FetchMatches of a given tournament
		// GET https://api.challonge.com/v1/tournaments/{tournament}/matches.{json|xml}
		FetchMatches(tournamentParticipants []models.TournamentParticipants) ([]models.TournamentMatches, error)
	}

	customClient struct {
		baseURL string
		client  *http.Client
		config  struct {
			apiKey string
		}
	}
)

func New(baseURL, apiKey string, client *http.Client) *customClient {
	return &customClient{
		baseURL: baseURL,
		client:  client,
		config: struct {
			apiKey string
		}{
			apiKey: apiKey,
		},
	}
}

func (c *customClient) FetchTournaments(date string) ([]models.Tournament, error) {
	// ctx, cancel := context.WithTimeout(ctx, c.timeout)
	// defer cancel()

	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	// if err != nil {
	// 	return nil, err
	// }
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/tournaments.json", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("api_key", c.config.apiKey)
	q.Add("state", "in_progress")
	// fmt.Println(date)
	q.Add("created_after", date)
	req.URL.RawQuery = q.Encode()
	// fmt.Printf("%v\n", req.URL)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode))
	}

	var tournaments models.Tournaments
	err = json.NewDecoder(res.Body).Decode(&tournaments)
	if err != nil {
		return nil, fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusInternalServerError))
	}
	// fmt.Printf("%+v, %v\n", tournaments, len(tournaments))

	if len(tournaments) == 0 {
		return []models.Tournament{}, nil
	}
	var tournamentList []models.Tournament
	for _, t := range tournaments {
		tournamentList = append(tournamentList, t.Tournament)
	}
	return tournamentList, err
}

func (c *customClient) fetchAllParticipants(tournament models.Tournament, participantResultChan chan<- participantResult, wg *sync.WaitGroup) {
	defer wg.Done()
	tournamentID := tournament.ID
	gameName := tournament.GameName
	url := fmt.Sprintf("%s/tournaments/%v/participants.json", c.baseURL, tournamentID)
	// fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// return nil, err
		participantResultChan <- participantResult{
			tournamentParticipant: nil,
			error:                 err,
		}
		return
	}
	q := req.URL.Query()
	q.Add("api_key", c.config.apiKey)
	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		// return nil, err
		participantResultChan <- participantResult{
			tournamentParticipant: nil,
			error:                 err,
		}
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		// return nil, fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode))
		participantResultChan <- participantResult{
			tournamentParticipant: nil,
			error:                 fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode)),
		}
		return
	}

	var participants models.Participants
	err = json.NewDecoder(res.Body).Decode(&participants)
	if len(participants) == 0 {
		// return nil, fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusNotFound))
		participantResultChan <- participantResult{
			tournamentParticipant: nil,
			error:                 fmt.Errorf("%w. %s", ErrNoData, http.StatusText(http.StatusNotFound)),
		}
		return
	}
	if err != nil {
		// return nil, fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusInternalServerError))
		participantResultChan <- participantResult{
			tournamentParticipant: nil,
			error:                 fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusInternalServerError)),
		}
		return
	}
	// fmt.Printf("%+v, %v\n", participants, len(participants))

	tournamentParticipant := models.TournamentParticipants{
		GameName:     gameName,
		TournamentID: tournamentID,
		Participant:  map[int]string{},
	}
	for _, p := range participants {
		tournamentParticipant.Participant[p.Participant.ID] = p.Participant.Name
	}

	participantResultChan <- participantResult{
		tournamentParticipant: &tournamentParticipant,
		error:                 nil,
	}

}

func (c *customClient) FetchParticipants(tournaments []models.Tournament) ([]models.TournamentParticipants, error) {
	var tournamentParticipants []models.TournamentParticipants

	cResponse := make(chan participantResult)
	var wg sync.WaitGroup
	for _, tournament := range tournaments {
		wg.Add(1) // add one to the waitgroup
		go c.fetchAllParticipants(tournament, cResponse, &wg)
	}

	go func() {
		wg.Wait()
		close(cResponse)
	}()

	for tournamentParticipantResult := range cResponse {
		if tournamentParticipantResult.error != nil {
			return nil, tournamentParticipantResult.error
		}
		tournamentParticipants = append(tournamentParticipants, *tournamentParticipantResult.tournamentParticipant)

	}

	// fmt.Printf("Final game participants: %+v", tournamentParticipants)
	return tournamentParticipants, nil
}

func (c *customClient) fetchAllMatches(tournamentParticipant models.TournamentParticipants, matchResultChan chan<- matchResult, wg *sync.WaitGroup) {
	defer wg.Done()
	tournamentID := tournamentParticipant.TournamentID
	gameName := tournamentParticipant.GameName
	participants := tournamentParticipant.Participant

	// fmt.Println(tournamentID, gameName, participants)

	url := fmt.Sprintf("%v/tournaments/%v/matches.json", c.baseURL, tournamentID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		matchResultChan <- matchResult{
			tournamentMatches: nil,
			error:             err,
		}
	}
	q := req.URL.Query()
	q.Add("api_key", c.config.apiKey)
	q.Add("state", "open")
	// fmt.Println(date)
	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		matchResultChan <- matchResult{
			tournamentMatches: nil,
			error:             err,
		}
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		matchResultChan <- matchResult{
			tournamentMatches: nil,
			error:             fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(res.StatusCode)),
		}
	}

	var matches models.Matches
	err = json.NewDecoder(res.Body).Decode(&matches)
	if err != nil {
		matchResultChan <- matchResult{
			tournamentMatches: nil,
			error:             fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusInternalServerError)),
		}
		// return nil, fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusInternalServerError))
	}
	if len(matches) == 0 {
		matchResultChan <- matchResult{
			tournamentMatches: nil,
			error:             nil,
		}
		// return nil, fmt.Errorf("%w. %s", ErrServerProblem, http.StatusText(http.StatusNotFound))
	}
	tournamentMatches := models.TournamentMatches{
		GameName:     gameName,
		TournamentID: tournamentID,
		MatchList:    make([]models.CustomMatch, 0),
	}
	for _, m := range matches {
		customMatch := models.CustomMatch{
			ID:          m.Match.ID,
			Player1Name: participants[m.Match.Player1ID],
			Player2Name: participants[m.Match.Player2ID],
			Round:       m.Match.Round,
			Underway:    len(m.Match.UnderwayAt) != 0,
		}
		fmt.Println("customMatch:", customMatch)
		tournamentMatches.MatchList = append(tournamentMatches.MatchList, customMatch)
	}
	// fmt.Printf("%+v\n", tournamentMatches)

	matchResultChan <- matchResult{
		tournamentMatches: &tournamentMatches,
		error:             nil,
	}
}

func (c *customClient) FetchMatches(tournamentParticipants []models.TournamentParticipants) ([]models.TournamentMatches, error) {

	var tournamentMatches []models.TournamentMatches

	cResponse := make(chan matchResult)
	var wg sync.WaitGroup
	for _, tournamentParticipant := range tournamentParticipants {
		wg.Add(1) // add one to the waitgroup
		go c.fetchAllMatches(tournamentParticipant, cResponse, &wg)
	}

	go func() {
		wg.Wait()
		close(cResponse)
	}()

	for tournamentMatchResult := range cResponse {
		// fmt.Printf("%+v", tournamentMatchResult.tournamentMatches)
		if tournamentMatchResult.error != nil {
			return nil, tournamentMatchResult.error
		}
		if tournamentMatchResult.tournamentMatches != nil {
			sort.SliceStable(tournamentMatchResult.tournamentMatches.MatchList, func(i, j int) bool {
				matchList := tournamentMatchResult.tournamentMatches.MatchList
				// sort the same name if round matches
				match1 := math.Abs(float64(matchList[i].Round))
				match2 := math.Abs(float64(matchList[j].Round))
				if match1 == match2 {
					return matchList[i].Player1Name <= matchList[j].Player1Name
				}
				return match1 < match2
			})
			tournamentMatches = append(tournamentMatches, *tournamentMatchResult.tournamentMatches)
		}

	}

	// fmt.Printf("All matches: %+v\n", tournamentMatches)
	return tournamentMatches, nil
}
