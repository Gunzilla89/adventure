package main

import (
	"fmt"

	c "github.com/gunzilla89/adventure/computer"
)

func main() {

	//loop game
	runGame := true
	for runGame {

		c.ComTalk(0)
		answer := c.TakeInput()

		fmt.Print(answer)

	}

}
