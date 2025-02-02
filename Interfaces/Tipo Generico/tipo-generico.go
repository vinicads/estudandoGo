package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipo Genérico em Go")
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1: "string",
		"string": 1,
		true: false,
	}

	fmt.Println(mapa)
}