package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança em GO")
	pessoa1 := pessoa{nome: "Vinicius", sobrenome: "Augusto", idade: 20, altura: 177}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "ADS", "FATEC"}
	fmt.Println(estudante1)
}
