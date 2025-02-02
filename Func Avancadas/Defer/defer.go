package main

import "fmt"

func main() {
	fmt.Println("Defer em Go")
	defer func1()
	func2()
}

func func1() {
	fmt.Println("Executando func1")
}

func func2() {
	fmt.Println("Executando func2")
}