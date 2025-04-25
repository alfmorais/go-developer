package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	var posicao uint = 16

	for index := uint(1); index < posicao; index++ {
		fmt.Println(fibonacci(index))
	}
}
