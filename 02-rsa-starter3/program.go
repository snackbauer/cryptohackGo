package main

import (
	"fmt"
	"math/big"
)

func eulerTotient(p, q *big.Int) *big.Int {
	pMinus1 := new(big.Int).Sub(p, big.NewInt(1))
	qMinus1 := new(big.Int).Sub(q, big.NewInt(1))
	totient := new(big.Int).Mul(pMinus1, qMinus1)
	return totient
}

func main() {
	p := new(big.Int)
	p, _ = p.SetString("857504083339712752489993810777", 10)
	q := new(big.Int)
	q, _ = q.SetString("1029224947942998075080348647219", 10)

	fmt.Printf("%v", eulerTotient(p, q))
}
