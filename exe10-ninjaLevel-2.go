package main

import "fmt"

/*
 - Crie uma variável de tipo string utilizando uma raw string literal.
 - Demonstre-a.
*/

var str string

func main() {
	str = `THIS IS
	
			A EXEMPLE
				
						OF A RAW STRING LITERAL`

	fmt.Print(str)
}
