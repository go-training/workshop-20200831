package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	outChan := make(chan int)
	errChan := make(chan error)
	finishedChan := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int, outChan chan<- int, errChan chan<- error) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			outChan <- i
		}(i, outChan, errChan)
	}

	go func() {
		wg.Wait()
		close(finishedChan)
	}()

LOOP:

	for {
		select {
		case v := <-outChan:
			log.Println("out:", v)
		case v := <-errChan:
			log.Println("error:", v)
			break LOOP
		case <-time.After(5 * time.Second):
			log.Println("timeout")
			break LOOP
		case <-finishedChan:
			log.Println("All Job Done!")
			break LOOP
		}
	}

	log.Println("main finished")
}
