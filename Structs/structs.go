package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	maiorDeIdade bool
	redes redesSociais
}

type redesSociais struct {
	linkedin string
	github string
	instagram string
}

func main() {
	fmt.Println("Arquivo de Structs")
	
	var user1 usuario
	user1.nome = "Vinicius"
	user1.idade = 20
	user1.redes = redesSociais{linkedin: "urlLinkedin", github: "urrlGithub"}
	verificaMaioridade(user1.idade)

	fmt.Println(user1)

	var redesExemplo = redesSociais{linkedin: "exemploLinkedin", github: "exemploGithub", instagram: "exemploInstagram"}
	user2 := usuario{nome: "Augusto", idade: 17, redes: redesExemplo}
	user2.maiorDeIdade = verificaMaioridade(user2.idade)
	fmt.Println(user2)
}

func verificaMaioridade(idade uint8) bool {
	if (idade >= 18){
		return true
	}else{
		return false
	}
}
