/*package main

import (
	"fmt"
	"sync"
)


func contador() int {
	return 1

}

func main() {
	dataChan := make(chan int)

	go func(){
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(){
				defer wg.Done()
				result := contador()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan{
		fmt.Println("lia", n)
	}
}
*/

/*package main

import(
	"fmt"
	//"time"
)

func contador(nombre string, ini int, fin int, chandata chan int ){
	var aux int
	for i:=ini, i<=fin, i++{
		aux = <- chandata
		fmt.Printf("%v esta contando en %v\n", nombre, aux)
	}
}

func main(){
	chanData := make(chan int)

	go contador("Lia", 1, 10, chanData)
	go contador("Mateo", 1, 10, chanData)

	i := 0

	for {
		i++
		chanData <- i
	}

}*/

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

