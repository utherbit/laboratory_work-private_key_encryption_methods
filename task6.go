package main

import (
	"errors"
	"github.com/utherbit/laboratory_work-private_key_encryption_methods/utilities"
	"strconv"
	"strings"
)

func task6FindGamma(beforeEncrypt, afterEncrypt string, baseBefore, baseAfter int) ([]int, error) {
	if len([]rune(beforeEncrypt)) != len([]rune(afterEncrypt)) {
		return nil, errors.New("lengths are not equal")
	}

	binBefore := utilities.StrToBitArr(beforeEncrypt, baseBefore)
	binAfter := utilities.StrToBitArr(afterEncrypt, baseAfter)

	res := make([]int, 0)
	for i := 0; i < len(binBefore); i++ {
		res = append(res, binBefore[i]^binAfter[i])
	}

	var resStr string
	for _, r := range res {
		resStr += strconv.Itoa(r)
	}

	for countr := 2; true; countr++ {
		if len(resStr)%countr != 0 {
			continue
		} else {
			lenSub := len(resStr) / countr
			sub := resStr[:lenSub]
			if strings.Contains(resStr, sub) {
				return res[:lenSub], nil
			}
		}
	}

	return res, nil
}
