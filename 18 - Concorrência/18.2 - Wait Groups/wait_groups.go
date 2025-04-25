package main

import (
	"sync"
	"time"
)

func main() {
	// Concorrência != Paralelismo
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func() {
		escrever("Gourotine 1")
		waitGroup.Done()
	}()
	go func() {
		escrever("Gourotine 2")
		waitGroup.Done()
	}()
	go func() {
		escrever("Gourotine 3")
		waitGroup.Done()
	}()
	go func() {
		escrever("Gourotine 4")
		waitGroup.Done()
	}()
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		println(texto)
		time.Sleep(time.Second)
	}
}
