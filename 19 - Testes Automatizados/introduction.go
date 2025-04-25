package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoDeEndereco)

	tipoDeEndereco = enderecos.TipoDeEndereco("Rua das Rosas")
	fmt.Println(tipoDeEndereco)

	tipoDeEndereco = enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoDeEndereco)

	tipoDeEndereco = enderecos.TipoDeEndereco("Praça das Rosas")
	fmt.Println(tipoDeEndereco)
}
