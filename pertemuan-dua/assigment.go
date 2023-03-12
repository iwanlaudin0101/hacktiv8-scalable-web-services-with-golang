package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)

		if i == 4 {
			for j := 0; j <= 10; j++ {
				fmt.Println("Nilai j =", j)

				if j == 4 {
					character := []rune{'С','А','Ш','А','Р','В','О'}
					for i, char := range character {
						fmt.Printf("character %U ", char)
						fmt.Printf("'%s' starts at byte position %d\n", []byte(string(char)), i*2)
					}
				}
			}
		}
	}
}