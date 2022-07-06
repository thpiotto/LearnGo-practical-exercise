package main

import "fmt"

/*
 - Anote (à mão) o resultado das expressões:
   1. fmt.Println(true && true)
   2. fmt.Println(true && false)
   3. fmt.Println(true || true)
   4. fmt.Println(true || false)
   5. fmt.Println(!true)
*/

func main() {
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
