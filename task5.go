package main

import (
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
)

/*
5.	Первый байт фрагмента текста в шестнадцатеричном виде имеет вид А5.
На него накладывается по модулю два 4-х битовая гамма 0111 (в двоичном виде).
Что получится после шифрования?
*/

// task5gammaEncrypt
// inputBase - система счисления для вводного текста
// outputBase - система счисления для выводного текста
// input - текст для шифрования
// gamma - гамма  шифрования
func task5gammaEncrypt(inputBase, gammaBase, outputBase int, input string, gamma string) string {
	binInput := utilities.StrToBitArr(input, inputBase)
	binGamma := utilities.StrToBitArr(gamma, gammaBase)

	res := make([]int, len(binInput))
	for i := 0; i < len(binInput); i++ {
		res[i] = binInput[i] ^ binGamma[i%len(binGamma)]
	}

	return utilities.BitArrToStr(res, outputBase)
}
