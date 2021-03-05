package main

import (
	"fmt"
	f "github.com/jlehtimaki/toornament-csgo/pkg/faceit"
	t "github.com/jlehtimaki/toornament-csgo/pkg/toornament"
)

func main()  {
	team, err := t.GetParticipant("Polar Squad")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(team)
}
