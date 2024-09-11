package main

import "fmt"

// Loops in Go podem omitir algumas partes das
// sessoes que compoem um loop, por exemplo
// a parte onde e declarada CONDITION (parte do meio)
// pode ser omitida, o que faz o loop rodar para sempre
//
// for INITIAL; ; AFTER {
// // faz alguma coisa
// }
// **
// EM GO NAO EXISTE LOOP WHILE
//
// A maioria das linguagens de programacao possuem
// um conceito conhecido como WHILE(enquanto)
// Como o Go permite omissoes nas sessoes do loop for
// um loop while e somente um loop for onde
// existe apenas a CONDITION
// for CONDITION {
// // faz alguma coisa
// }
// EXEMPLO:
func whileForLoopExample() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grow to ", plantHeight, "inches")
}

// O que vai printar:
// still growing! current height: 1
// still growing! current height: 2
// still growing! current height: 3
// still growing! current height: 4
// plant has grown to 5 inches
