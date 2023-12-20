package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask3decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 3\nШифрование сообщения по шифру Вижинера")
	decipherFunc := func(message string) string {
		return task3encryptMessage("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", message, "ЯБЛОКО")
	}

	utilities.CheckValidEncryption(t, decipherFunc, "КРИПТОСТОЙКОСТЬ", "ЙСФЮЭЭРУЪШХЭРУЗ")
	utilities.CheckValidEncryption(t, decipherFunc, "ГАММИРОВАНИЕ", "ВБШЫУЯНГЛЬУУ")

	// Контрольная проверка по значениям из 4 задания
	decipher2Func := func(message string) string {
		return task3encryptMessage("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ_", message, "ОРЕХ")
	}

	utilities.CheckValidEncryption(t, decipher2Func, "КУЗНЕЧИК", "ШВМБУЖНЯ")
	utilities.CheckValidEncryption(t, decipher2Func, "СТРЕКОЗА", "ЯБХЪШЮМХ")
	print("\n")
}
