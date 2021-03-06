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

type Match struct {
	Name	string
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