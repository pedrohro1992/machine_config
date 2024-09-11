package main

import "fmt"

// Continuar atraves de um loop
// A keyword "continue" para a interecao corrente
// do loop e continua para a proxima interacao
// "continue" e um metodo poderoso para a implementacao
// do pattern "guard clause" dentro dos loops
func loopContinueExample() {
	for i := 0; 1 < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Resultado da funcao:
// 1
// 3
// 5
// 7
// 9

// Sair de um loop
// a keyword "break" para a interecao corrente do loop e
// sai de dentro dele
func loopBreakExample() {
	for i := 0; 1 < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

// Resultado da funcao:
// 0
// 1
// 2
// 3
// 4
