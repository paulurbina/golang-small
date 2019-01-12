package main

import "fmt"

func contar(canal chan int) {
	for x := 0; x <= 101; x++ {
		canal <- x
		x++
	}
}

func imprimir(canal chan int) {
	var valor int
	for {
		valor = <-canal
		fmt.Println(valor)
	}
}

func main() {
	canal := make(chan int)
	go contar(canal)
	go imprimir(canal)
	var fin string
	fmt.Scanln(&fin)
}
