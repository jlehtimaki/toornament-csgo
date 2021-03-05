package toornament

type Team struct {
	Id				string	`json:"id"`
	Name			string	`json:"name"`
	Players			[]Player `json:"lineup"`
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
	MMRank			string
	FaceitRank		int
	FaceitElo		int
	FaceitUrl		string
	EsportalRank	string
}