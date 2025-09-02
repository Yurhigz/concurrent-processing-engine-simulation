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

func fanin(ch1 chan int, ch2 chan int) chan int {
	var wg sync.WaitGroup
	wg.Add(2)
	aggregatedChannel := make(chan int, cap(ch1)+cap(ch2))
	go func() {
		defer wg.Done()
		for val := range ch1 {
			aggregatedChannel <- val
		}
	}()
	go func() {
		defer wg.Done()
		for val := range ch2 {
			aggregatedChannel <- val
		}

	}()
	go func() {
		wg.Wait()
		close(aggregatedChannel)

	}()

	return aggregatedChannel
}

func consumer(worker int, values chan int) {
	for value := range values {
		fmt.Printf("worker %v processed %v \n", worker, value)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	nbProcess := 100
	ch1 := make(chan int, nbProcess)
	ch2 := make(chan int, nbProcess)
	workers := 3

	producer(nbProcess, ch1)
	producer(nbProcess, ch2)

	channel := fanin(ch1, ch2)

	for w := range workers {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			consumer(w, channel)
		}(w)
	}
	wg.Wait()

}
