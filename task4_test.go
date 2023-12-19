package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask4decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 4\nРасшифровка сообщения по шифру Вижинера")
	utilities.CheckValidEncryption(t, task4decipherMessage, "ШВМБУЖНЯ", "КУЗНЕЧИК")
	utilities.CheckValidEncryption(t, task4decipherMessage, "ЯБХЪШЮМХ", "СТРЕКОЗА")
	print("\n")
}
