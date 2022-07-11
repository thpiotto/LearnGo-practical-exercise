package main

import "fmt"

/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
	"Nome", "Sobrenome", "Hobby favorito"
	Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
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
