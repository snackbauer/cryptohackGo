package main

import "fmt"

func main() {
	numbers := [23]int{99, 114, 121, 112, 116, 111, 123, 65, 83, 67, 73, 73, 95, 112, 114, 49, 110, 116, 52, 98, 108, 51, 125}
	for i := 0; i <= 22; i++ {
		fmt.Print(string(numbers[i]))
	}
	fmt.Println("")
}