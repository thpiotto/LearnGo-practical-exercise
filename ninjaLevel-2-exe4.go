package main

import (
	"fmt"
)

/*
 - Crie um programa que:
 	- Atribua um valor int a uma variável;
	- Demonstre este valor em decimal, binário e hexadecimal
	- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
	- Demonstre esta outra variável em decimal, binário e hexadecimal
*/

var x uint8 = 40
var y uint8

func main() {
	y = x << 1
	fmt.Printf("DECIMAL\tBINARY\t\tHEX\n%v\t\t%#b\t%#x\n", x, x, x)
	fmt.Printf("DECIMAL\tBINARY\t\tHEX\n%v\t\t%#b\t%#x", y, y, y)
}
