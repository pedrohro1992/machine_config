package main

import "math"

// Interfaces sao uma collecao de signatures de metodo
// um tipo "implementa" uma interface se tiver todos os metodos
// da interface fornecedida definidos nele
//
// No exemplo abaixo, o "shape(forma)" tem que ser capaz
// de retornar sua area e perimitro
// tanto "rect(retangulo)" quanto "circle" preenchem a interface

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Quando um tipo implementa uma interface, ele pode
// ser usado como um tipo de interface (interface type)
//
// Interfaces sao implementadas IMPLICITAMENTE
//
// Um tipo nunca declara que implementa uma interface,
// se uma interface existe e o tupo tem os metodos
// apropiados definidos, entao o tipo preenche
// automaticamente essa interface
