package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calcularMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)
	resultadoSoma, resultadoSubtracao := calcularMatematicos(10, 15)
	resultadoSoma2, _ := calcularMatematicos(10, 99)
	_, resultadoSubtracao2 := calcularMatematicos(10, 99)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	fmt.Println(resultadoSoma2)
	fmt.Println(resultadoSubtracao2)
}
