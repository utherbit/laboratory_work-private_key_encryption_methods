package main

import "strings"

//var task4vigenereDict = "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ_"

//var task4vigenereKey = []rune("ОРЕХ")

func task4decipherMessage(dict, inp, key string) string {
	keyRunes := []rune(key)
	dictRunes := []rune(dict)
	// Сдвиг для символа из ключа task4vigenereKey
	dictByKeyChar := make(map[rune][]rune)
	for _, keyChar := range keyRunes {
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
		idxKey := idx % len(keyRunes)
		keyChar := keyRunes[idxKey]
		dict := dictByKeyChar[keyChar]

		dictIdx := 0
		// Поиск индекса руны в словаре -> dictIdx
		for i, dictChar := range dict {
			if dictChar == char {
				dictIdx = i
				break
			}
		}

		dictChar := dictRunes[dictIdx]
		decryptedChars[idx] = string(dictChar)
	}

	return strings.Join(decryptedChars, "")
}
