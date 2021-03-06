package main

import (
	"fmt"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
)

func main()  {
	// Get team
	team, err := t.GetParticipant("Polar Squad")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// Loop through players and get their data
	for i, _ := range team.Players {
		err = f.GetData(&team.Players[i])
		break
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}

	fmt.Println(team)
}
