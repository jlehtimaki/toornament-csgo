# toornament-csgo
This project is meant to provide player data (MM, Faceit, Esportal ranks, etc...) from a team in Toornament
Project is optimised to be run in Cloud Functions.
 
## Usage
This fuction expects the payload to have two values `type` and `value`

| Type          | Values        | Description           |
|------         |:---------:    |---------------------- |
|team           | Team Name     | Gets data of the team |
|standings      | Division name | Gets the standings either based on Division name or Team name   |
|seed           | Empty for all Divs or Div name for specific             | Get the current seed of the tournament

### Environment variables
Function also requires certain environment variables to be able to fetch data

| Type                      | Description           |
|------                     |---------------------- |
| TOORNAMENT_API_KEY        | Toornament API key    |
| SEASON_ID                 | Toornament Season ID  |
| FACEIT_API_KEY            | Faceit API key        |
