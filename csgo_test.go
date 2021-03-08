package CSGO

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIntegrationCSGOFunction(t *testing.T){
	payload := strings.NewReader(`
{
	"team_name": "Polar Squad"
}`)
	req := httptest.NewRequest("GET", "/", payload)
	values := map[string][]string{}
	values["team_name"] = append(values["team_name"], "Polar Squad")
	w := httptest.NewRecorder()
	CSGO(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}