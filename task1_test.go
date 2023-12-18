package main

import (
	"fmt"
	"testing"
)

func TestTask1decipherMessage(t *testing.T) {
	{
		input := "И.РЮУ.ЪФОБГНО"
		result := task1decipherMessage(input)
		fmt.Printf("\nРезультат расшифровки %s - %s;", input, result)
	}
	{
		input := "CЛХГ.ЪЛХО.ФОО.ЩВ"
		result := task1decipherMessage(input)
		fmt.Printf("\nРезультат расшифровки %s - %s;", input, result)
	}
}
