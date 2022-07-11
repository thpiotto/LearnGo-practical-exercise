package main

import "fmt"

/*
 - Usando uma literal composta:
	Crie um array que suporte 5 valores to tipo int
	Atribua valores aos seus índices
	Utilize range e demonstre os valores do array.
	Utilizando format printing, demonstre o tipo do array.
*/

func main() {
	arr := [5]int{365, 2, 744, -14, 0}
	for index, value := range arr {
		fmt.Printf("i[%v]: v[%v]\n", index, value)
	}
	fmt.Printf("\nO tipo do array demonstrado: %T", arr)
}
