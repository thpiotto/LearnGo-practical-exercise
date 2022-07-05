package main

import "fmt"

/*
 - Utilizando a solução anterior, adicione as opções else if e else.
*/

func main() {
	var num uint8
	num = 255
	if num == 0 {
		fmt.Println("Ocorreu um overflow")
	} else {
		fmt.Println("Cuidado pra não passar da limitação de memória")
	}
}
