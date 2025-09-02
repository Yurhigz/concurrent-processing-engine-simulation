package main

import (
	"fmt"
	"sync"
)

// Producer envoie des valeurs dans un channel et le ferme ensuite
func Producer(value int, values chan int) {
	for i := 0; i < value; i++ {
		values <- i
	}
	close(values)
}

// FanIn fusionne deux channels d'entiers en un seul
func FanIn(ch1, ch2 chan int) chan int {
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

// Consumer lit depuis un channel et affiche quel worker traite quelle valeur
func Consumer(worker int, values chan int) {
	for value := range values {
		fmt.Printf("worker %v processed %v\n", worker, value)
	}
}

// Main

// wg := &sync.WaitGroup{}

// nbProcess := 100
// ch1 := make(chan int, nbProcess)
// ch2 := make(chan int, nbProcess)
// workers := 3

// producer(nbProcess, ch1)
// producer(nbProcess, ch2)

// channel := fanin(ch1, ch2)

// for w := range workers {
// 	wg.Add(1)
// 	go func(worker int) {
// 		defer wg.Done()
// 		consumer(w, channel)
// 	}(w)
// }
// wg.Wait()
