package toornament

import (
	"testing"
)
import log "github.com/sirupsen/logrus"

func TestIntegrationGetNextMatch(t *testing.T) {
	data, err := NextMatch("Polar Squad")
	if err != nil {
		log.Fatal(err)
	}
	if data == nil {
		log.Fatal("next match - data was empty")
	}
}
