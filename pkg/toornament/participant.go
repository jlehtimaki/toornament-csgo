package toornament

import (
	"encoding/json"
	"fmt"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	log "github.com/sirupsen/logrus"
	"strings"
)


func GetParticipant(teamName string) (s.Team, error){
	// First get all teams and get the ID of the team
	var teams []s.Team
	teamName = strings.ReplaceAll(teamName, " ", "+")
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/participants?name=%s", seasonId,teamName)
	rangeKey := "participants=0-49"

	data, err := toornamentRest(subPath, rangeKey)
	if err != nil {
		return s.Team{}, err
	}

	// Parse team data
	err = json.Unmarshal(data, &teams)
	if err != nil {
		return s.Team{}, err
	}

	// Get matches
	matches(&teams[0])
	return teams[0], nil
}

func matches(team *s.Team){
	var matches s.Matches
	subPath := fmt.Sprintf(
		"viewer/v2/tournaments/%s/matches?participant_ids=%s", seasonId, team.Id)
	rangeString := "matches=0-50"
	data, err := toornamentRest(subPath, rangeString)
	if err != nil {
		log.Fatalf("could not get next games: %s", err)
		panic("something went wrong")
	}
	_ = json.Unmarshal(data, &matches)
	team.Matches = matches
}