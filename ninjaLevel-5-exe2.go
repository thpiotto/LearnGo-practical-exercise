package main

import "fmt"

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
	Demonstre os valores do map utilizando range.
	Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

type Pessoa struct {
	Nome          string
	Sobrenome     string
	SaborSorverte []string
}

func main() {
	listaPessoas := make(map[string]Pessoa)

	listaPessoas["Vitória"] = Pessoa{
		Nome:          "Maria",
		Sobrenome:     "Vitória",
		SaborSorverte: []string{"Açaí", "Chocolate Branco"},
	}

	listaPessoas["Lira"] = Pessoa{
		Nome:          "Thales",
		Sobrenome:     "Lira",
		SaborSorverte: []string{"Uva", "Limão"},
	}

	for _, pessoa := range listaPessoas {
		fmt.Printf("Meu nome é %v e meus sorvetes favoritos são: ", pessoa.Nome)
		for _, sabor := range pessoa.SaborSorverte {
			fmt.Printf("\n - %v", sabor)
		}
		fmt.Println("\n")
	}
}
