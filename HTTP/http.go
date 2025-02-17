package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTP em Go")
	// http eh um protocolo de comunicacao - base da comunicacao de dados na web
	// cliente (faz requisicao) - servidor (processa requisicao e envia resposta)
	// request - response
	// metodos: GET, POST, PUT, DELETE
	// status code: 1xx (informativo), 2xx (sucesso), 3xx (redirecionamento), 4xx (erro no cliente), 5xx (erro no servidor)
	// headers: metadados sobre a requisicao ou resposta
	// body: corpo da mensagem
	// URL: https://www.google.com.br
	// URI: https://www.google.com.br/search
	// Rotas
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	log.Fatal(http.ListenAndServe(":3000", nil))

}