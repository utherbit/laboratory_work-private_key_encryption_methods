package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func Max(slice []int) int {
	v := math.MinInt
	for _, i := range slice {
		if v < i {
			v = i
		}
	}
	return v
}

func task8decipherMessage(key []int, encrypted string) (string, error) {
	// Проверка валидности ключа
	maxKey := Max(key)
	if maxKey != len(key) {
		return "", errors.New(fmt.Sprintf("ERROR: Длинна ключа %d не равна максимальному значению ключа %d\n", len(key), maxKey))
	}
	validateKeyMap := make(map[int]*struct{})
	for _, v := range key {
		if _, ok := validateKeyMap[v]; ok {
			return "", errors.New(fmt.Sprintf("ERROR: Ключ %d повторяется\n", v))
		}
		validateKeyMap[v] = nil
	}

	// Создаю чанки, для заполнения значениями
	encryptedRunes := []rune(encrypted)
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
		idxKey := key[idx%d]
		chunks[idx/d][idxKey-1] = string(char)
	}

	joinChunks := make([]string, countChunks)
	for i, chunk := range chunks {
		joinChunks[i] = strings.Join(chunk, "")
	}
	result := strings.Join(joinChunks, "")

	return result, nil
}
