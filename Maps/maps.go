package main

import "fmt"

func main() {
	fmt.Println("Maps em GO")

	usuario := map[string]string {
		"nome": "Vinicius",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Vinicius",
			"ultimo": "Santos",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro"])
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string {
		"nome": "Virgem",
	}
	fmt.Println(usuario2)
}