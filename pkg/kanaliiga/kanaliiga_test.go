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
