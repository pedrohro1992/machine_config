package main

import "fmt"

// Na maioria das vezes, nao precisamos pensar no array
// em baixo de um slice. Pode se criar um slice com
// a funcao 'make'
func sliceCreateWithMake() {
	// func make([]T, len, cap) []T
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)

	// o argumento "capacidade"(o 10 no exemplo acima)
	// geralmente e omitido e o defaulta para o
	// length
	mySlice2 := make([]int, 5)
	fmt.Println(mySlice2)
}

// Slices criadas com 'make' sao preenchidas
// com valores zerados referente ao tipo
//
// Se ha necissidade de criar um slice com
// valores especificos pode se usar uma
// slice literal
func sliceLiteral() {
	mySlice := []string{"I", "love", "go"}
	fmt.Println(mySlice)
}
