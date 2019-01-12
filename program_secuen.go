package main

import "fmt"

func main() {
	var num1, num2, suma, producto int

	fmt.Print("ingrese el primer valor")
	fmt.Scan(&num1)

	fmt.Print("ingrese el segundo valor")
	fmt.Scan(&num2)

	//Operacion
	suma = num1 + num2
	producto = num1 * num2

	//devuelve
	fmt.Println("la suma de los valores es:", suma)
	fmt.Println("y su producto es:", producto)
}
