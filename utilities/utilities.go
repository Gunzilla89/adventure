package utilities

import (
	"bufio"
	"os"
	"strconv"
)

type MyReader interface {
	TakeInput()
	TakeIntInput()
}

type stringInput struct{}

type intInput struct{}

//reader to take input from user
func (si *stringInput) TakeStringInput() string {

	//reader for user input
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return response
}

func (intInput *intInput) TakeIntInput() int {

	//reader for user input
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	intResponse, _ := strconv.Atoi(response)
	return intResponse
}
