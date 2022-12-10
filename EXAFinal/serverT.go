package main

import (
	"encoding/gob"
	"fmt"
	"math"
	"net"
)

func ReglaTrapecio(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}	
	return sum * h
}

func Particion(Parts chan int, areas chan float64) {
	f := func(x float64) float64 {
		return (math.Pow(x, 2) + 1)
	}
	for a := range Parts {
		areas <- ReglaTrapecio(f, 2, 17, a)
	}
}

func servidor() {
	servidor, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := servidor.Accept()
		c2, err2 := servidor.Accept()
		c3, err3 := servidor.Accept()
		if err != nil && (err2 != nil) && (err3 != nil) {
			fmt.Println(err, err2, err3)
			continue
		}

		go handleClientA(c)
		go handleClientB(c2)
		go handleClientT(c3)
	}
}

func handleClientA(c net.Conn) {
	var a2 int

	err := gob.NewDecoder(c).Decode(&a2)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", a2)
	}
}

func handleClientB(c2 net.Conn) {
	var b2 int

	err2 := gob.NewDecoder(c2).Decode(&b2)
	if err2 != nil {
		fmt.Println(err2)
		return
	} else {
		fmt.Println("Mensaje: ", b2)
	}
}

func handleClientT(c3 net.Conn) {
	var n int
	Parts := make(chan int, n)
	areas := make(chan float64, n)
	err3 := gob.NewDecoder(c3).Decode(&n)
	if err3 != nil {
		fmt.Println(err3)
		return
	} else {
		for i := 0; i < n; i++ {
			go Particion(Parts, areas)
		}
		for i := 0; i < n; i++ {
			Parts <- i
		}
		close(Parts)
		for i := 0; i < n; i++ {
			fmt.Println(<-areas)
		}
	}
}

func main() {

	go servidor()

	var input string
	fmt.Scanln(&input)
}