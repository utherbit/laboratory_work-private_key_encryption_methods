package main

import (
	"fmt"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"testing"
)

func TestTask1decipherMessage(t *testing.T) {
	fmt.Printf("\n\nЗадача 1\nРасшифровка сообщения по шифру 1")
	utilities.CheckValidEncryption(t, task1decipherMessage, "И.РЮУ.ЪФОБГНО", "БОЛЬШОИ ВЗРЫВ")
	utilities.CheckValidEncryption(t, task1decipherMessage, "СЛХГ.ЪЛХО.ФОО.ЩВ", "УСТРОИСТВО ВВОДА")
	print("\n")
}
