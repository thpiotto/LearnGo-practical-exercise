package main

import (
	"fmt"
	"time"
)

/*
 - Crie um loop utilizando a sintaxe: for condition {}
 - Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main() {

	initYr := 1990

	for initYr <= time.Now().Year() {
		fmt.Println(initYr)
		initYr++
	}
}
