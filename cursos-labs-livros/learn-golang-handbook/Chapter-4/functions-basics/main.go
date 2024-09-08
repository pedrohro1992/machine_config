package main

// Funcoes em go recebem zero ou mais argumentos
// Pra deixar mais facil de ler, o tipo da variavel vem depois da declaracao da variavel
func sub(x int, y int) int { // Essa linha representa a "Function signature"
	return x - y
	// Funcao aceita dois parametros do tipo int e retorna outro int
}

// -- MULTIPLOS PARAMETROS --
// Quando multilos argumentos sao do mesmo tipo, o tipo so precisa
// ser declarado depois do ultimo parametro, assumindo que eles estao em ordem

func add(x, y int) int {
	return x + y
}

// Se eles nao estao em ordem, os tipos precisam ser definidos separadamente
