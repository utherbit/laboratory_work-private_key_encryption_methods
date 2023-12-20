package main

import (
	"fmt"
	"testing"
)

func TestTask6FindGamma(t *testing.T) {
	fmt.Printf("\n\nЗадача 6\nПоиск ключа гаммирования")

	gamma, err := task6FindGamma("74", "9А", 16)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("\nвывод: %v", gamma)
	print("\n")
}
