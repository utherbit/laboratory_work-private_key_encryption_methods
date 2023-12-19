package main

import (
	"fmt"
	"strings"
)

// Словарь дешифрации для шифра 1
var decryptionDict1 = map[rune]rune{
	'В': 'А',
	'И': 'Б',
	'О': 'В',
	'А': 'Г',
	'Щ': 'Д',
	'П': 'Е',
	'К': 'Ж',
	'Б': 'З',
	'Ъ': 'И',
	' ': 'К',
	'Р': 'Л',
	'Т': 'М',
	'Ц': 'Н',
	'.': 'О',
	'Ж': 'П',
	'Г': 'Р',
	'Л': 'С',
	'Х': 'Т',
	'С': 'У',
	'Ь': 'Ф',
	'Ч': 'Х',
	'З': 'Ц',
	'М': 'Ч',
	'У': 'Ш',
	'Д': 'Щ',
	'Э': 'Ъ',
	'Н': 'Ы',
	'Ю': 'Ь',
	'Ы': 'Э',
	'Ш': 'Ю',
	'Е': 'Я',
	'Ф': ' ',
	'Я': '.',
}

func task1decipherMessage(inp string) string {
	decryptedChars := make([]string, len(inp))
	for idx, char := range inp {
		decryptedChar, ok := decryptionDict1[char]
		if !ok {
			fmt.Printf("\nWARNING! Не удалось расшифровать символ '%v' (%s)", char, string(char))
			decryptedChar = '_'
		}
		decryptedChars[idx] = string(decryptedChar)
	}

	return strings.Join(decryptedChars, "")
}
