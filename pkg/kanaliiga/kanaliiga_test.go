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
		log.Fatalf("get participant: %s", err)
	}
}

func TestIntegrationGetTeamID(t *testing.T) {
	wanted := 48
	got, err := GetTeamLeagueID("Polar Squad")
	if err != nil {
		log.Fatalf("getTeamId: %s", err)
	}
	if wanted != got {
		log.Fatalf("got %d - wanted %d", got, wanted)
	}
}

//func TestIsScheduled(t *testing.T) {
//	if IsScheduled("Polar Squad", "Nets") {
//		fmt.Println("ya baby")
//		return
//	}
//	log.Fatal("did not work")
//}

func TestGetScheduledMatches(t *testing.T) {
	_, err := scheduledMatches("Polar Squad")
	if err != nil {
		log.Fatal(err)
	}
}
