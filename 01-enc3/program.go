package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	ct := "72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf"

	id, err := hex.DecodeString(ct)
	d := base64.StdEncoding.EncodeToString(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))

}
