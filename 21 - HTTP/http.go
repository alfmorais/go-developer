package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello World")
	w.Write(message)
}

func main() {
	/*
		HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
		CLIENTE - SERVIDOR
		REQUEST - RESPONSE
		ROTAS -> GET, POST, PUT, DELETE, PATCH
		URI - IDENTIFICADOR DO RECURSO
		MÉTODO HTTP
	*/
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
