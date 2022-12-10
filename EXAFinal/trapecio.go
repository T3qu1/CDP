package main

import (
	"fmt"
	"math"
	"time"
)

func particion(parts chan int, areas chan float64) {
	f := func(x float64) float64 {
		return (math.Pow(x, 2) + 1)
	}
	for a := range parts {
		areas <- reglatrapecio(f, 2, 17, a)
	}
}

func reglatrapecio(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

func main() {
	fmt.Println("--- Inicio programa ---")
	n := 100

	parts := make(chan int, n)
	areas := make(chan float64, n)

	for i := 1; i<13; i++{
		go particion(parts, areas)
	}
	
	ini := time.Now()
	for i := 0; i < n; i++ {
		parts <- i
	}
	fin := time.Since(ini).Nanoseconds()
	fmt.Println(fin)
	fmt.Println("--- Programa terminado ---")

	close(parts)

	for i:=1; i<n; i++{
		fmt.Println(<-areas)
	}

}

