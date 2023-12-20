package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\nПроизошла ошибка при попытке сканировать введенное значение. Ошибка: %v", err)
		return ""
	}
	return strings.TrimSpace(strings.Replace(readString, "\n", "", -1))
}
func InputInt() int {
	message := Input()
	value, err := strconv.Atoi(message)
	if err != nil {
		fmt.Printf("\nПроизошла ошибка при попытке получить число из ввода. Ошибка: %v", err)
		return 0
	}
	return value
}
