package main

import "strings"

var task3vigenereDict = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
var task3vigenereKey = []rune("ЯБЛОКО")

func task3decipherMessage(inp string) string {

	dictRunes := []rune(task3vigenereDict)
	// Сдвиг для символа из ключа task4vigenereKey
	dictByKeyChar := make(map[rune][]rune)
	for _, keyChar := range task3vigenereKey {
		dictIdx := 0
		// Поиск индекса руны в словаре -> dictIdx
		for i, dictChar := range dictRunes {
			if dictChar == keyChar {
				dictIdx = i
				break
			}
		}

		dictByKeyChar[keyChar] = append(dictRunes[dictIdx:], dictRunes[:dictIdx]...)
	}

	decryptedChars := make([]string, len(inp))
	for idx, char := range []rune(inp) {
		idxKey := idx % len(task3vigenereKey)
		keyChar := task3vigenereKey[idxKey]
		dict := dictByKeyChar[keyChar]

		dictIdx := 0
		// Поиск индекса руны в словаре -> dictIdx
		for i, dictChar := range dictRunes {
			if dictChar == char {
				dictIdx = i
				break
			}
		}

		dictChar := dict[dictIdx]
		decryptedChars[idx] = string(dictChar)
	}

	return strings.Join(decryptedChars, "")
}
