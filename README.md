# toornament-csgo
This project is meant to provide player data (MM, Faceit, Esportal ranks, etc...) from a team in Toornament
Project runs in `https://csgoapi.lehtux.com` . For access, please contact me.
 
## Abilities
This REST API can deliver these values

| Path                |     Value     | Description                                                   |
|---------------------|:-------------:|---------------------------------------------------------------|
| team/:id            |   Team Name   | Gets data of the team                                         |
| standings/:id       |   Team Name   | Gets the standings either based on Division name or Team name |
| match/next/:id      |   Team Name   | Get Next match for the team                                   |
| match/scheduled/:id |   Team Name   | Get Next scheduled match for the team                         |
| match/all/:id       |   Team Name   | Get all matches for that team                                 |
| seed                |       -       | Get Seedings for all the divisions                            |
| seed/:id             | Division Name | Get Seeding for the division                                  |
| rank/mm/:id         |   Steam ID    | Gives MM rank fetched from csgostats                          |

### Environment variables
Function also requires certain environment variables to be able to fetch data

| Type                      | Description           |
|------                     |---------------------- |
| TOORNAMENT_API_KEY        | Toornament API key    |
| SEASON_ID                 | Toornament Season ID  |
| FACEIT_API_KEY            | Faceit API key        |
