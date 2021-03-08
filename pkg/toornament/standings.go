package toornament

import (
	"encoding/json"
	"fmt"
)

func getStages()(Stages, error){
	var stages Stages
	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/stages", seasonId)

	data, err := toornamentRest(subPath, "")
	if err != nil {
		return stages, nil
	}
	_ = json.Unmarshal(data, &stages)
	return stages, nil
}

func getRanking(stageId string) ([]byte, error) {
	var standings Standings
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
	return data,nil
}

func GetStandings(s string) ([]byte, error){
	stages, err := getStages()
	if err != nil {
		return nil, err
	}

	//team, err := GetParticipant(s)
	//if err == nil {
	//	for _, stage := range stages {
	//		stage.
	//	}
	//}

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