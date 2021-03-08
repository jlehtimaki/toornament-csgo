package toornament

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)


func GetParticipant(teamName string) (Team, error){
	// First get all teams and get the ID of the team
	var teams []Team
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/participants?name=%s", seasonId,teamName)
	rangeKey := "participants=0-49"

	data, err := toornamentRest(subPath, rangeKey)
	if err != nil {
		return Team{}, err
	}

	// Parse team data
	err = json.Unmarshal(data, &teams)
	if err != nil {
		return Team{}, err
	}

	// Get matches
	teams[0].matches()
	return teams[0], nil
}

func (t *Team) matches(){
	var matches Matches
	subPath := fmt.Sprintf(
		"viewer/v2/tournaments/%s/matches?participant_ids=%s", seasonId, t.Id)
	rangeString := "matches=0-50"
	data, err := toornamentRest(subPath, rangeString)
	if err != nil {
		log.Fatalf("could not get next games: %s", err)
		panic("something went wrong")
	}
	_ = json.Unmarshal(data, &matches)
	t.Matches = matches
}