package main

import (
	c "github.com/gunzilla89/adventure/computer"
)

func main() {

	playGame := true
	var computer c.Computer = c.Computer{Level: 0}

	for playGame {
		c.PlayCurrentLevel(computer.Level)
		computer.GoUpLevel(computer.Level)
	}

}
