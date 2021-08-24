package faceit

import (
	"github.com/jlehtimaki/toornament-csgo/pkg/structs"
	"log"
	"testing"
)

func TestIntegrationParticipantGetTeam(t *testing.T) {
	player := structs.Player{}
	player.CustomFields.SteamId = "76561198049624788"
	err := GetData(&player)
	if err != nil {
		log.Fatal(err)
	}
	if player.Faceit.Url == "" || player.Faceit.Id == "" {
		log.Fatalf("could not get Faceitdata for player: %s", player.Name)
	}
}
