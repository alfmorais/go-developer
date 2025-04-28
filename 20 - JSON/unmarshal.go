package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	kyraEmJSON := `{"nome":"Kyra","raca":"Red Heller","idade":5}`
	estrelaEmJSON := `{"nome":"Estrela","raca":"Blue Heller"}`

	var kyra dog
	if erro := json.Unmarshal([]byte(kyraEmJSON), &kyra); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(kyra)

	estrela := make(map[string]string)
	if erro := json.Unmarshal([]byte(estrelaEmJSON), &estrela); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(estrela)
}
