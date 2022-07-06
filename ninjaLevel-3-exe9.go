package main

import "fmt"

/*
 - Crie um programa que utilize a declaração switch, 
 onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

func main() {
	esporteFavorito := "MMA"

	switch esporteFavorito {
	case "Basquete":
		fmt.Println("Esporte favorito é o Basquete.")
	case "Futebol":
		fmt.Println("Esporte favorito é o Futebol.")
	case "MMA":
		fmt.Println("Esporte favorito é o MMA.")
	default:
		fmt.Println("Sem esporte favorito")
	}
}
