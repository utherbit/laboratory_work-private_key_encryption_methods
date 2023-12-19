package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask3decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 3\nШифрование сообщения по шифру Вижинера")
	utilities.CheckValidEncryption(t, task3decipherMessage, "КРИПТОСТОЙКОСТЬ", "ЙСФЮЭЭРУЪШХЭРУЗ")
	utilities.CheckValidEncryption(t, task3decipherMessage, "ГАММИРОВАНИЕ", "ВБШЫУЯНГЛЬУУ")
	print("\n")
}
