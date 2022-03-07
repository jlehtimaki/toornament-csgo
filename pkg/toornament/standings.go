package toornament

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	str "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	"net/http"
)

func getStages() (str.Stages, error) {
	var stages str.Stages
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/stages", seasonId)

	data, err := toornamentRest(subPath, "")
	if err != nil {
		return stages, nil
	}
	_ = json.Unmarshal(data, &stages)
	return stages, nil
}

func getRanking(stageId string) (str.Standings, error) {
	var standings str.Standings
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/stages/%s/ranking-items", seasonId, stageId)
	rangeString := "items=0-49"
	data, err := toornamentRest(subPath, rangeString)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &standings)
	if err != nil {
		return nil, err
	}
	return standings, nil
}

func GetStandings(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	s := c.Param("id")
	standings, err := standings(s)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, standings)
}

func standings(s string) (str.Standings, error) {
	standings := str.Standings{}
	stages, err := getStages()
	if err != nil {
		return standings, err
	}

	// Get standings for a team
	team, err := GetTeam(s)
	if err == nil {
		for _, stage := range stages {
			standings, err := getRanking(stage.ID)
			if err != nil {
				return standings, err
			}
			for _, div := range standings {
				if div.Participant.Name == team.Name {
					subUrl := fmt.Sprintf(
						"viewer/v2/tournaments/%s/stages/%s/ranking-items?group_ids=%s",
						seasonId,
						stage.ID,
						div.GroupID)
					rangeUrl := "items=0-49"
					ret, err := toornamentRest(subUrl, rangeUrl)
					if err != nil {
						return standings, err
					}
					err = json.Unmarshal([]byte(ret), &standings)
					if err != nil {
						return standings, err
					}
					return standings, err
				}
			}
		}
	}

	// Get standings based on division name
	for _, stage := range stages {
		if stage.Name == s {
			standings, err := getRanking(stage.ID)
			if err != nil {
				return standings, err
			}
			return standings, nil
		}
	}
	return standings, fmt.Errorf("could not find correct standings")
}
