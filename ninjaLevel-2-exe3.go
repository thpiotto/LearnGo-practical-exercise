package main

import (
	"fmt"
)

/*
 - Crie constantes tipadas e não-tipadas.
 - Demonstre seus valores.
*/

const utyped = 255
const typed uint8 = 255

func main() {
	fmt.Printf("utyped: %v\ntyped: %v", utyped, typed)
}
