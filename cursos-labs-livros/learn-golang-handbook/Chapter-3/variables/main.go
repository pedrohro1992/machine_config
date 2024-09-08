package main

import "fmt"

func main() {
	// Formatando string com golang
	// Desse jeito da pra fazer porque o golang vai assumir o tipo, mas e melhor usar os tipos especificos quando possivel
	s := fmt.Sprintf("I am %v years old\n", 10)
	fmt.Printf(s)

	//%s interpola uma string
	st := fmt.Sprintf("I am %s years old\n", "way too many")
	fmt.Printf(st)

	//%d interpola um integer em forma decimal
	d := fmt.Sprintf("I am %d years old\n", 10)
	fmt.Printf(d)

	//%f interpola um decimal
	f := fmt.Sprintf("I am %f years old\n", 10.523)
	fmt.Printf(f)
	// Tambem pode adicionar numeros depois do ponto, apontando eles antes do f
	f2 := fmt.Sprintf("I am %.2f years old\n", 10.523)
	fmt.Printf(f2)
}
