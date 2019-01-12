package main

import (
	"fmt"
	"time"
)

func fechaPc() {
	var ahora time.Time
	ahora = time.Now()
	fmt.Println("anio:", ahora.Year())
	fmt.Println("mes:", ahora.Month())
	fmt.Println("Dia:", ahora.Day())
	fmt.Println("Hora:", ahora.Hour())
	fmt.Println("Minuto:", ahora.Minute())
	fmt.Println("Segundo:", ahora.Second())
}

func pasadoDosFechas() {
	navidad2019 := time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC)
	var ahora time.Time
	ahora = time.Now()
	fmt.Println(ahora)
	diferencia := ahora.Sub(navidad2019)
	fmt.Println("cantidad de horas de diferencia:", diferencia.Hours())
}

func main() {
	// fechaPc()
	pasadoDosFechas()
}
