package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}
type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return 3.14 * (c.raio * c.raio)
}

func main() {
	fmt.Println("Interfaces em Go")

	r := retangulo{10, 15}
	escreverArea(r)
}