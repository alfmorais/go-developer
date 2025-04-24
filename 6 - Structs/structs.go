package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Hello, world.")

	enderecoUsuarioComum := endereco{logradouro: "Rua dos Bobos, 0", numero: 0}

	var usuario1 usuario
	fmt.Println(usuario1)
	usuario1.nome = "Davi"
	usuario1.idade = 21
	usuario1.endereco = enderecoUsuarioComum
	fmt.Println(usuario1)

	usuario2 := usuario{"Davi", 21, enderecoUsuarioComum}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
