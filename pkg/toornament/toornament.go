package toornament

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	toornament = "https://api.toornament.com"
	seasonId = os.Getenv("SEASON_ID")
	toornamentApiKey = os.Getenv("TOORNAMENT_API_KEY")
)

func toornamentRest(subPath string, rangeString string) ([]byte, error){
	apiUrl := fmt.Sprintf("%s/%s",toornament, subPath)

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Key", toornamentApiKey)
	request.Header.Set("Range", rangeString)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		return nil, fmt.Errorf("Statuscode: %d, message: %s", resp.StatusCode, resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getParticipantsToornament() ([]Team, error) {
	var teams []Team
	rangeMin := 0
	rangeMax := 49

	// if SeasonID or ToornamentApiKey is empty fail
	if seasonId == "" || toornamentApiKey == "" {
		return teams,
		fmt.Errorf("SEASON_ID (%s) is empty or TOORNAMENT_API_KEY (%s) is empty", seasonId, toornamentApiKey)
	}

	subPath := fmt.Sprintf("viewer/v2/tournaments/%s/participants", seasonId)
	apiUrl := fmt.Sprintf("%s/%s",toornament, subPath)

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Key", toornamentApiKey)
	client := &http.Client{}

	for rangeMin < rangeMax {
		var tmpTeam []Team
		rangeHeader := fmt.Sprintf("participants=%d-%d", rangeMin, rangeMax)
		request.Header.Set("Range", rangeHeader)
		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		maxPagination, err := strconv.Atoi(strings.Split(response.Header.Get("Content-Range"),"/")[1])
		if err != nil {
			return nil, err
		}

		// Read the response
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		// Unmarshal the return value to Team struct and append it to return value
		_ = json.Unmarshal(data, &tmpTeam)
		teams = append(teams, tmpTeam...)

		// If we get any other than Partial content HTTP return message, break the loop since no pagination exists
		if response.Status != "206 Partial Content" {
			break
		}

		// Increase pagination by 50 which is the Toornament pagination limit, stupid, I know...
		// If range is greater than maxPagination, assign it to rangeMax
		rangeMin = rangeMax + 1
		rangeMax = rangeMax + 50
		if rangeMax >= maxPagination {
			rangeMax = maxPagination
		}
	}
	return teams, nil
}

func GetParticipant(teamName string) (Team, error){
	// First get all teams and get the ID of the team
	teams, err := getParticipantsToornament()
	if err != nil {
		return Team{}, err
	}

	for _, t := range teams{
		if t.Name == teamName {
			t.matches()
			return t, nil
		}
	}

	return Team{}, fmt.Errorf("could not find the team")
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