package main

import (
	"fmt"
	"main/utilities"
)

func main() {
	var input string
	_, err := fmt.Scanf("%s", &input)
	utilities.PanicIfErr(err)

	result := task1decipherMessage(input)
	fmt.Printf("\n%s", result)
}
