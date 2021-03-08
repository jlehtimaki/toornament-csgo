package toornament

import "time"

type Team struct {
	Id				string	`json:"id"`
	Name			string	`json:"name"`
	Players			[]Player `json:"lineup"`
	Matches			Matches
	CustomFields	struct{
		CaptainDiscord			string	`json:"kapteenin_discord_nick"`
		ReserveCaptainDiscord	string	`json:"varakapteenin_discord"`
	} `json:"custom_fields"`
}

type Player struct {
	Name			string	`json:"name"`
	CustomFields 	struct{
		SteamId		string	`json:"steam_id_"`
	} `json:"custom_fields"`
	MM				struct{
		Rank		string
	}
	Faceit			struct{
		Id			string
		Rank		int
		Elo			int
		Url			string
		Avatar		string
		MostPlayedMap	Map
		LeastPlayedMap	Map
		Weapon		string
		KD			string
		HSP			string
	}
	Esportal		struct{
		Rank		string
	}
}

type Map struct {
	Name	string
	Icon	string
	WinRate	string
	Matches	int
	KD		string
}

type Matches []Match

type Match struct {
	ID                string    `json:"id"`
	StageID           string    `json:"stage_id"`
	GroupID           string    `json:"group_id"`
	RoundID           string    `json:"round_id"`
	Number            int       `json:"number"`
	Type              string    `json:"type"`
	Status            string    `json:"status"`
	ScheduledDatetime time.Time `json:"scheduled_datetime"`
	PlayedAt          time.Time `json:"played_at"`
	Opponents         []struct {
		Number      int    `json:"number"`
		Position    int    `json:"position"`
		Result      string `json:"result"`
		Rank        int    `json:"rank"`
		Forfeit     bool   `json:"forfeit"`
		Score       int    `json:"score"`
		Participant struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			CustomFields struct {
			} `json:"custom_fields"`
		} `json:"participant"`
	} `json:"opponents"`
}