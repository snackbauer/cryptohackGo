package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	big, ok := new(big.Int).SetString("11515195063862318899931685488813747395775516287289682636499965282714637259206269", 10)
	if !ok {
		panic(ok)
	}
	strVer := fmt.Sprintf("%x", big)

	id, err := hex.DecodeString(strVer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(id))
}
