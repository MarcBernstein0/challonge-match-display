package businesslogic

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/MarcBernstein0/match-display/src/models"
	"github.com/stretchr/testify/assert"
)

var server *httptest.Server

const (
	MOCK_API_KEY        = "mock_api_key"
	MOCK_API_USERNAME   = "mock_api_username"
	NO_TOURNAMENTS_DATE = "2025-02-03"
)

func testApiAuth(apiKey string) bool {
	if len(apiKey) == 0 {
		return false
	} else if apiKey != MOCK_API_KEY {
		return false
	}
	return true
}

func readJsonFile(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	return byteValue, err
}

func mockFetchTournamentEndpoint(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("api_key")
	if !testApiAuth(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sc := http.StatusOK
	w.WriteHeader(sc)

	date := r.URL.Query().Get("created_after")
	if date == NO_TOURNAMENTS_DATE {
		w.Write([]byte("[]"))
		return
	}

	byteValue, err := readJsonFile("./test-data/testTournamentData.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// fmt.Println(string(byteValue))

	w.Write(byteValue)
}

func mockFetchParticipantEndpoint(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("api_key")
	if !testApiAuth(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sc := http.StatusOK
	w.WriteHeader(sc)

	byteValue, err := readJsonFile("./test-data/testParticipantData.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// fmt.Println(string(byteValue))

	w.Write(byteValue)
}

func mockFetchParticipantEndpoint2(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("api_key")
	if !testApiAuth(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sc := http.StatusOK
	w.WriteHeader(sc)

	jsonFile, err := os.Open("./test-data/testParticipantData.json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	// fmt.Println(string(byteValue))

	w.Write(byteValue)
}

func mockFetchMatchesEndpoint(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("api_key")
	if !testApiAuth(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sc := http.StatusOK
	w.WriteHeader(sc)

	jsonFile, err := os.Open("./test-data/testMatchesData.json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	// fmt.Println(string(byteValue))

	w.Write(byteValue)
}

func mockFetchMatchesEndpoint2(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("api_key")
	if !testApiAuth(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sc := http.StatusOK
	w.WriteHeader(sc)

	jsonFile, err := os.Open("./test-data/testMatchesData2.json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	// fmt.Println(string(byteValue))

	w.Write(byteValue)
}

func TestMain(m *testing.M) {
	fmt.Println("Starting Mock Server")
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// mock calls go here
		switch strings.TrimSpace(r.URL.Path) {
		case "/tournaments.json":
			mockFetchTournamentEndpoint(w, r)
		case "/tournaments/10879090/participants.json":
			mockFetchParticipantEndpoint(w, r)
		case "/tournaments/10879091/participants.json":
			mockFetchParticipantEndpoint2(w, r)
		case "/tournaments/10879090/matches.json":
			mockFetchMatchesEndpoint(w, r)
		case "/tournaments/10879091/matches.json":
			mockFetchMatchesEndpoint2(w, r)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))

	fmt.Println("mocking CustomClient")

	fmt.Println("run tests")
	m.Run()
}

func TestCustomClient_FetchTournaments(t *testing.T) {
	tt := []struct {
		name      string
		date      string
		fetchData *Tournaments
		client    *CustomClient
		wantData  *Tournaments
		wantErr   error
	}{
		{
			name:      "response not ok",
			date:      time.Now().Local().Format("2006-01-02"),
			fetchData: NewTournament(),
			client:    NewClient(server.URL, "ashdfhsf", "asdfhdsfh", http.DefaultClient),
			wantData:  NewTournament(),
			wantErr:   fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(http.StatusUnauthorized)),
		},
		{
			name:      "data found",
			date:      time.Now().Local().Format("2006-01-02"),
			fetchData: NewTournament(),
			client:    NewClient(server.URL, MOCK_API_USERNAME, MOCK_API_KEY, http.DefaultClient),
			wantData: &Tournaments{
				TournamentInfo: map[int]struct {
					Game         string
					Participants map[int]string
				}{
					10879090: {
						Game: "Guilty Gear -Strive-",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
					10879091: {
						Game: "DNF Duel",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
				},
			},
		},
		{
			name:      "no data found but response ok",
			date:      NO_TOURNAMENTS_DATE,
			fetchData: NewTournament(),
			client:    NewClient(server.URL, MOCK_API_USERNAME, MOCK_API_KEY, http.DefaultClient),
			wantData:  NewTournament(),
			wantErr:   nil,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {

			gotErr := testCase.fetchData.FetchTournaments(testCase.date, testCase.client)
			if testCase.wantErr != nil {
				assert.Equal(t, testCase.wantData.TournamentInfo, testCase.fetchData.TournamentInfo)
				assert.EqualError(t, gotErr, testCase.wantErr.Error())
			} else {
				if assert.NotNil(t, testCase.fetchData.TournamentInfo) {
					assert.Equal(t, testCase.wantData.TournamentInfo, testCase.fetchData.TournamentInfo)
				}
				assert.NoError(t, gotErr)
			}
		})
	}
}

func TestCustomClient_FetchMatches(t *testing.T) {
	tt := []struct {
		name      string
		date      string
		fetchData *Tournaments
		client    *CustomClient
		wantData  []models.TournamentMatches
		wantErr   error
	}{
		{
			name: "response not ok",
			fetchData: &Tournaments{
				TournamentInfo: map[int]struct {
					Game         string
					Participants map[int]string
				}{
					10879090: {
						Game: "Guilty Gear -Strive-",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
				},
			},
			client:   NewClient(server.URL, "asdfhjk", "asdfghj", http.DefaultClient),
			wantData: nil,
			wantErr:  fmt.Errorf("%w. %s", ErrResponseNotOK, http.StatusText(http.StatusUnauthorized)),
		},
		{
			name: "found matches",
			fetchData: &Tournaments{
				TournamentInfo: map[int]struct {
					Game         string
					Participants map[int]string
				}{
					10879090: {
						Game: "Guilty Gear -Strive-",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
					10879091: {
						Game: "DNF Duel",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
				},
			},
			client: NewClient(server.URL, MOCK_API_USERNAME, MOCK_API_KEY, http.DefaultClient),
			wantData: []models.TournamentMatches{
				{
					GameName:     "Guilty Gear -Strive-",
					TournamentID: 10879090,
					MatchList: []models.Match{
						{
							ID:          267800918,
							Player1ID:   166014671,
							Player1Name: "test",
							Player2ID:   166014674,
							Player2Name: "test4",
							Round:       1,
						},
						{
							ID:          267800919,
							Player1ID:   166014672,
							Player1Name: "test2",
							Player2ID:   166014673,
							Player2Name: "test3",
							Round:       1,
						},
					},
				},
				{
					GameName:     "DNF Duel",
					TournamentID: 10879091,
					MatchList: []models.Match{
						{
							ID:          267800918,
							Player1ID:   166014671,
							Player1Name: "test",
							Player2ID:   166014674,
							Player2Name: "test4",
							Round:       1,
						},
						{
							ID:          267800919,
							Player1ID:   166014672,
							Player1Name: "test2",
							Player2ID:   166014673,
							Player2Name: "test3",
							Round:       1,
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			gotData, gotErr := testCase.fetchData.FetchMatches(testCase.client)
			assert.ElementsMatch(t, testCase.wantData, gotData)
			if testCase.wantErr != nil {
				assert.EqualError(t, gotErr, testCase.wantErr.Error())
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}

func TestTournamentCacheUpdate(t *testing.T) {
	tt := []struct {
		name      string
		date      string
		fetchData *Tournaments
		client    *CustomClient
		wantData  *Tournaments
		wantErr   error
	}{
		{
			name: "Reset Data",
			date: NO_TOURNAMENTS_DATE,
			fetchData: &Tournaments{
				TournamentInfo: map[int]struct {
					Game         string
					Participants map[int]string
				}{
					10879090: {
						Game: "Guilty Gear -Strive-",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
					10879091: {
						Game: "DNF Duel",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
				},
				Timestamp: time.Now().Add(-1 * (30 * time.Hour)),
			},
			wantData: NewTournament(),
			client:   NewClient(server.URL, MOCK_API_USERNAME, MOCK_API_KEY, http.DefaultClient),
			wantErr:  nil,
		},
		{
			name: "call for update but no update happens",
			date: time.Now().Local().Format("2006-01-02"),
			fetchData: &Tournaments{
				TournamentInfo: map[int]struct {
					Game         string
					Participants map[int]string
				}{
					10879090: {
						Game: "Guilty Gear -Strive-",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
					10879091: {
						Game: "DNF Duel",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
				},
				Timestamp: time.Now().Add(-1 * (30 * time.Minute)),
			},
			wantData: &Tournaments{
				TournamentInfo: map[int]struct {
					Game         string
					Participants map[int]string
				}{
					10879090: {
						Game: "Guilty Gear -Strive-",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
					10879091: {
						Game: "DNF Duel",
						Participants: map[int]string{
							166014671: "test",
							166014672: "test2",
							166014673: "test3",
							166014674: "test4",
						},
					},
				},
			},
			client:  NewClient(server.URL, MOCK_API_USERNAME, MOCK_API_KEY, http.DefaultClient),
			wantErr: nil,
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			gotErr := testCase.fetchData.UpdateTournamentCache(testCase.date, testCase.client)
			if testCase.wantErr != nil {
				assert.Equal(t, testCase.wantData.TournamentInfo, testCase.fetchData.TournamentInfo)
				assert.EqualError(t, gotErr, testCase.wantErr.Error())
			} else {
				if assert.NotNil(t, testCase.fetchData.TournamentInfo) {
					assert.Equal(t, testCase.wantData.TournamentInfo, testCase.fetchData.TournamentInfo)
				}
				assert.NoError(t, gotErr)
			}
		})
	}
}
