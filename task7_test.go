package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask7encryptMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 7\nШифрование методом перестановки с фиксированным периодом")
	errHandleDecryptionF1 := func(s string) string {
		message, err := task7encryptMessage([]int{4, 3, 6, 2, 1, 5}, s)
		if err != nil {
			t.Error(err)
		}
		return message
	}
	// функция для проверки работы алгоритма, по заведомо известным значениям.
	errHandleDecryptionF2 := func(s string) string {
		message, err := task7encryptMessage([]int{6, 4, 2, 7, 5, 8, 1, 3}, s)
		if err != nil {
			t.Error(err)
		}
		return message
	}

	utilities.CheckValidEncryption(t, errHandleDecryptionF1, "ЖЕЛТЫЙ_ОГОНЬ", "ТЛЙЕЖЫОГЬО_Н")
	utilities.CheckValidEncryption(t, errHandleDecryptionF1, "МЫ_НАСТУПАЕМ", "Н_СЫМААПМУТЕ")

	fmt.Printf("\nПроверка шифрования, с заведомо известными значениями")
	utilities.CheckValidEncryption(t, errHandleDecryptionF2, "АПЕЛЬСИН", "СЛПИЬНАЕ")
	utilities.CheckValidEncryption(t, errHandleDecryptionF2, "ВИНОГРАД", "РОИАГДВН")
	print("\n")
}
