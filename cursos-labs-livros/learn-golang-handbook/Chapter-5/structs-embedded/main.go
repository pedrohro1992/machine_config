package main

import "fmt"

// Go nao e uma linguagem orientada a objetos, mas
// estruturas embedded(integradas) provem um tipo
// de de heranca somente de dados (data-only inheritance)
//
// Go nao suporta classes ou heranca no senso completo
// do conceito, Embedded Structs(estruturas integradas)
// sao apenas um jeito de elevar e compartilha
// campos entre definicoes de structs

type car struct {
	make  string
	model string
}

type truck struct {
	// car esta integrado, entao, toda a definicao de
	// truck agora tambem adicionalmente contem todos os
	// campos da estrutura car
	car
	bedSize int
}

// Embedded vs Nested
// 1 - um campo de uma estrutura embedded sao acessados no nivel
// mais alto, ao contrario das estruturas nested
// ---------------------**------------------------------------
// campos promovidos (passados de uma estrutura para outra, no embedded) sao acessados
// como campos normais da estrutura, mas nao podem ser usados em composite literals
func main() {
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)

	// campos integrados subiram para o nivel mais alto da estrutura
	// ao inves de acessados via lanesTruck.car.make

	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)
}
