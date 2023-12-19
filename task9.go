package main

/*
	9.	Определите ключи в системе шифрования, использующей
	перестановку с фиксированным периодом d=5 по парам открытых и зашифрованных сообщений:
*/

func task9KeySelection(encrypted, unencrypted string, d int) []int {

	unencryptedFirstD := []rune(unencrypted)[:d]
	encryptedFirstD := []rune(encrypted)[:d]

	key := make([]int, d)
	for iu, unencryptedRune := range unencryptedFirstD {

		findIdx := 0
		for ie, encryptedRune := range encryptedFirstD {
			if unencryptedRune == encryptedRune {
				findIdx = ie
				break
			}
		}
		key[iu] = findIdx + 1
	}
	// todo добавить проверку, по второй части фразы, подходит ли ключ (для проверки, правильности переданных параметров)
	// todo будет ломаться, в случае если встречаются повторяющиеся символы.. (в примере таких нет)
	return key
}
