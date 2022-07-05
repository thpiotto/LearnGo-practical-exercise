package main

import "fmt"

/*
 - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos
 - Demonstre estes valores
*/

const (
	_ = iota + 2022
	yrPlus1
	yrPlus2
	yrPlus3
	yrPlus4
)

func main() {
	fmt.Printf("Os próximos 4 anos:\n%v\n%v\n%v\n%v", yrPlus1, yrPlus2, yrPlus3, yrPlus4)
}
