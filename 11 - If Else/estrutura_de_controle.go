package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	var numero int = 0

	if numero > 12 {
		fmt.Println("Maior que 12 anos")
	} else {
		fmt.Println("Menor que 12 anos")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que zero")
	}
}
