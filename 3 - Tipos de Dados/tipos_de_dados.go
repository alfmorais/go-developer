package main

import (
	"errors"
	"fmt"
)

func main() {
	// Existe outros tipos de int: int8, int16, int32, int64
	var numero int16 = 100
	fmt.Println(numero)

	// Existe outros uint: uint8, uint16, uint32, uint64
	// uint = unsigned int (sem sinal)
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	var numeroReal2 float64 = 12300000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.45
	fmt.Println(numeroReal3)

	// FIM NUMERO REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FIM STRING
	var texto int
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
