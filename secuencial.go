package main

import (
	"fmt"
	"time"
)

func contador(nombre string, ini int, fin int ) {
	
	for i := ini ; i <= fin; i++ {
		fmt.Printf("%v esta contando en %v\n", nombre, i)
	}
	
}

func main() {
	start := time.Now()
	contador("Lia", 1, 10)
	fmt.Println("Ejecución: ", time.Since(start).Round(time.Nanosecond))
}


