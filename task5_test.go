package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask5decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 5\nШифрование при помощи метода гаммирования")
	encryptFunc := func(inp string) string {
		return task5gammaEncrypt(16, 2, 2, inp, "0111")
	}

	utilities.CheckValidEncryption(t, encryptFunc, "a5", "11010010")

	fmt.Printf("\ntestTask6: %v", task5gammaEncrypt(16, 2, 16, "ffffffff", "01011111"))
	print("\n")
}
