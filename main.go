package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Printf("\n===========================\nВведите номер задания: ")

		var taskNumber int
		_, err := fmt.Scanln(&taskNumber)
		if err != nil {
			fmt.Printf("\n1Произошла ошибка при попытке сканировать введенное значение. Ошибка: %v", err)
			continue
		}

		switch taskNumber {
		case 1:
			task1HandlerConsoleInput()

		case 2:
			task2HandlerConsoleInput()

		case 3:
			task3HandlerConsoleInput()

		case 4:
			task4HandlerConsoleInput()

		case 5:
			task5HandlerConsoleInput()

		case 6:
			task6HandlerConsoleInput()

		case 7:
			task7HandlerConsoleInput()

		case 8:
			task8HandlerConsoleInput()

		case 9:
			task9HandlerConsoleInput()

		case 10:
			task10HandlerConsoleInput()

		case 11:
			task11HandlerConsoleInput()

		case 12:
			task12HandlerConsoleInput()
		default:
			fmt.Printf("\nНе найдено задание с номером %d", taskNumber)
		}
	}
}
