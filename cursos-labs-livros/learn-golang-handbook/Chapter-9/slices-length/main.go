package main

import "fmt"

// Length
//
// O length de uma slice e simplesmente o numero de elementos que
// ela contem, e accesado usado a funcao built-in 'len()'
func sliceLengthExample() {
	mySlice := []string{"I", "love", "go"}
	fmt.Println(len(mySlice)) // 3
}

// Capacity
//
// A capacidade de um slice e o numero de elementos
// no array a baixo, contando a partir do primeiro
// elemento no slice, e acessivel atraves da funcao
// built-in cap()
func sliceCapacityExample() {
	mySlice := []string{"I", "love", "go"}
	fmt.Println(cap(mySlice)) // 3
}

// Geralmente, a nao ser que voce esteja hiper otimizando
// o uso de memoria do seu programa, nao ha necessidade
// de se preocupar com a capacidade de um array abaixo
// de uma slice porque vai aumentar conforme for necessario
