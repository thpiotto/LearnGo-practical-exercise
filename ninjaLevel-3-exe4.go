package main

import (
	"fmt"
	"time"
)

/*
 - Crie um loop utilizando a sintaxe: for {}
 - Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main() {

	initYr := 1990

	for {
		if initYr <= time.Now().Year() {
			fmt.Println(initYr)
			initYr++
		} else {
			break
		}
	}
}
