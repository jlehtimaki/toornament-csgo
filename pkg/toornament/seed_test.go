package toornament

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestIntegrationGetAllSeeds(t *testing.T) {
	data, err := GetSeed("")
	if err != nil {
		log.Fatal(err)
	}
	if data == nil {
		log.Fatal("data was empty")
	}
}
