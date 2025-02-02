package main

import "fmt"

func inverterSinal(numero int) int {
	numero = numero * -1;
	return numero
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1;
}

func main() {
	fmt.Println("Ponteiros em Go");
	fmt.Println(inverterSinal(5));

	var numero int = 20;
	inverterSinalComPonteiro(&numero);
	fmt.Print(numero)
}