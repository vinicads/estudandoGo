package main

import "fmt"

func main() {
	fmt.Println("Funcoes variaticas em O")
	fmt.Println(soma(1,2,3,4,5,6,7,8,9,10,11,100))
}

func soma(numeros ...int)(resultado int){
	for _, numero := range numeros{
		resultado += numero
	}
	return
}