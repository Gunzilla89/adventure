package main

import (
	"bufio"
	"os"

	c "./computer"
)

type myReader struct{}

type player struct {
	Name string
}

//reader to take input from user
func (mr myReader) TakeInput() string {
	//readering for user input
	reader := bufio.NewReader(os.Stdin)

	myname, _ := reader.ReadString('\n')
	return myname
}

//player stats
func (player player) playerStats(playerName string) {
	player.Name = playerName
}

func main() {
	//game set up
	//create level counter
	//currentLevel := 0
	//create reader to take in input
	//myReader := myReader{}
	//create the computer

	//create the player
	//player := player{Name: myReader.TakeInput()}

	//loop game
	runGame := true
	for runGame {

		c.ComTalk(0)

	}

}
