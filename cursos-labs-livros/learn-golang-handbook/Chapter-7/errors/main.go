package main

import (
	"fmt"
	"strconv"
)

// Programas em Go expressa erros com valores error, um erro
// e qualquer tipo que implementa a interface error nativa:
type error interface {
	Error() string
}

// Quando alguma coisa pode dar errado. essa funcao retorna um
// error como o seu ultimo valor de return. Qualquer codigo que
// chama uma funcao que pode retornar um error deve lidar
// com os erros testando se error e nil
func main() {
	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("couldn't convert:", err)
		// Porque "42b" nao e um int valido, printamos:
		// couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
		// Nota:
		// parsing "42b": invalid syntax e retornado pelo metodo .Error()
		return
	}
	// Se chegemos ate aqui, entao a conversao ocorreu com sucesso
	fmt.Println(i)
}
