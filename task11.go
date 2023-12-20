package main

/*
11.	Расшифруйте сообщения, зашифрованные методом перестановки по таблице 4*4
(символ подчеркивания заменяет пробел). Ключ указывает порядок считывания
столбцов при шифровании.
*/

func task11DecryptMessage(encryptedText string, key []int) string {
	encryptedTextRunes := []rune(encryptedText)
	keyLen := len(key)

	fullLines := len(encryptedTextRunes) / keyLen
	extraChars := len(encryptedTextRunes) % keyLen
	lines := fullLines
	if extraChars != 0 {
		lines += 1
	}

	plainTextRunes := make([]rune, lines*keyLen)

	encryptedIndex := 0
	for _, index := range key {

		for x := 0; x < lines; x++ {
			if x == fullLines && index-1 >= extraChars {
				continue
			}

			actualIndex := (x*keyLen + index) - 1
			if encryptedIndex < len(encryptedTextRunes) {
				plainTextRunes[actualIndex] = encryptedTextRunes[encryptedIndex]
				encryptedIndex++
			}
		}
	}

	plainTextRunes = plainTextRunes[:len(encryptedTextRunes)]
	return string(plainTextRunes)
}
