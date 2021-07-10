package utilities

import (
	"bufio"
	"os"
)

//reader to take string input from user
func TakeUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = response[:len(response)-2]
	return response
}
