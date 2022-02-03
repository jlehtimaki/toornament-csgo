package kanaliiga

func rankConverter(level int, league string) string {
	mmRanks := map[int]string{
		1:  "Silver I",
		2:  "Silver II",
		3:  "Silver III",
		4:  "Silver IV",
		5:  "Silver Elite",
		6:  "Silver Elite Master",
		7:  "Gold Nova I",
		8:  "Gold Nova II",
		9:  "Gold Nova III",
		10: "Gold Nova Master",
		11: "Master Guardian I",
		12: "Master Guardian II",
		13: "Master Guardian Elite",
		14: "Distinguished Master Guardian",
		15: "Legendary Eagle",
		16: "Legendary Eagle Master",
		17: "Supreme Master First Class",
		18: "Global Elite",
	}
	esportalRanks := map[int]string{
		1:  "Silver",
		2:  "Gold I",
		3:  "Gold II",
		4:  "Veteran I",
		5:  "Veteran II",
		6:  "Master I",
		7:  "Master II",
		8:  "Elite I",
		9:  "Elite II",
		10: "Pro I",
		11: "Pro II",
		12: "Legend",
	}
	if league == "MM" {
		return mmRanks[level]
	} else if league == "Esportal" {
		return esportalRanks[level]
	}
	return ""
}
