package main

import (
	"math"
	"strconv"
)

/*
10.	Зашифруйте сообщения методом перестановки по таблице 5*5.
Ключ указывает порядок считывания столбцов при шифровании.
*/

func encrypt(plainText string, key string) string {
	plainTextRunes := []rune(plainText)
	keyLen := len(key)
	lines := int(math.Floor(float64(len(plainTextRunes))/float64(keyLen)) + 1)
	res := ""

	for _, n := range key {
		for x := 0; x < lines; x++ {
			index, err := strconv.Atoi(string(n))
			if err != nil {
				panic(err)
			}
			index += keyLen*x - 1

			if index > len(plainTextRunes)-1 {
				continue
			}

			res += string(plainTextRunes[index])
		}
	}
	return res
}
func task10EncryptMessage(message string, key []int) string {
	messageRunes := []rune(message)
	keyLen := len(key)
	lines := int(math.Floor(float64(len(messageRunes))/float64(keyLen)) + 1)
	res := ""

	for _, i := range key {
		for x := 0; x < lines; x++ {
			index := i
			index += keyLen*x - 1

			if index > len(messageRunes)-1 {
				continue
			}

			res += string(messageRunes[index])
		}
	}
	return res
}
