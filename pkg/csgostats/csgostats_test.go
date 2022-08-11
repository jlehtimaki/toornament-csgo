package csgostats

import (
	"testing"
)

func TestIntegrationParticipantGetTeam(t *testing.T) {
	steamId := "76561198049624788"
	_ = GetRank(steamId)
	steamId = "76561197960288540"
	_ = GetRank(steamId)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if player.Faceit.Url == "" || player.Faceit.Id == "" {
	//	log.Fatalf("could not get Faceitdata for player: %s", player.Name)
	//}
}
