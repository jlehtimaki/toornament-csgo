package faceit

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	faceItAPI = "https://open.faceit.com/data/v4"
)

func faceItRest(subPath string) ([]byte, int, error) {
	apiUrl := fmt.Sprintf("%s/%s", faceItAPI, subPath)
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil,500, err
	}
	request.Header.Set("Bearer", faceitApiKey)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil,response.StatusCode, err
	}
	if response.StatusCode != 200{
		return nil, response.StatusCode, nil
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,response.StatusCode, err
	}
	return data,response.StatusCode, nil
}

func GetRank(player Player) (string, error){
	getPlayerUrl := fmt.Sprintf("players?nickname=%s&game=csgo", player.Name)
	playerData, statusCode, err := faceItRest(getPlayerUrl)
	if err != nil {
		return "", err
	}
	if statusCode == 404 {
		//searchPlayer := fmt.Sprintf("search/players?nickname=%s&game=csgo&offset=0&limit=20", player.Name)
	}
}