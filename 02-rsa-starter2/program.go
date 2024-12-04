package main

import (
	"fmt"
	"math/big"
)

func main() {
	pt := big.NewInt(12)
	e := big.NewInt(65537)

	p := big.NewInt(17)
	q := big.NewInt(23)

	n := new(big.Int).Mul(p, q)

	//fmt.Printf("%v, %v, %v, %v %v", pt, e, p, q, n)

	result := new(big.Int)
	result.Exp(pt, e, n)
	fmt.Printf("%s", result)
}
