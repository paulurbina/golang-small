package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cargar(vec *[5000]int, aleatorio *rand.Rand) {
	for f := 0; f < len(vec); f++ {
		vec[f] = aleatorio.Intn(101)
	}
}

func ordenar(vec *[5000]int) {
	for k := 1; k < len(vec); k++ {
		for f := 0; f < len(vec)-k; f++ {
			if vec[f] > vec[f+1] {
				aux := vec[f]
				vec[f] = vec[f+1]
				vec[f+1] = aux
			}
		}
	}
	wg.Done()
}

func diferenciaTiempo(hora1, hora2 time.Time) time.Duration {
	diferencia := hora2.Sub(hora1)
	return diferencia
}

func main() {
	aleatorio := rand.New(rand.NewSource(time.Now().UnixNano()))
	var vec1 [5000]int
	var vec2 [5000]int
	cargar(&vec1, aleatorio)
	cargar(&vec2, aleatorio)

	var hora1, hora2 time.Time
	hora1 = time.Now()
	wg.Add(2)
	go ordenar(&vec1)
	go ordenar(&vec2)
	wg.Wait()

	hora2 = time.Now()
	di := diferenciaTiempo(hora1, hora2)
	fmt.Println("la cantidad de segundo de diferencia: ", di.Seconds())
	fmt.Println("N# procesadores de maquina: ", runtime.NumCPU())

}
