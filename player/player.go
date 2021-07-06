package main

type player struct {
	Name string
}

//player stats
func (player player) playerStats(playerName string) {
	player.Name = playerName
}
