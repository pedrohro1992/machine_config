package main

import "fmt"

// Funcoes variadicas
//
// Muitas funcoes, especialmente aquelas que estao
// presentes na biblioteca padrao, podem receber
// um numero arbritario de argumentos finais.
// Isso pode ser feito atraves da syntas "..."
// na assinatura da funcao
//
// Uma funcao variadica recebe os argumentois variadicos
// de um slice
func concat(strs ...string) string {
	final := ""
	// strs is just a slice of strings
	for str := range strs {
		final = strs
	}
	return final
}

func main() {
	final := concat("Hello ", "there ", "friend!")
	fmt.Println(final)
	// Output: Hello there friend!
}
