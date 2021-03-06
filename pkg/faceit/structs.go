package faceit

type Player struct {
	Nickname	string	`json:"nickname"`
	Id			string	`json:"player_id"`
	Avatar		string	`json:"avatar"`
	Steam64		string	`json:"steam_id_64"`
	FaceitUrl	string	`json:"faceit_url"`
	Games		struct{
		CSGO	struct{
			SkillLevel	int	`json:"skill_level"`
			Elo			int	`json:"faceit_elo"`
		} `json:"csgo"`
	} `json:"games"`
}

type Stats struct {
	Overall	struct{
		HSP		string	`json:"Average Headshot %"`
		KD		string	`json:"Average K/D Ratio"`
	}`json:"lifetime"`
	Maps		[]Map	`json:"segments"`
}

type Map struct {
	Stats	struct{
		WinRate	string	`json:"Win Rate %"`
		Matches	string	`json:"Matches"`
		KD		string	`json:"K/D Ratio"`
	}`json:"stats"`
	Name	string	`json:"label"`
	Image	string	`json:"img_small"`
	Type	string	`json:"type"`
}