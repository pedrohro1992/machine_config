package main

import "fmt"

// Mesmo Go NAO e uma linguagem orientada a objetos, ainda sim suporta
// que metodos sejam definidos dentro de estruturas
//
// Metodos sao apenas funcoes que possuem um receptor
// Um receptor e um parametro especial que sintaticamente
// vai ANTES do nome da funcao

type rect struct {
	width  int
	height int
}

// area tem um reciver de (r rect)
func (r rect) area() int {
	return r.width * r.height
}

func main() {
	r := rect{
		width:  5,
		height: 10,
	}
	fmt.Println(r.area())
}

// Os recivers sao apenas um tipo especial de parametro de funcao.
// Recivers sao importantes porque eles permitem a definicao de
// interfaces que as structs(e outros tipos) podem implementar
