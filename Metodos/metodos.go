package main

import "fmt"

type Usuario struct {
	nome string
	idade int
}

func (u Usuario) salvar(){
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u *Usuario) fazerAniversario(){
	u.idade++
}

func main() {
	fmt.Println("Métodos em Go")
	usuario1 := Usuario{"Fulano", 20}
	fmt.Println(usuario1)
	usuario1.fazerAniversario()
	fmt.Println(usuario1)
	usuario1.salvar()
}