package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func xorBytes(a, b []byte) []byte {
	if len(a) != len(b) {
		log.Fatalf("byte slices must be the same length")
	}
	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func xorHex(a, b string) string {
	bytesA, err := hex.DecodeString(a)
	bytesB, err := hex.DecodeString(b)
	if err != nil {
		log.Fatalf("Error decoding hex as string: %v", err)
	}

	result := xorBytes(bytesA, bytesB)
	return hex.EncodeToString(result)
}

func main() {
	ct := "04ee9855208a2cd59091d04767ae47963170d1660df7f56f5faf"
	key2and3 := "c1545756687e7573db23aa1c3452a098b71a7fbf0fddddde5fc1"
	key2and1 := "37dcb292030faa90d07eec17e3b1c6d8daf94c35d4c9191a5e1e"
	key1 := "a6c8b6733c9b22de7bc0253266a3867df55acde8635e19c73313"

	key2 := xorHex(key2and1, key1)
	key3 := xorHex(key2and3, key2)

	decrypted := xorHex(xorHex(xorHex(ct, key1), key2), key3)
	d, err := hex.DecodeString(decrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}
