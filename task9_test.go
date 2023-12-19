package main

import (
	"fmt"
	"testing"
)

func TestTask9KeySelection(t *testing.T) {
	fmt.Printf("\n\nЗадача 8\nРасшифровка методом перестановки с фиксированным периодом")

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
	if !compareSlice(key, validKey) {
		t.Errorf("Неверный результат расшифровки, ожидалось %v, получено: %v", validKey, key)
	}
}

func compareSlice[T interface{ comparable }](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}
