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

	n := new(big.Int)
	n, _ = n.SetString("882564595536224140639625987659416029426239230804614613279163", 10)

	p := new(big.Int)
	p, _ = p.SetString("857504083339712752489993810777", 10)

	q := new(big.Int)
	q, _ = q.SetString("1029224947942998075080348647219", 10)

	e := big.NewInt(65537)

	phiN := eulerTotient(p, q) // Calculate phi(n)

	d := new(big.Int)
	d.ModInverse(e, phiN) // Calculate d mod phi(n)

	c := new(big.Int)
	c, _ = c.SetString("77578995801157823671636298847186723593814843845525223303932", 10)

	m := new(big.Int).Exp(c, d, n)
	fmt.Printf("%v", m)
}
