package main

import "fmt"

func main() {
	var num1, num2 int
	var suma, dif, prod int
	var div int

	fmt.Print("ingrese primer numero", num1)
	fmt.Scan(&num1)

	fmt.Print("ingrese segundo numero", num2)
	fmt.Scan(&num2)

	//operaciones
	suma = num1 + num2
	dif = num1 - num2
	prod = num1 * num2
	div = num1 / num2

	//conficones
	if num1 > num2 {
		fmt.Println("suma es", suma)
		fmt.Println("diferencia es", dif)
	} else {
		fmt.Println("producto es:", prod)
		fmt.Println("la division es:", div)
	}
}
