package main

import "fmt"

func first() {
	fmt.Println("Primeira")
}

func second() {
	fmt.Println("Segunda")
}

func alunoEstaAprovado(nota1, nota2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (nota1 + nota2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer first()
	second()
	fmt.Println(alunoEstaAprovado(7, 8))
}
