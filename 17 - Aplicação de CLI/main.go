package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	aplication := app.Gerar()
	if appError := aplication.Run(os.Args); appError != nil {
		log.Fatal(appError)
	}
}
