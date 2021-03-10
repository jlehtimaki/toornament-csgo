package structs

type Kanaliiga struct {
	Stats   []KanaliigaStat `json:"data"`
}

type KanaliigaStat struct {
	Kills             int     `json:"Kills"`
	Deaths            int     `json:"Deaths"`
	Assists           int     `json:"Assists"`
	MVPs              int     `json:"MVPs"`
	KDR				  int
	ADR               float64 `json:"ADR"`
	HsPercent         int     `json:"hsPercent"`
	KAST              int     `json:"KAST"`
	KanaRating        float64 `json:"kanaRating"`
}