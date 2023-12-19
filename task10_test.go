package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTask10decipherMessage(t *testing.T) {

	fmt.Printf("%v", strconv.FormatInt(42, 2))

	fmt.Printf("\nRes: %s", task10decipherMessage("ЕАУПД_КЕАЗАРЧВ", []int{4, 1, 2, 3}))
	fmt.Printf("\nRes: %s", task10decipherMessage("А_Н_СЫИЛБСА_ЛЙГ_", []int{3, 1, 4, 2}))
}
