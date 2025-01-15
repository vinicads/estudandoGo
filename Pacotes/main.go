package main

import (
	"fmt"
	"modulo/auxiliar"
	 "github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Write()
	erro := checkmail.ValidateFormat("viniciusadstrab@outlook.com")
	fmt.Println(erro)
}