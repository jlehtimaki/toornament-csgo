package toornament

import (
	"encoding/json"
	"fmt"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	log "github.com/sirupsen/logrus"
)

func GetTeam(teamName string) (s.Team, error) {
	// First get all teams and get the ID of the team
	var team s.Team
	var rangeKey string

	teamID, err := getTeamID(teamName, seasonId)
	if err != nil {
		return s.Team{}, err
	}
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/participants/%s", seasonId, teamID)

	data, err := toornamentRest(subPath, rangeKey)
	if err != nil {
		return s.Team{}, err
	}

	err = json.Unmarshal(data, &team)
	if err != nil {
		return s.Team{}, err
	}

	// Get matches
	matches(&team)
	return team, nil
}

func getTeamID(teamName string, tournamentID string) (string, error) {
	type Participant struct {
		ID           string `json:"id"`
		TournamentID string `json:"tournament_id"`
		Name         string `json:"name"`
	}
	var participants []Participant
	subPath := fmt.Sprintf("viewer/v2/participants?tournament_ids=%s", tournamentID)

	participantId1 := 0
	participantId2 := 49
	for true {
		rangeString := fmt.Sprintf("participants=%d-%d", participantId1, participantId2)
		data, err := toornamentRest(subPath, rangeString)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal([]byte(data), &participants)

		for _, p := range participants {
			if p.Name == teamName {
				return p.ID, nil
			}
		}
		participantId1 = participantId1 + 50
		participantId2 = participantId2 + 50
	}
	return "", fmt.Errorf("could not find correct team")
}

func matches(team *s.Team) {
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
