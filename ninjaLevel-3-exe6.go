package main

import "fmt"

/*
 - Crie um programa que demonstre o funcionamento da declaração if.
*/

func main() {
	var num uint8
	num = 255
	if num+1 == 0 {
		fmt.Println("Ocorreu um overflow")
	}
}
