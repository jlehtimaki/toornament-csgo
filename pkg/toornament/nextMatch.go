package toornament

import (
	"encoding/json"
	"fmt"
)

func NextMatch(team string) ([]byte, error) {
	t, err := GetTeam(team)
	if err != nil {
		return nil, err
	}

	for _, t := range t.Matches[0].Opponents {
		if t.Participant.Name != team && t.Result == "" {
			t2, err := GetTeam(t.Participant.Name)
			if err != nil {
				return nil, err
			}
			b, err := json.MarshalIndent(t2, "", "  ")
			if err != nil {
				return nil, err
			}
			return b, nil
		}
	}
	return nil, fmt.Errorf("something went wrong with getting next match")
}
