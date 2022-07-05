package main

import (
	"fmt"
)

/*
 - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/

func main() {
	fmt.Println("Resto da divisão de:")
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v mod 4 = %v\n", i, i%4)
	}
}
