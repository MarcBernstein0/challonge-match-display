package models

type (
	Matches []struct {
		Match Match `json:"match"`
	}
	Match struct {
		ID          int    `json:"id"`
		Player1ID   int    `json:"player1_id"`
		Player1Name string `json:"player1_name"`
		Player2ID   int    `json:"player2_id"`
		Player2Name string `json:"player2_name"`
		Round       int    `json:"round"`
		UnderwayAt  string `json:"underway_at"`
	}
	CustomMatch struct {
		ID          int    `json:"id"`
		Player1Name string `json:"player1_name"`
		Player2Name string `json:"player2_name"`
		Round       int    `json:"round"`
		Underway    bool   `json:"underway"`
	}
	TournamentMatches struct {
		GameName     string        `json:"game_name"`
		TournamentID int           `json:"tournament_id"`
		MatchList    []CustomMatch `json:"match_list"`
	}
)
