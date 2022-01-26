package structs

type Kanaliiga struct {
	Stats []KanaliigaStat `json:"data"`
}

type KanaliigaStat struct {
	Kills      int `json:"Kills"`
	Deaths     int `json:"Deaths"`
	Assists    int `json:"Assists"`
	MVPs       int `json:"MVPs"`
	KDR        int
	ADR        float64 `json:"ADR"`
	HsPercent  int     `json:"hsPercent"`
	KAST       int     `json:"KAST"`
	KanaRating float64 `json:"kanaRating"`
}

type KanaliigaMatch struct {
	ID           int    `json:"id"`
	Team1        int    `json:"team1"`
	Team2        int    `json:"team2"`
	Team1Score   int    `json:"team1Score"`
	Team1HTScore int    `json:"team1HTScore"`
	Team1OTScore int    `json:"team1OTScore"`
	Team2Score   int    `json:"team2Score"`
	Team2HTScore int    `json:"team2HTScore"`
	Team2OTScore int    `json:"team2OTScore"`
	Map          string `json:"map"`
	LeagueID     int    `json:"leagueID"`
	Demofile     string `json:"demofile"`
	Type         int    `json:"type"`
}

type KanaliigaRanks struct {
	SteamID      string  `json:"steamID"`
	Rank         int     `json:"rank"`
	Level        int     `json:"level"`
	Hours        int     `json:"hours"`
	FaceELO      int     `json:"faceELO"`
	Kanaelo      int     `json:"kanaelo"`
	Fkd          float64 `json:"fkd"`
	Ekd          float64 `json:"ekd"`
	EsportalElo  int     `json:"esportalElo"`
	EsportalRank int     `json:"esportalRank"`
}
