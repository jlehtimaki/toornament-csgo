package main

import (
	"fmt"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
)

func main()  {
	team, err := t.GetParticipant("Polar Squad")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, player := range team.Players {
		player.FaceitRank, player.FaceitElo, player.FaceitUrl, err = f.GetRank(player.Name, player.CustomFields.SteamId)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
	fmt.Println(team)
}

//func getTeamPlayersInfo(team Team) (Team, error){
//	var err error
//	for _, player := range team.Players {
//		player.FaceitRank, player.FaceitElo, player.FaceitUrl, err = faceit.GetRank(player.Name, player.CustomFields.SteamId)
//		if err != nil {
//			return team, nil
//		}
//		fmt.Println(player)
//	}
//	return team, nil
//}