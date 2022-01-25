package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	k "github.com/jlehtimaki/toornament-csgo/pkg/kanaliiga"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/team/:id", getTeam)
	router.GET("/standings/:id", t.GetStandings)
	router.GET("/match/next/:id", t.NextMatch)
	router.Run("localhost:8080")
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
}
