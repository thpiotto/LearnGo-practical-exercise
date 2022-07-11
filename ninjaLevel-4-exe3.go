package main

import "fmt"

/*
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
	Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	Do quinto ao último item do slice (incluindo o último item!)
	Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)

	Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/

func main() {
	slice := []int{365, 2, 744, -14, 0, 77, 32, 1, 15, -88}
	fmt.Printf("Do primeiro ao terceiro item, incluso: %v\n", slice[:3])
	fmt.Printf("Do quinto ao ultimo item, incluso: %v\n", slice[4:])
	fmt.Printf("Do segundo ao sétimo item, incluso: %v\n", slice[1:7])
	fmt.Printf("Do terceiro ao penultimo item, incluso, sem len(): %v\n", slice[2:9])
	fmt.Printf("Do terceiro ao penultimo item, incluso, com len(): %v", slice[2:len(slice)-1])
	// for index, value := range slice {
	// 	fmt.Printf("i[%v]: v[%v]\n", index, value)
	// }
	fmt.Printf("\nO tipo do array demonstrado: %T", slice)
}
