package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
    Name  string `json:"name"`
    Age   uint   `json:"age"`
    Breed string `json:"breed"`
}


func main() {
	fmt.Println("Struct para JSON em GO")

	c := dog{"Rex", 3, "Poodle"}
	fmt.Println(c)

	dogJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(dogJson))
}