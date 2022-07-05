package main

import (
	"fmt"
	"math"
)

/*
 - Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis:
*/

var expresion float64
var num uint8

func main() {
	expresion = 20*math.Sqrt(144) + 15
	num = uint8(expresion) //255
	fmt.Printf("DECIMAL\tBINARY\t\tHEX\n%v\t\t%#b\t%#x", num, num, num)
}
