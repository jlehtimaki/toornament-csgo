package faceit

type Player struct {
	Nickname	string	`json:"nickname"`
	Steam64		string	`json:"steam_id_64"`
	FaceitUrl	string	`json:"faceit_url"`
	Games		struct{
		CSGO	struct{
			SkillLevel	int	`json:"skill_level"`
			Elo			int	`json:"faceit_elo"`
		} `json:"csgo"`
	} `json:"games"`
}

type SearchResult struct {
	Items	[]struct{
		Nickname	string	`json:"nickname"`
	} `json:"Items"`
}