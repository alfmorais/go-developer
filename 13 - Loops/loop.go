package main

import "time"

func main() {
	// for i := 0; i < 10; i++ {
	// 	println("Incrementando i: ", i)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joaquim", "Maria", "Joana"}
	for index, nome := range nomes {
		println("Posição: ", index, " - Nome: ", nome)
		time.Sleep(time.Second)
	}

	for _, letra := range "PALAVRA" {
		println("Letra: ", string(letra), " Tabela ASCII: ", letra)
		time.Sleep(time.Second)
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		println("Chave: ", chave, " - Valor: ", valor)
		time.Sleep(time.Second)
	}
}
