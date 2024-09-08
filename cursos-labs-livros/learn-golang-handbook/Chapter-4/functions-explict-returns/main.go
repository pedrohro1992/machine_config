package main

import "errors"

// Mesmo a funcao tendo valores de retorno nominados, ainda
// podemos realizar retornos explicitos

func getCoords() (x, y int) {
	return x, y // isso e explicito
}

// Usando o padrao explicito, podemos sobreescrever os valores de return
func getCoordsOverwrite() (x, y int) {
	return 5, 6 // isso e explicito, x e y nao serao reornados
}

// De outra forma, se desejamos retornar os valores definidos na signature da funcao
// podemos usar o retorno naked
func getCoordsNaked() (x, y int) {
	return // retorno implicito de x e y
}

// -- Benefecios dos retornos nominados --
//
// 1 - Ajuda na documentacao de uma funcao, sabemos o que uma funcao vai retornar
// so olhando para sua signature, sem necessitar de comentarios
// Valores de retorno nominados sao importantes em funcoes mais longas
func calculatorNamed(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

// O que e mais facil de entender que:
func calculatorLong(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}

	mul := a * b
	div := a / b
	return mul, div, nil
}

// No primeiro exemplo, com os valores nominados a gente consegue
// deduzir cada valor de retorno olhando somente para a signature
// da funcao
//
// -- ** ---
// Em resumo, a escolha entre retornos naked e nominados
// vai depender da complexidade da funcao
