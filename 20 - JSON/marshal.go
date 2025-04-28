// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type cachorro struct {
// 	Nome  string `json:"nome"`
// 	Raca  string `json:"raca"`
// 	Idade uint   `json:"idade"`
// }

// func main() {
// 	kyra := cachorro{"Kyra", "Red Heller", 5}
// 	kyraEmJSON, erro := json.Marshal(kyra)

// 	if erro != nil {
// 		log.Fatal(erro)
// 	}

// 	fmt.Println(kyraEmJSON)
// 	fmt.Println(bytes.NewBuffer(kyraEmJSON))

// 	estrela := map[string]string{
// 		"nome": "Estrela",
// 		"raca": "Blue Heller",
// 	}
// 	estrelaEmJSON, erro := json.Marshal(estrela)

// 	if erro != nil {
// 		log.Fatal(erro)
// 	}

// 	fmt.Println(estrelaEmJSON)
// 	fmt.Println(bytes.NewBuffer(estrelaEmJSON))
// }