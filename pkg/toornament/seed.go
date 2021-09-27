package toornament

import (
	"encoding/json"
	"fmt"
	"github.com/jlehtimaki/toornament-csgo/pkg/structs"
	"sort"
	"strings"
)

// Get the current seed based on rankings in all divisions
func GetSeed(div string) ([]byte, error) {
	seed := map[string][]structs.SeedTeam{}

	if div != "" {
		seeding, err := getStandings(div)
		if err != nil {return nil, err}

		// Set the division and it's seeding to the JSON
		seed[div] = seeding
	} else {
		// Get stages i.e. Divisions
		stages, err := getStages()
		if err != nil { return nil, err }

		// Loop through all divisions and get calculate the seedings accordingly
		for _, stage := range stages {
			if strings.Contains(stage.Name, "layoffs") {
				continue
			}

			seeding, err := getStandings(stage.Name)
			if err != nil {return nil, err}

			// Set the division and it's seeding to the JSON
			seed[stage.Name] = seeding
		}
	}
	// Pretify the JSON struct
	b, err := json.MarshalIndent(seed, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	return b, nil
}

func getStandings(div string) ([]structs.SeedTeam, error){
	// Get the current standings for group in division
	data,_ := GetStandings(div)
	var standings structs.Standings
	err := json.Unmarshal(data, &standings)
	if err != nil { return nil, err}

	// Get all teams that are 4th or upper in the group
	var teams []structs.Division
	for _, s := range standings {
		if s.Position <= 4 {
			teams = append(teams, s)
		}
	}

	// Order the teams
	seeds, err := orderTeams(teams)
	if err != nil { return nil, nil}

	return seeds, nil
}

func orderTeams(teams []structs.Division) ([]structs.SeedTeam, error) {
	var seedTeams []structs.SeedTeam
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Position < teams[j].Position
	})

	for _, t := range teams {
		s := structs.SeedTeam{
			Name:             t.Participant.Name,
			PlacementInGroup: t.Position,
			Points:           t.Points,
			PlusMinus:        t.Properties.ScoreDifference,
			Wins:             t.Properties.Wins,
			GroupID:          t.GroupID,
		}

		if len(seedTeams) == 0 {
			s.Seed = 1
			seedTeams = append(seedTeams, s)
			continue
		}
		seedTeams = orderTeam(s, seedTeams)
	}

	// Set correct seed
	for n, _ := range seedTeams {
		seedTeams[n].Seed = n + 1
	}

	return seedTeams, nil
}

func orderTeam(team structs.SeedTeam, seedTeams []structs.SeedTeam) []structs.SeedTeam {
	for n, st := range seedTeams {
		if st.PlacementInGroup == team.PlacementInGroup {
			if st.Points == team.Points {
				if st.PlusMinus < team.PlusMinus {
					return insert(seedTeams, n, team)
				}
				if st.PlusMinus == team.PlusMinus {
					if st.Wins < team.Wins {
						return insert(seedTeams, n, team)
					}
				}
			}
			if st.Points < team.Points {
				return insert(seedTeams, n, team)
			}
		}
	}
	team.Seed = len(seedTeams) + 1
	seedTeams = append(seedTeams, team)
	return seedTeams
}

func insert(s []structs.SeedTeam, index int, team structs.SeedTeam) []structs.SeedTeam {
	if len(s) == index { // nil or empty slice or after last element
		return append(s, team)
	}
	s = append(s[:index+1], s[index:]...) // index < len(a)
	s[index] = team
	return s
}