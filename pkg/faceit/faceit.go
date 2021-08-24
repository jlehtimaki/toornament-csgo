package faceit

import (
	"encoding/json"
	"fmt"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	faceItAPI = "https://open.faceit.com/data/v4"
)

func faceItRest(subPath string) ([]byte, int, error) {
	// Import Faceit API Key
	faceitApiKey := os.Getenv("FACEIT_API_KEY")
	if faceitApiKey == "" {
		return nil, 0, fmt.Errorf("could not find FACEIT_API_KEY")
	}
	apiUrl := fmt.Sprintf("%s/%s", faceItAPI, subPath)
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil,500, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", faceitApiKey))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil,response.StatusCode, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,response.StatusCode, err
	}
	return data,response.StatusCode, nil
}

func GetData(player *s.Player) error{
	var faceitPlayer s.FaceitPlayer
	subPath := fmt.Sprintf("players?game=csgo&game_player_id=%s", player.CustomFields.SteamId)
	data, statusCode, err := faceItRest(subPath)
	if err != nil {
		return err
	}
	if statusCode == 404 || statusCode == 400 {
			return fmt.Errorf("could not find user: %s", player.Name)
	}
	_ = json.Unmarshal(data, &faceitPlayer)

	player.Faceit.Elo = faceitPlayer.Games.CSGO.Elo
	player.Faceit.Rank = faceitPlayer.Games.CSGO.SkillLevel
	player.Faceit.Url = strings.ReplaceAll(faceitPlayer.FaceitUrl, "{lang}", "en")
	player.Faceit.Avatar = faceitPlayer.Avatar
	player.Faceit.Id = faceitPlayer.Id

	err = getStats(player)
	if err != nil {
		return err
	}
	return nil
}

func getStats(player *s.Player) error {
	var stats s.Stats
	subPath := fmt.Sprintf("players/%s/stats/csgo", player.Faceit.Id)
	data, statusCode, err := faceItRest(subPath)
	if statusCode != 200 || err != nil{
		return err
	}
	_ = json.Unmarshal(data, &stats)

	// Find favourite map and least favourite + stats from them
	mostPlayed := stats.Maps[0]
	leastPlayed := stats.Maps[0]
	for _, m := range stats.Maps {
		if m.Mode != "5v5" {
			continue
		}
		mInt, _ := strconv.Atoi(m.Stats.Matches)
		mostInt, _ := strconv.Atoi(mostPlayed.Stats.Matches)
		leastInt, _ := strconv.Atoi(leastPlayed.Stats.Matches)
		if mInt > mostInt {
			mostPlayed = m
		}
		if mInt < leastInt {
			leastPlayed = m
		}
	}

	// Save data to player object
	player.Faceit.KD = stats.Overall.AverageKDRatio
	player.Faceit.HSP = stats.Overall.AverageHeadshots
	// Save MostPlayedMap data
	player.Faceit.MostPlayedMap.Name = mostPlayed.Name
	player.Faceit.MostPlayedMap.Matches, _ = strconv.Atoi(mostPlayed.Stats.Matches)
	player.Faceit.MostPlayedMap.KD = mostPlayed.Stats.KD
	player.Faceit.MostPlayedMap.WinRate = mostPlayed.Stats.WinRate
	player.Faceit.MostPlayedMap.Icon = mostPlayed.Image
	// Save LeastPlayedMap data
	player.Faceit.LeastPlayedMap.Name = leastPlayed.Name
	player.Faceit.LeastPlayedMap.Matches, _ = strconv.Atoi(leastPlayed.Stats.Matches)
	player.Faceit.LeastPlayedMap.KD = leastPlayed.Stats.KD
	player.Faceit.LeastPlayedMap.WinRate = leastPlayed.Stats.WinRate
	player.Faceit.LeastPlayedMap.Icon = leastPlayed.Image

	return nil
}