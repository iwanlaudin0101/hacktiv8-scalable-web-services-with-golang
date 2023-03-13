package main

import (
	"fmt"
)

func main() {
	
	tokenize := func (str string) {
		for _, char := range str {
			fmt.Println(string(char))
		}
		// return ""
	}

	tokenize("selamat malam")
}