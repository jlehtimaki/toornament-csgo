package faceit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	faceItAPI = "https://open.faceit.com/data/v4"
)

func faceItRest(subPath string) ([]byte, int, error) {
	// Import Faceit API Key
	faceitApiKey := os.Getenv("FACEIT_API_KEY")
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

func GetRank(nick string, steamid string) (int, int, string, error){
	var faceitPlayer Player
	getPlayerUrl := fmt.Sprintf("players?nickname=%s&game=csgo", nick)
	data, statusCode, err := faceItRest(getPlayerUrl)
	if err != nil {
		return 0, 0, "", err
	}
	if statusCode == 404 {
		data, err = searchPlayer(nick, steamid)
		if err != nil {
			return 0, 0, "", err
		}
	}
	_ = json.Unmarshal(data, &faceitPlayer)

	if steamid != faceitPlayer.Steam64 {
		data, err = searchPlayer(nick, steamid)
		if err != nil {
			return 0,0,"", err
		}
		_ = json.Unmarshal(data, &faceitPlayer)
	}

	return faceitPlayer.Games.CSGO.SkillLevel,
	faceitPlayer.Games.CSGO.Elo,
	strings.ReplaceAll(faceitPlayer.FaceitUrl,"{lang}", "en"),
	nil
}

func searchPlayer(nick string, steamid string) ([]byte ,error) {
	var searchResult SearchResult
	var player	Player
	searchSubPath := fmt.Sprintf("search/players?nickname=%s&game=csgo&offset=0&limit=20", nick)
	result, _, err := faceItRest(searchSubPath)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(result, &searchResult)
	for _, item := range searchResult.Items {
		playerSubPath := fmt.Sprintf("players?nickname=%s&game=csgo", item.Nickname)
		playerData, _, err := faceItRest(playerSubPath)
		if err != nil {
			return nil, err
		}
		_ = json.Unmarshal(playerData, &player)
		if steamid == player.Steam64 {
			return playerData, nil
		}
	}
	return nil, fmt.Errorf("could not find any faceit user for: %s", nick)
}