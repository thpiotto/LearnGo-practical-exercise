package main

import "fmt"

/*
	- Desafio surpresa:
	- Format printing:
		- Decimal     %d
		- Hexadecimal %#x
		- Unicode     %#U
		- Tab         \t
		- Linha nova  \n

	- Faça um loop dos números 33 a 122, utilize format printing para demonstr=a-los como string
*/

func main() {
	fmt.Println("DEC\t   HEX\t\tUNICODE")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t   %#x\t\t%c\n", i, i, i)
	}
}
