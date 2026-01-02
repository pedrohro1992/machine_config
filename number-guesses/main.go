package main

import (
	"fmt";
	"github.com/number-guesses/internals"
)

const (
	ArrayLength = 9
)

var numberToGuess [ArrayLength]int = [ArrayLength]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	newNumbers := internals.RandomizerNumbers(numberToGuess[:])
	for _, number := range newNumbers {
		sumNumber := internals.SumNumber(number)
		subNumber := internals.SubtracANumber(sumNumber)
		checkResuts := internals.CheckResults(subNumber)
		if checkResuts{
			fmt.Println("Processo executado com sucesso")
		} else {
			fmt.Println("Falha encontrada no processo")
		}
	}
}
