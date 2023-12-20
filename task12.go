package main

import "strings"

/*
12.	Известно, что при использовании шифра пропорциональной замены
каждой русской букве поставлено в соответствие одно или несколько
трехзначных чисел по таблице замен:
*/

func task12decipherMessage(message string) string {
	messageRunes := []rune(message)
	decryptedChars := make([]string, 0)
	for i := 0; i < len(messageRunes); i += 3 {
		key := string(messageRunes[i]) + string(messageRunes[i+1]) + string(messageRunes[i+2])
		decryptedChar, ok := task12dictionary[key]
		if ok {
			decryptedChars = append(decryptedChars, string(decryptedChar))
		} else {
			decryptedChars = append(decryptedChars, " ")
		}
	}
	return strings.Join(decryptedChars, "")
}

// ключ - 3 значное число, значение - символ
var task12dictionary = map[string]rune{
	"760": 'А', "128": 'А', "350": 'А', "201": 'А', // А
	"101": 'Б',             // Б
	"210": 'В', "106": 'В', // В
	"351": 'Г',                                     // Г
	"129": 'Д',                                     // Д
	"761": 'Е', "130": 'Е', "802": 'Е', "352": 'Е', // Е
	"102": 'Ж',                         // Ж
	"753": 'З',                         // З
	"762": 'И', "211": 'И', "131": 'И', // И
	"754": 'К', "764": 'К', // К
	"132": 'Л', "354": 'Л', // Л
	"755": 'М', "742": 'М', // М
	"763": 'Н', "756": 'Н', "212": 'Н', // Н
	"757": 'О', "213": 'О', "765": 'О', "133": 'О', "353": 'О', // О
	"743": 'П', "766": 'П', // П
	"134": 'Р', "532": 'Р', // Р
	"800": 'С', "767": 'С', "105": 'С', // С
	"759": 'Т', "135": 'Т', "214": 'Т', // Т
	"544": 'У',             // У
	"560": 'Ф',             // Ф
	"768": 'Х',             // Х
	"545": 'Ц',             // Ц
	"215": 'Ч',             // Ч
	"103": 'Ш',             // Ш
	"752": 'Щ',             // Щ
	"561": 'Ъ',             // Ъ
	"136": 'Ы',             // Ы
	"562": 'Ь',             // Ь
	"750": 'Э',             // Э
	"570": 'Ю',             // Ю
	"216": 'Я', "104": 'Я', // Я
}