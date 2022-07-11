package main

import "fmt"

/*
- Começando com a seguinte slice:
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	Anexe a ela o valor 52;
	Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
	Demonstre a slice;

- Anexe a ela a seguinte slice:
	y := []int{56, 57, 58, 59, 60}
	Demonstre a slice x.
*/

func main() {
	fmt.Print("> x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}\n")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("slice x: %v\n", x)
	fmt.Print("> x = append(x, 52)\n")
	x = append(x, 52)
	fmt.Printf("slice x: %v\n", x)
	fmt.Print("> x = append(x, 53, 54, 55)\n")
	x = append(x, 53, 54, 55)
	fmt.Printf("slice x: %v\n", x)
	fmt.Print("\n> y := []int{56, 57, 58, 59, 60}\n")
	y := []int{56, 57, 58, 59, 60}
	fmt.Printf("slice y: %v\n", y)
	fmt.Print("> x = append(x, y...)\n")
	x = append(x, y...)
	fmt.Printf("slice x: %v\n", x)
}
