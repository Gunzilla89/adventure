package computer

import (
	"bufio"
	"fmt"
	"os"
)

//computer talk
func ComTalk(currentLevel int) {

	if currentLevel == 0 {
		fmt.Print("Welcome to Josh's Adventure \nType your name and press Enter to submit: ")
	}

}

//reader to take input from user
func TakeInput() string {
	//readering for user input
	reader := bufio.NewReader(os.Stdin)

	myname, _ := reader.ReadString('\n')
	return myname
}
