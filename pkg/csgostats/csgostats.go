package csgostats

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

var csgostatsUrl = "http://csgostats:8080/stats"

type PlayerRank struct {
	SteamId string
	Current int
	Highest int
	Avg     int
}

type CsgoStats struct {
	Current string `json:"current"`
	Highest string `json:"highest"`
}

func GetRank(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	steamId := c.Param("id")
	csgoStats := CsgoStats{}
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
	err = json.Unmarshal(body, &csgoStats)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	highest, err := strconv.Atoi(csgoStats.Highest)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	current, err := strconv.Atoi(csgoStats.Current)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	playerRank := PlayerRank{
		SteamId: steamId,
		Current: current,
		Highest: highest,
		Avg:     (highest + current) / 2,
	}
	c.IndentedJSON(http.StatusOK, playerRank)
	return
}
