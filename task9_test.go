package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask9KeySelection(t *testing.T) {
	fmt.Printf("\n\nЗадача 9\nПодбор ключа для шифрования методом перестановки")

	// контрольная проверка
	checkValidT9KeySelection(t,
		task9KeySelection, "ВИНОГРАД", "РОИАГДВН", 8, []int{6, 4, 2, 7, 5, 8, 1, 3})

	checkValidT9KeySelection(t,
		task9KeySelection, "МОЙ ПАРОЛЬ", "ЙПМ ООЬАЛР", 5, []int{3, 5, 1, 4, 2})

	checkValidT9KeySelection(t,
		task9KeySelection, "СИГНАЛ БОЯ", "НИСАГО ЛЯБ", 5, []int{4, 2, 1, 5, 3})

}

func checkValidT9KeySelection(t *testing.T, f func(e string, u string, d int) []int,
	encrypted string, unencrypted string, d int, validKey []int) {

	key := f(encrypted, unencrypted, d)
	fmt.Printf("\nввод: {e:%s u:%s d:%d} вывод: %v", encrypted, unencrypted, d, key)
	if !utilities.CompareSlice(key, validKey) {
		t.Errorf("Неверный результат расшифровки, ожидалось %v, получено: %v", validKey, key)
	}
}
