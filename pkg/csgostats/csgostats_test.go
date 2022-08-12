package csgostats

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestIntegrationParticipantGetTeam(t *testing.T) {
	r := SetUpRouter()
	r.GET("/", GetRank)
	//steamId := "76561198049624788"
	req, _ := http.NewRequest("GET", fmt.Sprintf("/seed"), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(responseData))
	//GetRank(c)
	//steamId = "76561197960288540"
	//GetRank(c)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if player.Faceit.Url == "" || player.Faceit.Id == "" {
	//	log.Fatalf("could not get Faceitdata for player: %s", player.Name)
	//}
}
