package main

import "fmt"

/*
- Crie um map com key tipo string e value tipo []string.

    Key deve conter nomes no formato sobrenome_nome
    Value deve conter os hobbies favoritos da pessoa

	Demonstre todos esses valores e seus índices.
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

}
