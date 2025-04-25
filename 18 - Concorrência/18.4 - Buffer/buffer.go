package main

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	println(mensagem)
	msg := <-canal
	println(msg)
}
