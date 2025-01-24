package main

import "fmt"

func main() {
	fmt.Println("Funcoes avancadas em GO")
	resultado, _ := calculosMatematicos(5, 2)
	fmt.Println(resultado)

}

func calculosMatematicos(n1, n2 int) (soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return
}