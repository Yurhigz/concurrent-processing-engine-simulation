package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	randomTime := time.Duration(rand.IntN(5)+1) * time.Second
	fmt.Printf("Task %v started... It will take %v \n", id, randomTime)
	time.Sleep(randomTime)
	fmt.Printf("Task %v done\n", id)
}

func producer(value int, values chan int) {
	for i := range value {
		values <- i
	}
	close(values)
}

func consumer(worker int, values chan int) {
	for value := range values {
		fmt.Printf("worker %v processed %v \n", worker, value)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	values := make(chan int)
	workers := 3

	go producer(10, values)

	for w := range workers {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			consumer(w, values)
		}(w)
	}
	wg.Wait()

}
