package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask8decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 8\nРасшифровка методом перестановки с фиксированным периодом")
	errHandleDecryptionF := func(s string) string {
		message, err := task8decipherMessage([]int{6, 4, 2, 7, 5, 8, 1, 3}, s)
		if err != nil {
			t.Error(err)
		}
		return message
	}

	utilities.CheckValidEncryption(t, errHandleDecryptionF, "СЛПИЬНАЕ", "АПЕЛЬСИН")
	utilities.CheckValidEncryption(t, errHandleDecryptionF, "РОИАГДВН", "ВИНОГРАД")
	print("\n")
}
