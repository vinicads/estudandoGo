package main

import "fmt"

func main() {
	fmt.Println("Ponteiros em GO")
	var var1 int = 10;
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)

	//POnteiro e uma referencia de memoria
	var var3 int
	var ponteiro *int

	fmt.Println(var3, ponteiro)

	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, *ponteiro)

	
	var3++

	fmt.Println(var3, *ponteiro)
}