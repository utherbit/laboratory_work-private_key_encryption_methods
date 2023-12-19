package main

import (
	"fmt"
	"strings"
)

// Словарь дешифрации для шифра 2
var decryptionDict2 = map[rune]rune{
	'^': 'А',
	'@': 'Б',
	')': 'В',
	'+': 'Г',
	'>': 'Е',
	'<': 'Д',
	'∀': 'Ж',
	'♦': 'З',
	'*': 'И',
	'♥': 'К',
	'♠': 'Л',
	'№': 'М',
	'#': 'Н',
	'-': 'О',
	'=': 'П',
	'(': 'Р',
	'?': 'С',
	'%': 'Т',
	'⦻': 'У',
	'!': 'Ф',
	//'№': 'Х',
	'®': 'Ц',
	'∑': 'Ч',
	'∇': 'Ш',
	'ϒ': 'Щ',
	'א': 'Ъ',
	'⊕': 'Ы',
	'×': 'Ь',
	'ω': 'Э',
	'$': 'Ю',
	'△': 'Я',
	'∞': ' ',
	'♣': '.',
}

func task2decipherMessage(inp string) string {
	decryptedChars := make([]string, len(inp))
	for idx, char := range inp {
		decryptedChar, ok := decryptionDict2[char]
		if !ok {
			fmt.Printf("\nWARNING! Не удалось расшифровать символ '%v'", char)
			decryptedChar = '_'
		}
		decryptedChars[idx] = string(decryptedChar)
	}

	return strings.Join(decryptedChars, "")
}
