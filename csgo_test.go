package CSGO

import (
	"encoding/json"
	"fmt"
	"github.com/jlehtimaki/toornament-csgo/pkg/toornament"
	"io/ioutil"
	"log"
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
		log.Fatal("Something went wrong")
	}

	var team toornament.Team
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