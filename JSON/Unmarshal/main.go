package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Breed string `json:"breed"`
}

func main() {
	fmt.Println("JSON para STRUCT em GO")

	dogInJSON := `{"name":"Rex","age":3,"breed":"Poodle"}`
	fmt.Println(dogInJSON)

	var c dog

	if erro := json.Unmarshal([]byte(dogInJSON), &c) ; erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(c)
}
