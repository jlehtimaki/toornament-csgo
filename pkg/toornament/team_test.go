package toornament

import (
	"log"
	"testing"
)

func TestIntegrationParticipantGetTeam(t *testing.T) {
	_, err := GetTeam("Polar Squad")
	if err != nil {
		log.Fatal(err)
	}
}
