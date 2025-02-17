package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	fmt.Println("Introducao a testes automatizados em Go")
	tipoEndereco := enderecos.TipoDeEndereco("Rua ABC");
	fmt.Println(tipoEndereco)
}