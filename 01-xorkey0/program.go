package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

func xorWithByte(message []byte, key byte) []byte {
	result := make([]byte, len(message))
	for i := range message {
		result[i] = message[i] ^ key
	}
	return result
}

func worker(message []byte, key byte, wg *sync.WaitGroup) {
	defer wg.Done() // Signal the WaitGroup when the worker finishes

	xored := xorWithByte(message, key)
	cleanResult := strings.ReplaceAll(string(xored), "\n", "")
	fmt.Printf("Key: 0x%02x | Result: %s\n", key, cleanResult)
}

func main() {
	messageHex := "73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d"
	message, err := hex.DecodeString(messageHex)
	if err != nil {
		panic("Invalid hex string")
	}

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	for key := 0; key <= 0xFF; key++ {
		wg.Add(1)
		go worker(message, byte(key), &wg)
	}
	wg.Wait()
}
