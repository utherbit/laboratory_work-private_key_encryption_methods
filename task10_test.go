package main

import (
	"fmt"
	"testing"
)

func TestTask10EncryptMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 10\nШифрование сообщения, методом перестоновки по таблице")
	checkValidTask10EncryptMessage(t, task10EncryptMessage, "ШИРОКОПОЛОСНЫЙ_УСИЛИТЕЛЬ", []int{4, 1, 2, 3, 5}, "ОЛЙЛЬШОСУТИПНСЕРОЫИЛКО_И")
	checkValidTask10EncryptMessage(t, task10EncryptMessage, "ПЕРЕДАЧА_ИЗОБРАЖЕНИЯ", []int{2, 4, 5, 1, 3}, "ЕЧОЕЕ_РИДИАЯПАЗЖРАБН")

	// Контрольная проверка по значениям из 11 задачи
	checkValidTask10EncryptMessage(t, task10EncryptMessage, "ПЕРЕДАЧА_ЗВУКА", []int{4, 1, 2, 3}, "ЕАУПД_КЕАЗАРЧВ")
	checkValidTask10EncryptMessage(t, task10EncryptMessage, "СЛАБЫЙ_СИГНАЛ", []int{3, 1, 4, 2}, "А_НСЫИЛБСАЛЙГ")

	print("\n")
}

func checkValidTask10EncryptMessage(t *testing.T, f func(e string, key []int) string,
	e string, key []int, validDecrypted string) {

	decrypted := f(e, key)
	fmt.Printf("\nввод: {e: %s key: %v} вывод: %s", e, key, decrypted)
	if validDecrypted != decrypted {
		t.Errorf("Неверный результат шифрования, ожидалось %v, получено: %v", validDecrypted, decrypted)
	}
}
