package main

import (
	"encoding/gob"
	"fmt"
	"net" 
)

func cliente() {
	c, err := net.Dial("tcp", ":9999")
	c2, err2 := net.Dial("tcp", ":9999")
	c3, err3 := net.Dial("tcp", ":9999")
	if err != nil && (err2 != nil) && (err3 != nil) {
		fmt.Println("error1: %v, error2: %v, error3: %v", err, err2, err3)
		return
	}
	a2 := 2
	b2 := 17
	n := 100
	fmt.Println("Valor de a2: ", a2)
	fmt.Println("Valor de b2: ", b2)
	fmt.Println("Numero de trapecios: ", n)
	err = gob.NewEncoder(c).Encode(a2)
	err = gob.NewEncoder(c2).Encode(b2)
	err = gob.NewEncoder(c3).Encode(n)

	if err != nil && (err2 != nil) && (err3 != nil) {
		fmt.Println("error1: %v, error2: %v, error3: %v", err, err2, err3)
	}
	c.Close()
	c2.Close()
	c3.Close()
}

func main() {
	go cliente()
	
	var input string
	fmt.Scanln(&input)
}