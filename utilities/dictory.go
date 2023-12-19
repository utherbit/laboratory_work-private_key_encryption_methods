package utilities

import (
	"fmt"
	"testing"
)

func CheckValidEncryption(t *testing.T, encryptionFunc func(i string) string, inp, valid string) {
	result := encryptionFunc(inp)
	fmt.Printf("\nввод: '%s' вывод: '%s';", inp, result)
	if valid != result {
		t.Errorf("Неверный результат расшифровки, ожидалось %s, получено: %s", valid, result)
	}
}
