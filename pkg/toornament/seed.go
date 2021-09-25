package toornament

import (
	"encoding/json"
	"fmt"
	"github.com/jlehtimaki/toornament-csgo/pkg/structs"
	"sort"
)

func GetSeed() ([]byte, error) {
	seed := map[string][]structs.SeedTeam{}
	//stages, err := getStages()
	//if err != nil { return nil, err }

	standing,_ := GetStandings("3.Div")


	var standings structs.Standings
	err := json.Unmarshal(standing, &standings)
	if err != nil { return nil, err}

	var teams []structs.Division
	for _, s := range standings {
		if s.Position <= 4 {
			teams = append(teams, s)
		}
	}

	f, err := orderTeams(teams)
	if err != nil { return nil, nil}

	seed["2.Div"] = f

	b, err := json.MarshalIndent(seed, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))

	fmt.Println(len(seed["2.Div"]))

	return nil, nil
}

func orderTeams(teams []structs.Division) ([]structs.SeedTeam, error) {
	var seedTeams []structs.SeedTeam
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Position < teams[j].Position
	})

	for n, t := range teams {
		s := structs.SeedTeam{
			Name:      t.Participant.Name,
			PlacementInGroup: t.Position,
			Points:    t.Points,
			PlusMinus: t.Properties.ScoreDifference,
			Wins:      t.Properties.Wins,
			GroupID:   t.GroupID,
		}
		//fmt.Printf("%s - %d  - %d - %d - %d \n", t.Participant.Name, t.Position, t.Points, t.Properties.ScoreDifference, t.Properties.Wins)
		if len(seedTeams) == 0 {
			s.Seed = 1
			seedTeams = append(seedTeams, s)
			continue
		}

		if teams[n-1].Position == s.PlacementInGroup {
			seedTeams = orderTeam(s, seedTeams)
			continue
		}

		seedTeams = append(seedTeams, s)

	}

	return seedTeams, nil
}

func orderTeam(team structs.SeedTeam, seedTeams []structs.SeedTeam) []structs.SeedTeam {
	for n, st := range seedTeams {
		if st.Points == team.Points {
			if st.PlusMinus < team.PlusMinus {
				tmpT := seedTeams[n]
				seedTeams[n] = team
				seedTeams = append(seedTeams, tmpT)
				return seedTeams
			}
		}
		if st.Points < team.Points {
			tmpT := seedTeams[n]
			seedTeams[n] = team
			seedTeams = append(seedTeams, tmpT)
			return seedTeams
		}
	}
	seedTeams = append(seedTeams, team)
	return seedTeams
}