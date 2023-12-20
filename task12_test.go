package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask12DecryptMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 12\nРасшифровка сообщения по таблице замен")
	utilities.CheckValidEncryption(t, task12decipherMessage, "353214764134136759136762849754128212350354035767106216753211", "ОТКРЫТЫИ КАНАЛ СВЯЗИ")
	utilities.CheckValidEncryption(t, task12decipherMessage, "351761756130532128759353134758105757213101752352763211762", "ГЕНЕРАТОР СООБЩЕНИИ")
	print("\n")
}
