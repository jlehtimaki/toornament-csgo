package CSGO

import (
	"encoding/json"
	"fmt"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CSGO(w http.ResponseWriter, r *http.Request) {
	// Payload checks
	var d struct {
		TeamName string `json:"team_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Errorf( "wrong type of payload")
		panic("error occured")
	}
	if d.TeamName == "" {
		log.Errorf("could not parse payload")
		panic("error occured")
	}

	// Get information about the team
	team, err := t.GetParticipant(d.TeamName)
	if err != nil {
		log.Errorf("Error: %s", err)
		panic("could not get participant")
	}
	// Loop through players and get their data
	for i, _ := range team.Players {
		err = f.GetData(&team.Players[i])
		if err != nil {
			log.Warnf("%s\n", err)
		}
	}
	ret, err := json.Marshal(team)
	if err != nil {
		log.Errorf( "could not parse return value")
		panic("could not parse data")
	}
	_, _ = fmt.Fprintf(w, string(ret))
}