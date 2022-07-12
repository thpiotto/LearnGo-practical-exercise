package main

import "fmt"

/*
- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
*/

func main() {

	pessoa := map[string]string{
		"nome":       "Thales",
		"sobre_nome": "Lira",
		"hobbies":    "code, animes",
	}

	for key, value := range pessoa {
		fmt.Printf("%v: %v\n", key, value)
	}

	pessoa["formação"] = "cientista da computação"
	fmt.Println()
	for key, value := range pessoa {
		fmt.Printf("%v: %v\n", key, value)
	}

}
