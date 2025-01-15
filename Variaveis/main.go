package main

import "fmt"

func main() {
	//varuaveis explicitas
	var texto string = "Variavel String Explicita"
	fmt.Println(texto)

	//varuaveis implicitas
	texto2 := "Variavel String Implicita"
	fmt.Println(texto2)

	var (
		texto3 string = "Texto 3"
		texto4 string = "Texto 4"
	)

	fmt.Println(texto3, texto4)

	variavel5, variavel6 := "Texto 5", "Texto 6"
	fmt.Println(variavel5, variavel6)

	const contante1 string = "Constante 1"

	fmt.Println(contante1)

	//invertendo variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}