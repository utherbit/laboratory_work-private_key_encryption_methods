package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Printf("\nВведите номер задания: ")

		var taskNumber int
		_, err := fmt.Scanf("%d\n", &taskNumber)
		if err != nil {
			fmt.Printf("Некорректный номер задания! Ошибка: %v", err)
			continue
		}

		switch taskNumber {
		case 1:
			fmt.Printf("\nВведите сообщение для расшифровки: ")

			var input string
			_, err = fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("\nОшибка при сканировании ввода! Ошибка: %v", err)
			}

			result := task1decipherMessage(input)
			fmt.Printf("\nРезультат расшифровки: %s", result)

		case 2:
			fmt.Printf("\nВведите сообщение для расшифровки: ")

			var input string
			_, err = fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("\nОшибка при сканировании ввода! Ошибка: %v", err)
			}

			result := task1decipherMessage(input)
			fmt.Printf("\nРезультат расшифровки: %s", result)

		default:
			fmt.Printf("\nЗадание с номером %d не найдено!", taskNumber)
		}

		fmt.Printf("\n\n\n==========\n")
	}
}
