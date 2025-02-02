package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução...")
	if r := recover(); r != nil {
		fmt.Println("Recuperamos a execução com sucesso!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Calculando média do aluno...")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente a 6!")
}

func main() {
	fmt.Println("Panic e Recover em Go")
	fmt.Println(alunoAprovado(6, 6))
}