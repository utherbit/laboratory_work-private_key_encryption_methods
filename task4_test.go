package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask4decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 4\nРасшифровка сообщения по шифру Вижинера")
	decryptFunc := func(message string) string {
		return task4decipherMessage("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ_", message, "ОРЕХ")
	}

	utilities.CheckValidEncryption(t, decryptFunc, "ШВМБУЖНЯ", "КУЗНЕЧИК")
	utilities.CheckValidEncryption(t, decryptFunc, "ЯБХЪШЮМХ", "СТРЕКОЗА")
	print("\n")
}
