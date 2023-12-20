package utilities

import (
	"fmt"
	"testing"
)

func TestStrToBitArr(t *testing.T) {
	b := StrToBitArr("fff123abc", 16)
	fmt.Printf("\nS2B %v", b)
	a := BitArrToStr(b, 16)
	fmt.Printf("\nB2S %v", a)
}
