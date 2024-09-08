// Um tipo pode implemenetar qualquer numeros de
// interfaces em Go.
// Por exemplo, a interface vazia interface{}, e
// sempre implementada por qualquer tipo porque
// nao tem requerimentos
package main

import (
	"fmt"
	"log"
)

// Nomeado argumentos de interface
//
// Considerando a seguinte interface:
type Copier interface {
	Copy(string, string) int
}

// Se baseando somente pela declaracao da interface
// nao tem como saber quais valores tem que passar
// para a funcao Copy
//
// Pela signature da funcao, sabemos que a espera
// receber 2 strings, mas o que sao essas strings
// Data Raw? URL? Filename? E como a funcao recebe
// duas strings e retorna um int?
//
// Adicionando argumentos nominados pra deixar mais claro
type copierNamed interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// Agora fica mais claro o que a funcao espera, o primeiro argumento e
// sourceFile o segundo destinationFile e retorna o bytesCopied como
// interger
//
// AFIRMACAO DE TIPOS EM GO
// Quando se trabalha com interfaces em go, de vez em quando
// aparece a necessidade o de acessar o tipo subjacente de um valor
// presente em uma interface
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func defineCircle() {
	// "c" e um novo circle formado a partir de "s"
	// o que e uma nova instancia de shape
	// "ok" e um bool que, retorna true se for um circulo
	// ou false se nao for
	c, ok := s.(circle)
	if !ok {
		// s nao e um circulo
		log.Fatal("s nao e um circulo")
	}

	radius := c.radius
}

// ALTERACAO DE TIPO(SWITH CASE DE TIPO)
// Uma alteracao de tipo, deixa mais facil fazer
// varias afirmacoes de tipo em serie
//
// o switch de tipo e parecido com switch-case
// normal, mas nesse caso especifica tipos
// ao inves de valores

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
