package main

import "fmt"

/*
 - Usando uma literal composta:
	Crie uma slice de tipo int
	Atribua 10 valores a ela
	Utilize range para demonstrar todos estes valores.
	E utilize format printing para demonstrar seu tipo.
*/

func main() {
	slice := []int{365, 2, 744, -14, 0, 77, 32, 1, 15, -88}
	for index, value := range slice {
		fmt.Printf("i[%v]: v[%v]\n", index, value)
	}
	fmt.Printf("\nO tipo do array demonstrado: %T", slice)
}
