package main

import "fmt"

func main() {
	fmt.Println("Padao de concorrencia Worker Pools em GO")
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}

}

func worker (tarefas <-chan int, resultados chan<- int){
	for trabalho := range tarefas{
		resultados <- fibonacci(trabalho)
	}
}

func fibonacci(posicao int) int{
	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao-2)+fibonacci(posicao-1)
}