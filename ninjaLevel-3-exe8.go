package main

import "fmt"

/*
 - Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

func main() {
	x := 2

	switch {
	case x <= 5:
		fmt.Println("Primeira metade dos 10 primeiros números inteiros acima de zero.")
	case x <= 10:
		fmt.Println("Segunda metade dos 10 primeiros números inteiros acima de zero.")
	default:
		fmt.Println("Acima de 10")
	}
}
