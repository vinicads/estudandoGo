package main

import "fmt"

func main() {
	//aritmeticos
	soma := 10 + 5
	subtacao := 10 - 5
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDivisao := 10 % 4
	fmt.Println(soma, subtacao, divisao, multiplicacao, restoDivisao)
	fmt.Println(parOuImpar(1))

	//relacionais
	fmt.Println(1 > 3)
	fmt.Println(5 >= 3)
	fmt.Println(1 == 3)
	fmt.Println(1 != 3)


	//logicos
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(false || true)
	fmt.Println(false || false)

	//unarios
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero +=10
	fmt.Println(numero)
	numero *=10
	fmt.Println(numero)
}

func parOuImpar(numero int) string{
	if (numero % 2 != 0){
		return  "IMPAR"
	}
	return "PAR"
}