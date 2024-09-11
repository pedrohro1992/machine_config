package main

import "fmt"

// O loop basico em go e escrito como uma syntance padrao tipo C
// for INITIAL; CONDITION; AFTER{
// do something
// }
// INITIAL roda uma vez no comeco do loop e pode
// criar variavies dentro do escopo do loop
//
// CONDITION e checado no inicio de cada interecao. Se a
// CONDITION nao passar, entao o loop quebra
//
// AFTER roda depois de cada interacao
func loopExample() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// vai printar de 1 a 9
}
