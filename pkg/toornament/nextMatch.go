package toornament

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NextMatch(c *gin.Context) {
	team := c.Param("id")
	t, err := GetTeam(team)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	for _, t := range t.Matches[0].Opponents {
		if t.Participant.Name != team && t.Result == "" {
			t2, err := GetTeam(t.Participant.Name)
			if err != nil {
				log.Error(err)
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
			c.IndentedJSON(http.StatusOK, t2)
			return
		}
	}
	log.Error(fmt.Errorf("something went wrong with getting next match"))
	c.IndentedJSON(http.StatusInternalServerError, err)
	return
}
