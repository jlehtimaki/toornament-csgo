package kanaliiga

import (
	"github.com/jlehtimaki/toornament-csgo/pkg/structs"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestIntegrationParticipantGetTeam(t *testing.T) {
	player := structs.Player{}
	player.CustomFields.SteamId = "76561198049624788"
	err := GetData(&player)
	if err != nil {
		log.Fatal(err)
	}
}

func TestIntegrationGetTeamID(t *testing.T) {
	wanted := "5330666641945853952"
	got, err := GetTeamID("PS. Tykitellään")
	if err != nil {
		log.Fatal(err)
	}
	if wanted != got {
		log.Fatalf("got %s - wanted %s", got, wanted)
	}
}
