package CSGO

import (
	"encoding/json"
	"fmt"
	"github.com/jlehtimaki/toornament-csgo/pkg/structs"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIntegrationCSGOFunction(t *testing.T){
	payload := strings.NewReader(`
{
	"type": "team",
	"value": "Polar Squad"
}`)
	req := httptest.NewRequest("GET", "/", payload)
	w := httptest.NewRecorder()
	CSGO(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("Something went wrong: %d message: %s", resp.StatusCode, resp.Status)
	}

	var team structs.Team
	err := json.Unmarshal(body, &team)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(team)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

func TestStandingsIntegration(t *testing.T){
	payload := strings.NewReader(`
{
	"type": "standings",
	"value": "4.Div"
}`)
	req := httptest.NewRequest("GET", "/", payload)
	w := httptest.NewRecorder()
	CSGO(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal("Something went wrong")
	}

	var standings structs.Standings
	err := json.Unmarshal(body, &standings)
	if err != nil {
		log.Fatal(err)
	}

	_, err = json.Marshal(standings)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(jsonData))
}