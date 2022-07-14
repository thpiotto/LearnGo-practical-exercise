package main

import "fmt"

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
	Nome
	Sobrenome
	Sabores favoritos de sorvete

- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/

type Pessoa struct {
	Nome          string
	Sobrenome     string
	SaborSorverte []string
}

func main() {
	pessoa1 := Pessoa{}
	pessoa1.Nome = "Thales"
	pessoa1.Sobrenome = "Lira"
	pessoa1.SaborSorverte = append(pessoa1.SaborSorverte, "Uva", "Limão")
	pessoa2 := Pessoa{}
	pessoa2.Nome = "Maria"
	pessoa2.Sobrenome = "Vitória"
	pessoa2.SaborSorverte = append(pessoa2.SaborSorverte, "Açaí")

	for _, sabores := range pessoa1.SaborSorverte {
		fmt.Printf("%v\t", sabores)
	}
	fmt.Println()
	for _, sabores := range pessoa2.SaborSorverte {
		fmt.Printf("%v\t", sabores)
	}
}
