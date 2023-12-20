package utilities

import (
	"fmt"
	"math/big"
	"strconv"
)

func BitArrToStr(in []int, base int) string {
	// Convert the bit array to a big.Int
	num := new(big.Int)
	for _, bit := range in {
		num.Lsh(num, 1)
		num.Add(num, big.NewInt(int64(bit)))
	}

	// Convert the big.Int to a string in the specified base
	return num.Text(base)
}
func StrToBitArr(in string, base int) (out []int) {
	if base > 2 {
		inter := ""
		for _, r := range in {
			app, _ := strconv.ParseInt(string(r), base, 8)
			inter += fmt.Sprintf("%.4b", app)
		}
		in = inter
	}
	for _, r := range in {
		app, _ := strconv.Atoi(string(r))
		out = append(out, app)
	}
	return
}
