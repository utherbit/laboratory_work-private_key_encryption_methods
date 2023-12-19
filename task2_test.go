package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask2decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 2\nРасшифровка сообщения по шифру 2")
	utilities.CheckValidEncryption(t, task2decipherMessage, "∇*!(∞♦№>#⊕", "ШИФР ЗМЕНЫ")
	utilities.CheckValidEncryption(t, task2decipherMessage, "@♠-♥∞∇*!(-)#*△", "БЛОК ШИФРОВНИЯ")
	print("\n")
}
