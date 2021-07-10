package main

import (
	com "github.com/gunzilla89/adventure/computer"
	levelSim "github.com/gunzilla89/adventure/levelSim"
	utility "github.com/gunzilla89/adventure/utilities"
)

func main() {

	playGame := true
	var computer com.Computer = com.Computer{Level: 0}

	for playGame {
		com.PlayCurrentLevel(computer.Level)
		computer.GoUpLevel(computer.Level)

		com.PlayCurrentLevel(computer.Level)
		answer := utility.TakeUserInput()
		levelSim.MainLevel1(answer)

		playGame = false

	}

}
