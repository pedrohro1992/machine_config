package main

import "fmt"

// Arrays
// Arrays saoi grupos de tamanho fixo de variavies
// do mesmo tipo
// o tipo "[n]T" e um array de n valores do tipo T
// para declarar um array de 10 integers:
var myInts [10]int

// ou para declarar um literal inicializado
func arrayInitializedExample() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// SLICES
//
// 99 vezes em 100, usamos um slice ao inves de um array
// quando trabalhamos com lidas ordenadas
// Array sao fixos em tamanho. Uma vez que se cria um
// array com [10]int nao podemos adicionar um novo elemento
//
// Slices possuem tamanhos dinamicos, felixivel, levando
// em conta os elementos de um array
//
// Slices sempre possuem um array por baixo mas nao e sempre
// que o array precisa ser especificado esplicitamente.
func sliceArrayExplicity() {
	// para especificar o array, existe a seguinte maneira
	primes := [6]int{2, 3, 5, 7, 11, 13}
	mySlice := primes[1:4]
	fmt.Println(mySlice)
}

// A syntaxe e a seguinte:
// arrayname[lowIndex:highIndex]
// arrayname[lowIndex:]
// arrayname[:highIndex]
// arrayname[:]
//
// Aonde lowIndex e inclusivo e o highIndex e exclusivo
// Tanto lowIndex ou highIndex ou os dois podem ser omitidos
// para usar o array inteiro naquele lado
