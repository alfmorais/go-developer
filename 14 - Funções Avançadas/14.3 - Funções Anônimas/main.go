package main

import "fmt"

// Fun o principal do programa
func main() {
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido: %s", text)
	}("Ola")

	fmt.Println(retorno)
}
