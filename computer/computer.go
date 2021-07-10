package computer

import (
	"fmt"
	"strconv"
	"github.com/gunzilla89/utilities"
)

var myReader := utilities

type Computer struct {
	Level int
}

//computer talk
func PlayCurrentLevel(currentLevel int) {

	switch currentLevel {
	case 0:
		level1()

	case 1:
		level2()
	}

}

func (c *Computer) GoUpLevel(currentLevel int) {
	c.Level++
}

///////GAME LEVELS/////////////////

//LEVEL 1
func level1() {
	fmt.Print("Welcome to Josh's Adventure \nType your name and press Enter to submit: ")
	//Take input here
	fmt.Printf("Hello %s \n Let's being...", answer)
}

//LEVEL 2
func level2() {
	fmt.Print("you are on level 1")
	fmt.Print("You walked up and saw a flower. What do you do?")
	fmt.Print("1: Pick the flower")
	fmt.Print("2: Step on the flower")
	//take input here

	//need to slice away \n after input
	answerInt, err := strconv.Atoi(answer)
	if err != nil {
		panic(err)
	}

	switch answerInt {
	case 1:
		fmt.Print("You are so nice")

	case 2:
		fmt.Print("You are so mean!!!!")
	}

}
