package object

import (
	"fmt"
)

// create player
var player Player

type Computer struct {
	Level int
}

//computer talk
func PlayCurrentLevel(currentLevel int) {

	switch currentLevel {
	case 0:
		level0()

	case 1:
		level1()
	}

}

func (c *Computer) GoUpLevel(currentLevel int) {
	c.Level++
}

///////GAME LEVELS/////////////////
//LEVEL 1
func level0() {
	fmt.Print("Welcome to Josh's Adventure \nType your name and press Enter to submit: ")
	answer := "Josh"
	player.setName(answer)
	fmt.Printf("Hello %s \n Let's begin...", player.Name)
}

//LEVEL 2 Question
func level1() {

	fmt.Println("LEVEL 1")
	fmt.Print("\n")

	fmt.Println("Welcome to the Gloom Tavern! We hope that you enjoy your stay." +
		"You take a look around to see musicians and drunkards doing what they do best." +
		"as you scan the room you also notice a man in a cloak by himslef. You can't make out his face.")

	fmt.Println("What do you do?")
	fmt.Println("1: Cheerfuly greet the man")
	fmt.Println("2: Ignore him and get a drink")
	fmt.Println("2: Buy the man a drink and ask to sit with him")
	fmt.Println("4: Walk up and demand why he sits alone")

}
