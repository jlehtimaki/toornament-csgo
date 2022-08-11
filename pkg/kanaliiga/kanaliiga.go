package kanaliiga

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	kanaliigaApi   = "https://cssapi.kanalan.it"
	kanaliigaToken = os.Getenv("KANALIIGA_TOKEN")
	seasonID       = os.Getenv("SEASON")
)

func restCall(subPath string) ([]byte, error) {
	if kanaliigaToken == "" {
		return nil, fmt.Errorf("KANALIIGA_TOKEN is missing")
	}
	apiUrl := fmt.Sprintf("%s/%s?token=%s", kanaliigaApi, subPath, kanaliigaToken)
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("statuscode: %d message: %s", response.StatusCode, response.Status)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetData(player *s.Player) error {
	var kanaliigaPlayer s.Kanaliiga
	var (
		kills      int
		assists    int
		deaths     int
		mvps       int
		adr        float64
		hsp        int
		kast       int
		kanarating float64
	)

	subPath := fmt.Sprintf("player/%s", player.CustomFields.SteamId)
	data, err := restCall(subPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &kanaliigaPlayer)
	if err != nil {
		return err
	}

	if len(kanaliigaPlayer.Stats) > 0 {
		for _, s := range kanaliigaPlayer.Stats {
			kills = kills + s.Kills
			assists = assists + s.Assists
			deaths = deaths + s.Deaths
			mvps = mvps + s.MVPs
			adr = adr + s.ADR
			hsp = hsp + s.HsPercent
			kast = kast + s.KAST
			kanarating = kanarating + s.KanaRating
		}

		player.Kanaliiga.Kills = kills
		player.Kanaliiga.Assists = assists
		player.Kanaliiga.Deaths = deaths
		player.Kanaliiga.MVPs = mvps
		player.Kanaliiga.KDR = kills / deaths
		player.Kanaliiga.ADR = adr / float64(len(kanaliigaPlayer.Stats))
		player.Kanaliiga.HsPercent = hsp / len(kanaliigaPlayer.Stats)
		player.Kanaliiga.KAST = kast / len(kanaliigaPlayer.Stats)
		player.Kanaliiga.KanaRating = kanarating / float64(len(kanaliigaPlayer.Stats))
	}

	getRanks(player)

	return nil
}

func getKanaTeams() (s.KanaToornament, error) {
	var kanaData s.KanaToornament
	subPath := fmt.Sprintf("teams/toornament/%s", seasonID)
	data, err := restCall(subPath)
	if err != nil {
		return s.KanaToornament{}, err
	}

	err = json.Unmarshal(data, &kanaData)
	if err != nil {
		return s.KanaToornament{}, err
	}

	return kanaData, nil
}

func GetTeamID(teamName string) (int, error) {
	kanaData, err := getKanaTeams()
	if err != nil {
		return 0, err
	}
	for _, d := range kanaData.Data {
		if d.Name == teamName {
			return d.ID, nil
		}
	}

	return 0, fmt.Errorf("could not find correct team")
}

func GetTeamLeagueID(teamName string) (int, error) {
	kanaData, err := getKanaTeams()
	if err != nil {
		return 0, err
	}

	for _, d := range kanaData.Data {
		if d.Name == teamName {
			return d.LeagueID, nil
		}
	}

	return 0, fmt.Errorf("could not find correct team")
}

func GetMatches(teamName string) ([]s.KanaliigaMatch, error) {
	type Matches struct {
		Data []s.KanaliigaMatch `json:"data"`
	}
	matches := Matches{}

	// First get leagueID
	leagueID, err := GetTeamLeagueID(teamName)
	if err != nil {
		return []s.KanaliigaMatch{}, err
	}

	// Get all matches for that team
	subPath := fmt.Sprintf("matches/%d", leagueID)
	data, err := restCall(subPath)
	if err != nil {
		return []s.KanaliigaMatch{}, err
	}

	err = json.Unmarshal(data, &matches)
	if err != nil {
		return []s.KanaliigaMatch{}, err
	}
	return matches.Data, nil
}

func getRanks(player *s.Player) {
	type Ranks struct {
		Data []s.KanaliigaRanks `json:"data"`
	}
	ranks := Ranks{}
	subPath := fmt.Sprintf("ranks/%s", player.CustomFields.SteamId)
	data, err := restCall(subPath)
	if err != nil {
		log.Error(err)
		return
	}

	err = json.Unmarshal(data, &ranks)
	if err != nil {
		log.Error(err)
		return
	}

	player.Esportal.Rank = RankConverter(ranks.Data[0].EsportalRank, "Esportal")
	player.MM.Rank = RankConverter(ranks.Data[0].Rank, "MM")
}

func GetScheduledMatches(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	team := c.Param("id")
	s, err := scheduledMatches(team)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, s)
	return
}

func scheduledMatches(teamName string) ([]s.ScheduledMatch, error) {
	var calendar s.Calendar
	var scheduledMatches []s.ScheduledMatch

	teamID, err := GetTeamID(teamName)
	if err != nil {
		return nil, err
	}

	subPath := fmt.Sprintf("calendar/%d", teamID)
	data, err := restCall(subPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &calendar)
	if err != nil {
		return nil, err
	}

	teams, err := getKanaTeams()
	if err != nil {
		return nil, err
	}
	for _, t := range teams.Data {
		for n, m := range calendar.Data {
			if calendar.Data[n].Team1Name != "" && calendar.Data[n].Team2Name != "" {
				break
			}
			if m.Team1 == t.ID {
				calendar.Data[n].Team1Name = t.Name
			}
			if m.Team2 == t.ID {
				calendar.Data[n].Team2Name = t.Name
			}
		}
	}

	today := time.Now()
	for _, match := range calendar.Data {
		if today.Before(match.Date) {
			scheduledMatches = append(scheduledMatches, match)
		}
	}

	return scheduledMatches, nil
}

func IsScheduled(team1 string, team2 string) bool {
	matches, err := scheduledMatches(team1)
	if err != nil {
		log.Error(err)
		return false
	}

	team2ID, err := GetTeamID(team2)
	if err != nil {
		log.Error(err)
		return false
	}

	for _, match := range matches {
		if match.Team1 == team2ID || match.Team2 == team2ID {
			return true
		}
	}
	return false
}
