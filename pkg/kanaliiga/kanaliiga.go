package kanaliiga

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	s "github.com/jlehtimaki/toornament-csgo/pkg/structs"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	kanaliigaApi = "https://tilastot.kanaliiga.fi:8443"
	kanaliigaToken = os.Getenv("KANALIIGA_TOKEN")
)

func restCall(subPath string) ([]byte, error){
	if kanaliigaToken == "" {
		return nil, fmt.Errorf("KANALIIGA_TOKEN is missing")
	}
	apiUrl := fmt.Sprintf("%s/%s?token=%s", kanaliigaApi, subPath,kanaliigaToken)
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
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
		kills int
		assists int
		deaths int
		mvps	int
		adr		float64
		hsp		int
		kast	int
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

	return nil
}
