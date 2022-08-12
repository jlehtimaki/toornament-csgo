package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/jlehtimaki/toornament-csgo/pkg/csgostats"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	k "github.com/jlehtimaki/toornament-csgo/pkg/kanaliiga"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/team/:id", getTeam)
	router.GET("/standings/:id", t.GetStandings)
	router.GET("/match/next/:id", t.NextMatch)
	router.GET("/match/scheduled/:id", k.GetScheduledMatches)
	router.GET("/seed", t.Seed)
	router.GET("/seed/:id", t.Seed)
	router.GET("/rank/mm/:id", c.GetRank)
	if os.Getenv("GIN_MODE") == "release" {
		err := router.Run(":8080")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		router.SetTrustedProxies([]string{"localhost"})
		router.Run("localhost:8081")
	}
}

func getTeam(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	teamName := c.Param("id")
	// Get information about the team
	team, err := t.GetTeam(teamName)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	//Loop through players and get their data
	for i := range team.Players {
		err = f.GetData(&team.Players[i])
		if err != nil {
			log.Warnf("%s\n", err)
		}
		err = k.GetData(&team.Players[i])
		if err != nil {
			log.Warnf("%s\n", err)
		}
	}

	c.IndentedJSON(http.StatusOK, team)
	return
}

// Unfinished
func bestWorstMap(team *s.Team) {
	type Map struct {
		Wins  int
		Loses int
		WinP  float64
	}
	matches, err := k.GetMatches(team.Name)
	if err != nil {
		log.Error(err)
		return
	}

	id, err := k.GetTeamLeagueID(team.Name)
	if err != nil {
		log.Error(err)
		return
	}

	maps := map[string]Map{}

	for _, m := range matches {
		win := false
		if _, ok := maps[m.Map]; !ok {
			maps[m.Map] = Map{
				Wins:  0,
				Loses: 0,
				WinP:  0,
			}
		}

		if m.Team1 == id {
			if m.Team1OTScore > m.Team2OTScore {
				win = true
			}
			if m.Team1Score > m.Team2Score {
				win = true
			}
		} else {
			if m.Team2Score > m.Team1Score {
				win = true
			}
			if m.Team2OTScore > m.Team1OTScore {
				win = true
			}
		}

		if entry, ok := maps[m.Map]; ok {
			if win {
				entry.Wins = maps[m.Map].Wins + 1
			} else {
				entry.Loses = maps[m.Map].Loses + 1
			}
			if entry.Wins == 0 && entry.Loses != 0 {
				entry.WinP = 0
			} else if entry.Wins != 0 && entry.Loses == 0 {
				entry.WinP = 100
			} else {
				entry.WinP = float64(entry.Wins) / (float64(entry.Wins) + float64(entry.Loses)) * 100
			}
			maps[m.Map] = entry
		}
	}
}
