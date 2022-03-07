package toornament

import (
	"fmt"
	"github.com/gin-gonic/gin"
	k "github.com/jlehtimaki/toornament-csgo/pkg/kanaliiga"
	u "github.com/jlehtimaki/toornament-csgo/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NextMatch(c *gin.Context) {
	if !u.Verify(c) {
		c.IndentedJSON(http.StatusForbidden, "authentication failure")
		return
	}
	team := c.Param("id")
	t, err := GetTeam(team)
	if err != nil {
		log.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	for _, m := range t.Matches {
	opponent:
		for _, o := range m.Opponents {
			if o.Participant.Name != team && o.Result == "" {
				if k.IsScheduled(t.Name, o.Participant.Name) {
					break opponent
				}
				t2, err := GetTeam(o.Participant.Name)
				if err != nil {
					log.Error(err)
					c.IndentedJSON(http.StatusInternalServerError, err)
					return
				}
				c.IndentedJSON(http.StatusOK, t2)
				return
			}
		}
	}
	log.Error(fmt.Errorf("something went wrong with getting next match"))
	c.IndentedJSON(http.StatusInternalServerError, err)
	return
}
