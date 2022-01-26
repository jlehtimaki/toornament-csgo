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

func TestIntegrationParticipantGetTeam2(t *testing.T) {
	_, err := GetTeam("PS. Tykitellään")
	if err != nil {
		log.Fatal(err)
	}
}

func TestIntegrationGetTeamID(t *testing.T) {
	wanted := "5331313737645744128"
	got, err := getTeamID("PS. Tykitellään", "5161204601415041024")
	if err != nil {
		log.Fatal(err)
	}
	if wanted != got {
		log.Fatalf("wanted (%s) and got %s", wanted, got)
	}
}
