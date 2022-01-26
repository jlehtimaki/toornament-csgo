package toornament

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	toornament       = "https://api.toornament.com"
	seasonId         = os.Getenv("SEASON_ID")
	toornamentApiKey = os.Getenv("TOORNAMENT_API_KEY")
)

func toornamentRest(subPath string, rangeString string) ([]byte, error) {
	apiUrl := fmt.Sprintf("%s/%s", toornament, subPath)

	if toornamentApiKey == "" {
		return nil, fmt.Errorf("could not find TOORNAMENT_API_KEY")
	}

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Api-Key", toornamentApiKey)
	request.Header.Set("Range", rangeString)
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		return nil, fmt.Errorf("statuscode: %d, message: %s", resp.StatusCode, resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if string(data) == "[]" {
		return nil, fmt.Errorf("toornamentRest: got an empty array")
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
