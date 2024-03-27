package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	base   float64
	altura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.base * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func CalcularAreaForma(f forma) float64 {
	return f.area()
}

func main() {

	forma1 := retangulo{3, 2}
	forma2 := circulo{2.5}

	fmt.Println("Area de um retangulo 3 x 2 = ", CalcularAreaForma(forma1))
	fmt.Println("Area de um circulo com raio de 2.5 = ", CalcularAreaForma(forma2))
}
