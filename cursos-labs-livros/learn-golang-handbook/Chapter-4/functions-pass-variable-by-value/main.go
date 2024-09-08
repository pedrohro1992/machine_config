package main

import "fmt"

// Variaveis em Go sao passadas por valor,
// Passadas por valor(Pass by Values) significa que
// quando uma variavel e passada em uma funcao, essa funcao
// recebe somente uma copia da variavel
// a funcao nao consegue alterar o valor dessa variavel
func main() {
	x := 5
	increment(x)
	// ainda vai printar 5
	// porque a funcao increment
	// recebeu uma copia de x

	fmt.Println(x)
}

func increment(x int) {
	x++
}
