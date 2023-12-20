package main

import (
	"fmt"
	"testing"
)

func TestTask11DecryptMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 11\nРасшифровка сообщения, зашифрованного методом перестоновки по таблице")
	checkValidTask11DecryptMessage(t, task11DecryptMessage, "ЕАУПД_КЕАЗАРЧВ", []int{4, 1, 2, 3}, "ПЕРЕДАЧА_ЗВУКА")
	checkValidTask11DecryptMessage(t, task11DecryptMessage, "А_НСЫИЛБСАЛЙГ", []int{3, 1, 4, 2}, "СЛАБЫЙ_СИГНАЛ")

	// Контрольная проверка по значениям из 10 задачи
	checkValidTask11DecryptMessage(t, task11DecryptMessage, "ОЛЙЛЬШОСУТИПНСЕРОЫИЛКО_И", []int{4, 1, 2, 3, 5}, "ШИРОКОПОЛОСНЫЙ_УСИЛИТЕЛЬ")
	checkValidTask11DecryptMessage(t, task11DecryptMessage, "ЕЧОЕЕ_РИДИАЯПАЗЖРАБН", []int{2, 4, 5, 1, 3}, "ПЕРЕДАЧА_ИЗОБРАЖЕНИЯ")
	print("\n")
}

func checkValidTask11DecryptMessage(t *testing.T, f func(e string, key []int) string,
	e string, key []int, validDecrypted string) {

	decrypted := f(e, key)
	fmt.Printf("\nввод: {e: %s key: %v} вывод: %s", e, key, decrypted)
	if validDecrypted != decrypted {
		t.Errorf("Неверный результат расшифровки, ожидалось %v, получено: %v", validDecrypted, decrypted)
	}
}
