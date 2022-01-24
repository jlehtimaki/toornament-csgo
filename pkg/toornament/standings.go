package toornament

import (
	"encoding/json"
	"fmt"
	str "github.com/jlehtimaki/toornament-csgo/pkg/structs"
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

func getRanking(stageId string) ([]byte, error) {
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
	return data, nil
}

func GetStandings(s string) ([]byte, error) {
	stages, err := getStages()
	if err != nil {
		return nil, err
	}

	// Get standings for a team
	team, err := GetTeam(s)
	if err == nil {
		for _, stage := range stages {
			var standings str.Standings
			ret, err := getRanking(stage.ID)
			if err != nil {
				return nil, err
			}
			_ = json.Unmarshal(ret, &standings)
			for _, div := range standings {
				if div.Participant.Name == team.Name {
					subUrl := fmt.Sprintf(
						"viewer/v2/tournaments/%s/stages/%s/ranking-items?group_ids=%s",
						seasonId,
						stage.ID,
						div.GroupID)
					fmt.Println(subUrl)
					rangeUrl := "items=0-49"
					ret, err := toornamentRest(subUrl, rangeUrl)
					if err != nil {
						fmt.Println("foobar")
						return nil, err
					}
					return ret, nil
				}
			}
		}
	}

	// Get standings based on division name
	for _, stage := range stages {
		if stage.Name == s {
			ret, err := getRanking(stage.ID)
			if err != nil {
				return nil, err
			}
			return ret, nil
		}
	}
	return nil, fmt.Errorf("could not find any standings")
}
