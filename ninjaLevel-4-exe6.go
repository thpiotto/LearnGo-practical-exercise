package main

import "fmt"

/*
- Crie uma slice usando make que possa conter todos os estados do Brasil.

	Os estados:
	"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
	"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais",
	"Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
	"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
	"São Paulo", "Sergipe", "Tocantins"

- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice sem utilizar range.
*/

func main() {

	slice := make([]string, 26)
	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
		"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais",
		"Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Printf("Comprimento: %d, Capacidade: %d\n", len(slice), cap(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%s\n", slice[i])
	}
}
