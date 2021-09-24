package toornament

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestIntegrationGetGroups(t *testing.T) {
	data, err := getGroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}