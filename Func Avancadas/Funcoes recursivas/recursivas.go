package main

import "fmt"

func main() {
	fmt.Println("Funcoes recursivas em GO")

	posicao := uint(10)
	fmt.Println(fibonacci(posicao))
}


func fibonacci(posicao uint) uint{
	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao-2)+fibonacci(posicao-1)
}