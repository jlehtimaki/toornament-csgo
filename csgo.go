package CSGO

import (
	"encoding/json"
	"fmt"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	k "github.com/jlehtimaki/toornament-csgo/pkg/kanaliiga"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CSGO(w http.ResponseWriter, r *http.Request) {
	// Payload checks
	var ret []byte
	var err error
	var d struct {
		Type string `json:"type"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Errorf( "wrong type of payload")
		panic("error occured, check logs")
	}
	if d.Type == "" || d.Value == "" {
		log.Errorf("either value or type is empty")
		panic("error occured, check logs")
	}

	if d.Type == "team" {
		ret, err = getTeam(d.Value)
		if err != nil {
			log.Fatal(err)
			panic("error occured while getting team information")
		}
	} else if d.Type == "standings" {
		ret, err = t.GetStandings(d.Value)
		if err != nil {
			log.Fatal(err)
			panic("error occured while getting standings")
		}
	} else if d.Type == "seed" {
		ret, err = t.GetSeed()
		if err != nil {
			log.Fatal(err)
			panic("error occured while getting seeding list")
		}
	}
	_, _ = fmt.Fprintf(w, string(ret))
}

func getTeam(teamName string) ([]byte, error){
	// Get information about the team
	team, err := t.GetTeam(teamName)
	if err != nil {
		return nil, err
	}
	//Loop through players and get their data
	for i, _ := range team.Players {
		err = f.GetData(&team.Players[i])
		if err != nil {
			log.Warnf("%s\n", err)
		}
		err = k.GetData(&team.Players[i])
		if err != nil {
			log.Warnf("%s\n", err)
		}
	}
	ret, err := json.Marshal(team)
	if err != nil {
		return nil, fmt.Errorf( "could not parse return value")
	}
	return ret, nil
}