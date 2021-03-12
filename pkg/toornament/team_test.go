package toornament

import (
	"fmt"
	"log"
	"testing"
)

func TestIntegrationParticipantGetTeam(t *testing.T) {
	team, err := GetTeam("Polar Squad")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(team)
}
