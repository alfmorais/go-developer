package main

import "fmt"

func main() {
	total := soma(1, 2)
	fmt.Println(total)
	escrever("Ola", 1, 2, 3, 4, 5)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
