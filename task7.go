package main

import (
	"errors"
	"fmt"
	"strings"
)

func task7encryptMessage(key []int, input string) (string, error) {
	// Проверка валидности ключа
	maxKey := Max(key)
	if maxKey != len(key) {
		return "", errors.New(fmt.Sprintf("ERROR: Длинна ключа %d не равна максимальному значению ключа %d\n", len(key), maxKey))
	}
	// Ключ - сдвиг; значение - позиция
	keyMap := make(map[int]int)
	for i, v := range key {
		if _, ok := keyMap[v]; ok {
			return "", errors.New(fmt.Sprintf("ERROR: Значение ключа %d повторяется\n", v))
		}
		keyMap[v] = i
	}

	// Создаю чанки, для заполнения значениями
	encryptedRunes := []rune(input)
	d := len(key)
	countChunks := len(encryptedRunes) / d
	if len(encryptedRunes)%d > 0 {
		countChunks++
	}
	chunks := make([][]string, countChunks)
	for i := 0; i < countChunks; i++ {
		chunks[i] = make([]string, d)
	}

	for idx, char := range encryptedRunes {
		// Получить шифрующий индекс
		idxKey := keyMap[idx%d+1]
		// Заполнение чанков
		chunks[idx/d][idxKey] = string(char)
	}

	// Объединение чанков в строку
	joinChunks := make([]string, countChunks)
	for i, chunk := range chunks {
		joinChunks[i] = strings.Join(chunk, "")
	}
	result := strings.Join(joinChunks, "")

	return result, nil
}
