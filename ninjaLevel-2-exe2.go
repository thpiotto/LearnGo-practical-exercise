package main

import (
	"fmt"
	"math"
)

/*
 - Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis:
   1. ==
   2. !=
   3. >
   3. >=
   4. <
   5. <=

*/

func main() {
	expresion1 := (20*math.Sqrt(144) + 15) == 255 //true
	expresion2 := (20*math.Sqrt(144) + 15) != 255 //false
	expresion3 := (20*math.Sqrt(144) + 15) < 256  //true
	expresion4 := (20*math.Sqrt(144) + 15) <= 254 //false
	expresion5 := (20*math.Sqrt(144) + 15) > 254  //true
	expresion6 := (20*math.Sqrt(144) + 15) >= 256 //false

	fmt.Printf("%v, %v, %v, %v, %v, %v", expresion1, expresion2, expresion3, expresion4, expresion5, expresion6)
}
