package main

import (
	"fmt"
	"math/big"
)

func main() {
	b := big.NewInt(101)
	e := big.NewInt(17)
	m := big.NewInt(22663)

	result := new(big.Int)
	result.Exp(b, e, m)
	fmt.Printf("%s", result)
}
