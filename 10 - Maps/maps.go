package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Pedro",
			"segundo":  "Silva",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")

	usuario2["signo"] = map[string]string{
		"tipo": "Gêmeos",
	}
	fmt.Println(usuario2)
}
