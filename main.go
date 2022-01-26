package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	k "github.com/jlehtimaki/toornament-csgo/pkg/kanaliiga"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/team/:id", getTeam)
	router.GET("/standings/:id", t.GetStandings)
	router.GET("/match/next/:id", t.NextMatch)
	if os.Getenv("GIN_MODE") == "release" {
		err := router.Run(":8080")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		router.SetTrustedProxies([]string{"localhost"})
		router.Run("localhost:8080")
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	// Payload checks
	var ret []byte
	var err error
	var d struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Errorf("wrong type of payload")
		panic("error occured, check logs")
	}
	if d.Type == "" {
		log.Errorf("type is empty")
		panic("error occured, check logs")
	}

	if d.Type == "seed" {
		log.Info("getting seeds")
		if d.Value != "" {
			ret, err = t.GetSeed(d.Value)
		} else {
			ret, err = t.GetSeed("")
		}
		if err != nil {
			log.Fatal(err)
			panic("error occured while getting seeding list")
		}
	}
	_, _ = fmt.Fprintf(w, string(ret))
}

func getTeam(c *gin.Context) {
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
