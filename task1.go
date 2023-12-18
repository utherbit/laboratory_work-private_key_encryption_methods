package main

import (
	"fmt"
	"main/utilities"
)

func task1decipherMessage(inp string) string {
	decryptedChars := make([]rune, len(inp))
	for idx, char := range inp {
		decryptedChar, ok := utilities.DecryptionDict1[char]
		if !ok {
			fmt.Printf("\nWARNING! Не удалось расшифровать символ '%v'", char)
			decryptedChar = '_'
		}
		decryptedChars[idx] = decryptedChar
	}
	return string(decryptedChars)
}
