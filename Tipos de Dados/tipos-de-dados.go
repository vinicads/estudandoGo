package main

import (
	"errors"
	"fmt"
)

func main() {
    //numeros
	var numero int = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)



	//numeros reais
	var numeroReal1 float32 = 123.121
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1543543543323.1215435
	fmt.Println(numeroReal2)

	//strings
	var texto string = "Texto"
	fmt.Println(texto)

	char := 'B'
	fmt.Println(char)

	//valor 0
	var texto0 string
	fmt.Println(texto0)

	//booleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	//error
	var erro error = errors.New("Erro de teste")
	fmt.Println(erro)
}