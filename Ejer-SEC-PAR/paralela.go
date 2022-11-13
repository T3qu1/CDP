package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func contador(c chan int, num int){
	wg.Done()
	c <- num
}

func main() {
	start := time.Now()
	chanData := make(chan int, 10)
	
	for i:=1; i<=10; i++{
		wg.Add(1)
		go contador(chanData, i)
	}
	wg.Wait()
	close(chanData)

	for item := range chanData{
		fmt.Println("lia esta contando en ", item)
	}
	fmt.Println("Ejecución: ", time.Since(start).Round(time.Nanosecond))
}

