package csgostats

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	k "github.com/jlehtimaki/toornament-csgo/pkg/kanaliiga"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

var csgostatsUrl = "http://localhost:8080/stats"

type PlayerRank struct {
	SteamId string
	Current string `json:"current"`
	Highest string `json:"highest"`
	Avg     string
}

func GetRank(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	steamId := c.Param("id")
	playerRank := PlayerRank{}
	url := fmt.Sprintf("%s/%s", csgostatsUrl, steamId)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(body, &playerRank)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	playerRank.SteamId = steamId
	highest, err := strconv.Atoi(playerRank.Highest)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	current, err := strconv.Atoi(playerRank.Current)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	playerRank.Avg = k.RankConverter(((highest + current) / 2), "MM")
	playerRank.Current = k.RankConverter(current, "MM")
	playerRank.Highest = k.RankConverter(highest, "MM")
	c.IndentedJSON(http.StatusOK, playerRank)
	return
}
