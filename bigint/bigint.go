package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	var amount = "0x0234c8a3397aab58" // 158972490234375000
	pi, ok := ParseBig256(amount)
	if ok {
		str, err := ToString(pi)
		fmt.Println(str, err)
	}
}
func ToString(big *big.Int) (string, error) {
	if big == nil {
		return "0x0", nil
	}
	a := fmt.Sprintf("%d", big)

	fmt.Println(a, strings.Compare(a, "158972490234375000"))
	return fmt.Sprintf("%#x", big), nil
}

func ParseBig256(s string) (*big.Int, bool) {
	if s == "" {
		return new(big.Int), true
	}
	var bigint *big.Int
	var ok bool
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		bigint, ok = new(big.Int).SetString(s[2:], 16)
	} else {
		bigint, ok = new(big.Int).SetString(s, 10)
	}
	if ok && bigint.BitLen() > 256 {
		bigint, ok = nil, false
	}
	return bigint, ok
}
