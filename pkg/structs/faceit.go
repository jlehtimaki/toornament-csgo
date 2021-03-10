package structs

type FaceitPlayer struct {
	Nickname  string `json:"nickname"`
	Id        string `json:"player_id"`
	Avatar    string `json:"avatar"`
	Steam64   string `json:"steam_id_64"`
	FaceitUrl string `json:"faceit_url"`
	Games     struct {
		CSGO struct {
			SkillLevel int `json:"skill_level"`
			Elo        int `json:"faceit_elo"`
		} `json:"csgo"`
	} `json:"games"`
}

type Stats struct {
	Overall struct {
		AverageKDRatio   string `json:"Average K/D Ratio"`
		TotalHeadshots   string `json:"Total Headshots %"`
		LongestWinStreak string `json:"Longest Win Streak"`
		Wins             string `json:"Wins"`
		Matches          string `json:"Matches"`
		AverageHeadshots string `json:"Average Headshots %"`
		CurrentWinStreak string `json:"Current Win Streak"`
		KDRatio          string `json:"K/D Ratio"`
		WinRate          string `json:"Win Rate %"`
	} `json:"lifetime"`
	Maps []FaceitMap `json:"segments"`
}

type FaceitMap struct {
	Stats struct {
		WinRate string `json:"Win Rate %"`
		Matches string `json:"Matches"`
		KD      string `json:"K/D Ratio"`
	} `json:"stats"`
	Name  string `json:"label"`
	Image string `json:"img_small"`
	Type  string `json:"type"`
	Mode  string `json:"mode"`
}
