package toornament

import (
	"github.com/gin-gonic/gin"
	"github.com/jlehtimaki/toornament-csgo/pkg/structs"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sort"
	"strings"
)

func Seed(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	var data map[string]map[string][]structs.SeedTeam
	var err error
	division := c.Param("id")
	if division == "" {
		data, err = getSeed("")
	} else {
		data, err = getSeed(division)
	}
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, data)
	return
}

// Get the current seed based on rankings in all divisions
func getSeed(div string) (map[string]map[string][]structs.SeedTeam, error) {
	seed := map[string]map[string][]structs.SeedTeam{}
	seed["playoffs"] = map[string][]structs.SeedTeam{}
	seed["pityplayoffs"] = map[string][]structs.SeedTeam{}

	if div != "" {
		// Normal playoffs
		// Set the division and it's seeding to the JSON
		seeding, err := getStandings(div, false)
		if err != nil {
			return nil, err
		}

		seed["playoffs"][div] = seeding

		// Petty playoffs
		// Set the division and it's seeding to the JSON
		seeding, err = getStandings(div, true)
		if err != nil {
			return nil, err
		}

		seed["pityplayoffs"][div] = seeding
	} else {
		// Get stages i.e. Divisions
		stages, err := getStages()
		if err != nil {
			return nil, err
		}

		// Loop through all divisions and get calculate the seedings accordingly
		for _, stage := range stages {
			if strings.Contains(stage.Name, "layoffs") {
				continue
			}

			// Normal playoffs
			seeding, err := getStandings(stage.Name, false)
			if err != nil {
				return nil, err
			}

			// Set the division and it's seeding to the JSON
			seed["playoffs"][stage.Name] = seeding

			// Pitty playoffs
			seeding, err = getStandings(stage.Name, true)
			if err != nil {
				return nil, err
			}

			// Set the division and it's seeding to the JSON
			seed["pityplayoffs"][stage.Name] = seeding
		}
	}
	return seed, nil
}

func getStandings(div string, pity bool) ([]structs.SeedTeam, error) {
	// Get the current standings for group in division
	standings, err := standings(div)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("doing pity: %t \n", pity)
	// Get all teams that are 4th or upper in the group
	var teams []structs.Division
	for _, s := range standings {
		if pity {
			if s.Position > 4 {
				teams = append(teams, s)
			}
		} else {
			if s.Position <= 4 {
				teams = append(teams, s)
			}
		}
	}
	// Order the teams
	seeds, err := orderTeams(teams)
	if err != nil {
		return nil, nil
	}

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
	for n := range seedTeams {
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
