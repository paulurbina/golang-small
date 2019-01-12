package main

import (
	"fmt"
	"math/rand"
	"time"
)

func aleatorio1() {
	var ale int
	for f := 0; f < 200; f++ {
		ale = rand.Intn(101)
		fmt.Print(ale, "-")
	}
}

func separador() {
	for f := 0; f < 100; f++ {
		fmt.Print("*")
	}
}

func aleatorio2() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ale int
	for f := 0; f < 200; f++ {
		ale = r.Intn(101)
		fmt.Print(ale, "-")
	}
}

func main() {
	aleatorio1()
	separador()
	aleatorio2()
}
