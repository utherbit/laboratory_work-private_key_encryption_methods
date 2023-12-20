package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"strings"
)

func task1HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 1\nРасшифровка сообщения по первому шифру из таблицы")

	fmt.Printf("\nВведите сообщение которое нужно расшифровать: ")
	message := utilities.Input()

	decrypted := task1decipherMessage(strings.TrimSpace(message))
	fmt.Printf("\nРезультат расшифровки: %s", decrypted)
}

func task2HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 2\nРасшифровка сообщения по второму шифру из таблицы")

	fmt.Printf("\nВведите сообщение которое нужно расшифровать: ")
	message := utilities.Input()

	decrypted := task2decipherMessage(strings.TrimSpace(message))
	fmt.Printf("\nРезультат расшифровки: %s", decrypted)
}
func task3HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 3\nШифрование сообщения по шифру Вижинера")

	fmt.Printf("\nВведите исходный алфавит: ")
	dict := utilities.Input()

	fmt.Printf("\nВведите сообщение которое нужно зашифровать: ")
	message := utilities.Input()

	fmt.Printf("\nВведите ключ который хотите использовать для шифрования: ")
	key := utilities.Input()

	encrypted := task3encryptMessage(strings.TrimSpace(dict), strings.TrimSpace(message), key)
	fmt.Printf("\nРезультат шифрования: %s", encrypted)
}
func task4HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 4\nРасшифрование сообщения по шифру Вижинера")

	fmt.Printf("\nВведите исходный алфавит: ")
	dict := utilities.Input()

	fmt.Printf("\nВведите сообщение которое нужно расшифровать: ")
	message := utilities.Input()

	fmt.Printf("\nВведите ключ который хотите использовать для расшифрования: ")
	key := utilities.Input()

	decrypted := task4decipherMessage(dict, message, key)
	fmt.Printf("\nРезультат расшифрования: %s", decrypted)
}
func task5HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 5\nШифрование при помощи метода гаммирования")

	fmt.Printf("\nВведите число - в какой системе счисление хотите ввести сообщение для шифрования:")
	messageBase := utilities.InputInt()

	fmt.Printf("\nВведите сообщение в %d-ричной системе счисления, которое хотите зашифровать: ", messageBase)
	message := utilities.Input()

	fmt.Printf("\nВведите число - в какой системе счисление хотите ввести гамму для шифрования: ")
	gammaBase := utilities.InputInt()

	fmt.Printf("\nВведите гамму в %d-ричной системе, для шифрования: ", gammaBase)
	gamma := utilities.Input()

	fmt.Printf("\nВведите число - в какой системе счисление хотите получить результат шифрования: ")
	outputBase := utilities.InputInt()

	encrypted := task5gammaEncrypt(messageBase, gammaBase, outputBase, strings.TrimSpace(message), strings.TrimSpace(gamma))
	fmt.Printf("\nРезультат шифрования: %s", encrypted)
}
func task6HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 6\nПолучение ключа гаммирования")

	fmt.Printf("\nВведите число - в какой системе счисление хотите ввести сообщение до шифрования: ")
	baseBefore := utilities.InputInt()

	fmt.Printf("\nВведите сообщение до шифрования в %d-ричной системе: ", baseBefore)
	messageBefore := utilities.Input()

	fmt.Printf("\nВведите число - в какой системе счисление хотите ввести сообщение после шифрования: ")
	baseAfter := utilities.InputInt()

	fmt.Printf("\nВведите сообщение после шифрования в %d-ричной системе: ", baseAfter)
	messageAfter := utilities.Input()

	gamma, err := task6FindGamma(strings.TrimSpace(messageBefore), strings.TrimSpace(messageAfter), baseBefore, baseAfter)
	if err != nil {
		fmt.Printf("\nПроизошла ошибка при попытке получить ключ гаммирования. Ошибка: %v", err)
		return
	}
	fmt.Printf("\nКлюч гаммирования: %v", gamma)
}

func task7HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 7\nШифрование методом перестновки с фиксированным периодом")

	fmt.Printf("\nВведите число - период шифрования: ")
	d := utilities.InputInt()

	fmt.Printf("\nВведите ключ для шифрования, для разделения значений, используйте пробел: ")
	key := make([]int, d)
	{
		keyStr := utilities.Input()

		keyStrValues := strings.Split(keyStr, " ")
		fmt.Printf("\nKeyStrvalues: %v", keyStrValues)
		if len(keyStrValues) != d {
			fmt.Printf("\nНеверное количество значений в ключе. Ожидалось %d, получено %d", d, len(keyStrValues))
			return
		}
		for i := 0; i < d; i++ {
			_, err := fmt.Sscanf(keyStrValues[i], "%d", &key[i])
			if err != nil {
				fmt.Printf("\nПроизошла ошибка при попытке сканировать введенное значение. Ошибка: %v", err)
				return
			}
		}
	}

	fmt.Printf("\nВведите сообщение которое нужно зашифровать: ")
	message := utilities.Input()

	encrypted, err := task7encryptMessage(key, strings.TrimSpace(message))
	if err != nil {
		fmt.Printf("\nПроизошла ошибка при попытке зашифровать сообщение. Ошибка: %v", err)
		return
	}
	fmt.Printf("\nРезультат шифрования: %s", encrypted)
}

func task8HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 8\nРасшифрование сообщения, зашифрованного методом перестновки с фиксированным периодом")

	fmt.Printf("\nВведите число - период шифрования: ")
	d := utilities.InputInt()

	fmt.Printf("\nВведите ключ для шифрования, для разделения значений, используйте пробел: ")
	key := make([]int, d)
	{
		keyStr := utilities.Input()

		keyStrValues := strings.Split(keyStr, " ")
		fmt.Printf("\nKeyStrvalues: %v", keyStrValues)
		if len(keyStrValues) != d {
			fmt.Printf("\nНеверное количество значений в ключе. Ожидалось %d, получено %d", d, len(keyStrValues))
			return
		}
		for i := 0; i < d; i++ {
			_, err := fmt.Sscanf(keyStrValues[i], "%d", &key[i])
			if err != nil {
				fmt.Printf("\nПроизошла ошибка при попытке сканировать введенное значение. Ошибка: %v", err)
				return
			}
		}
	}

	fmt.Printf("\nВведите сообщение которое нужно расшифровать: ")
	message := utilities.Input()

	decrypted, err := task8decipherMessage(key, strings.TrimSpace(message))
	if err != nil {
		fmt.Printf("\nПроизошла ошибка при попытке зашифровать сообщение. Ошибка: %v", err)
		return
	}
	fmt.Printf("\nРезультат расшифрования: %s", decrypted)
}

func task9HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 9\nОпределение ключа шифрования методом перестновки с фиксированным периодом")

	fmt.Printf("\nВведите сообщение до шифрования: ")
	messageBefore := utilities.Input()

	fmt.Printf("\nВведите сообщение после шифрования: ")
	messageAfter := utilities.Input()

	fmt.Printf("\nВведите число - период шифрования: ")
	d := utilities.InputInt()

	key := task9KeySelection(strings.TrimSpace(messageBefore), strings.TrimSpace(messageAfter), d)
	fmt.Printf("\nКлюч, который был использован для шифрования: %v", key)
}

func task10HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 10\nШифорвание сообщения методом перестановки по таблице")

	fmt.Printf("\nВведите сообщение которое хотите зашифровать: ")
	message := utilities.Input()

	fmt.Printf("\nВведите число - размерность матрицы шифрования: ")
	d := utilities.InputInt()

	fmt.Printf("\nВведите ключ для шифрования, для разделения значений, используйте пробел: ")
	key := make([]int, d)
	{
		keyStr := utilities.Input()

		keyStrValues := strings.Split(keyStr, " ")
		fmt.Printf("\nKeyStrvalues: %v", keyStrValues)
		if len(keyStrValues) != d {
			fmt.Printf("\nНеверное количество значений в ключе. Ожидалось %d, получено %d", d, len(keyStrValues))
			return
		}
		for i := 0; i < d; i++ {
			_, err := fmt.Sscanf(keyStrValues[i], "%d", &key[i])
			if err != nil {
				fmt.Printf("\nПроизошла ошибка при попытке сканировать введенное значение. Ошибка: %v", err)
				return
			}
		}
	}

	encrypted := task10EncryptMessage(strings.TrimSpace(message), key)
	fmt.Printf("\nЗашифрованное сообщение: %s", encrypted)
}

func task11HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 11\nРасшифрование сообщения, зашифрованного методом перестановки по таблице")

	fmt.Printf("\nВведите сообщение которое хотите расшифровать: ")
	message := utilities.Input()

	fmt.Printf("\nВведите число - размерность матрицы шифрования: ")
	d := utilities.InputInt()

	fmt.Printf("\nВведите ключ для расшифрования, для разделения значений, используйте пробел: ")
	key := make([]int, d)
	{
		keyStr := utilities.Input()

		keyStrValues := strings.Split(keyStr, " ")
		fmt.Printf("\nKeyStrvalues: %v", keyStrValues)
		if len(keyStrValues) != d {
			fmt.Printf("\nНеверное количество значений в ключе. Ожидалось %d, получено %d", d, len(keyStrValues))
			return
		}
		for i := 0; i < d; i++ {
			_, err := fmt.Sscanf(keyStrValues[i], "%d", &key[i])
			if err != nil {
				fmt.Printf("\nПроизошла ошибка при попытке сканировать введенное значение. Ошибка: %v", err)
				return
			}
		}
	}

	decrypted := task11DecryptMessage(strings.TrimSpace(message), key)
	fmt.Printf("\nРезультат расшифрования: %s", decrypted)
}

func task12HandlerConsoleInput() {
	fmt.Printf("\nПрактическое задание 12\nРасшифрование сообщения, зашифрованного методом пропорциональной замены")

	fmt.Printf("\nВведите сообщение которое хотите расшифровать: ")
	message := utilities.Input()

	decrypted := task12decipherMessage(strings.TrimSpace(message))
	fmt.Printf("\nРезультат расшифрования: %s", decrypted)
}
