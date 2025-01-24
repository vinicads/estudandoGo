package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops em GO")
	i := 1

	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println(i)

	fmt.Println("----------------------------------------------------")
	for j := 0; j < 10; j++{
		fmt.Println(j)
		time.Sleep(1)
	}

	fmt.Println("----------------------------------------------------")
	nomes := [3]string{"Vinicius", "Augusto", "Santos"}

	for indice, nome := range nomes{
		fmt.Println(indice, " - ", nome)
		for indice2, letra := range nome{
			fmt.Println(indice2, " - ", string(letra))
		}
	}

	fmt.Println("----------------------------------------------------")
	//usando map
	usuario := map[string]string {
		"nome": "Vinicius",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario)
	
	for chave, valor := range usuario{
		fmt.Println(chave, " - ", valor)
	}


}