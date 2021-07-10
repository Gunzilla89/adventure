package levelSim

import (
	"fmt"
)

func MainLevel1(answer string) {

	//player decision
	switch answer {
	case "1":
		fmt.Println("The man stares and says nothing...")
		fmt.Println("Press enter to continue...")
		level1Branch1()

	case "2":
		fmt.Println("rude.. go say hello")

	}
}

func level1Branch1() {
	fmt.Print("you've entered a branch")
}
