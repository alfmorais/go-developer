package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", user.nome)
}

func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() {
	user.idade++
}

func main() {
	fmt.Println("Métodos")
	user := usuario{"Joaquim", 3}
	fmt.Println(user)
	user.salvar()
	user.maiorDeIdade()
	fmt.Println(user.idade)
	user.fazerAniversario()
	fmt.Println(user.idade)
}
