package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1-n2
	return soma, subtracao
}

func main() {
	fmt.Println(somar(1, 2))

	var f = func(texto string){
		fmt.Println(texto)
	}

	f("Enviando um texto")

	resultadoCalculos, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoCalculos)
}