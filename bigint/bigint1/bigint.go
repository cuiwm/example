package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func ToBestStr(pi *big.Int) (result string) {
	z := big.NewInt(0)
	z1 := big.NewInt(0)
	m := big.NewInt(0)

	y := big.NewInt(params.Ether)
	y1 := big.NewInt(params.Finney)

	z.DivMod(pi, y, m)
	z1.Div(m, y1)

	return fmt.Sprintf("%d.%03d", z, z1)
}

func RoundToBestStr(pi *big.Int) (result string) {
	z := big.NewInt(0)
	z1 := big.NewInt(0)
	m := big.NewInt(0)
	r := big.NewInt(0)

	y := big.NewInt(params.Ether)
	y1 := big.NewInt(params.Finney)
	hf := big.NewInt(params.Finney)
	hf.Div(hf, big.NewInt(2))

	z.Set(pi)
	r.Set(pi)
	z.Add(z, hf)
	z.DivMod(z, y, m)
	//m.Add(m, hf)
	z1.Div(m, y1)

	return fmt.Sprintf("%d.%03d vs %04d", z, z1, big.NewInt(0).Div(r, big.NewInt(1e14)))
}

func main() {
	//z := big.NewInt(params.Ether)
	//y := big.NewInt(params.Finney)
	//x := big.NewInt(0)

	i := 0
	k := big.NewInt(0)
	for x := big.NewInt(params.Ether); ; x.Sub(x, big.NewInt(1e14)) {
		if x.Cmp(big.NewInt(1e14)) < 0 {
			fmt.Println("break ", x)
			break

		}
		fmt.Println(ToBestStr(x), " --- round:", i, "  x: ", x, k.Div(x, big.NewInt(1e14)))
		i++
	}
}
