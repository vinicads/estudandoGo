package formas

type Forma interface {
	area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return 3.14 * (c.raio * c.raio)
}
