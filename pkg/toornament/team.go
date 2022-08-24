package toornament

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
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
	team.Matches = matches(team.Id)
	return team, nil
}

func getTeamID(teamName string, tournamentID string) (string, error) {
	type Participant struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	var participants []Participant
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/participants", tournamentID)

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
			if strings.ToLower(p.Name) == strings.ToLower(teamName) {
				return p.ID, nil
			}
		}
		participantId1 = participantId1 + 50
		participantId2 = participantId2 + 50
	}
	return "", fmt.Errorf("could not find correct team")
}

func GetMatches(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "getTeam: authentication failure")
		return
	}
	teamName := c.Param("id")
	teamId, err := getTeamID(teamName, seasonId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, matches(teamId))
	return
}

func matches(teamId string) s.Matches {
	var matches s.Matches
	subPath := fmt.Sprintf(
		"viewer/v2/tournaments/%s/matches?participant_ids=%s", seasonId, teamId)
	rangeString := "matches=0-50"
	data, err := toornamentRest(subPath, rangeString)
	if err != nil {
		log.Fatalf("could not get next games: %s", err)
	}
	_ = json.Unmarshal(data, &matches)
	return matches
}
