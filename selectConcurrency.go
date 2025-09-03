package main

import (
	"fmt"
	"time"
)

func selecConcurrency() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {

		time.Sleep(1 * time.Second)
		ch1 <- "Salut c'est fonction 1"
		close(ch1)
	}()

	go func() {

		time.Sleep(4 * time.Second)
		ch2 <- "Moi c'est fonction 2 "
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case m1, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Printf("Message reçu de la part de fonction 1, voici son message : %s", m1)
		case m2, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("Message reçu de la part de fonction 2, voici son message : %s", m2)
		}

	}
	fmt.Println("Select terminé")
}
