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

func main() {
	wg := &sync.WaitGroup{}
	for w := 0; w < 5; w++ {
		wg.Add(1)
		go task(w, wg)

	}
	wg.Wait()
	fmt.Println("All tasks completed")
}
