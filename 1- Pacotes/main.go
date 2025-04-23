package main

import (
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	println("Print do modulo do go")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	println(erro)
}
