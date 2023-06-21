package models

type (
	Tournaments []struct {
		Tournament Tournament `json:"tournament"`
	}
	Tournament struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		GameName string `json:"game_name"`
	}

	ResultData struct {
		Data []struct {
			ID         string `json:"id"`
			Attributes struct {
				Name     string `json:"name"`
				GameName string `json:"gameName"`
			}
		} `json:"data"`
	}
)
