package main

import "errors"

// Go suporta a abilidade de retorna cedo de uma funcao
// o que diminui codigo e pode ser usado como guard clauses
// Guard Clauses aproveita a abilidade de retornar mais cedo
// ou continuar por um loop fazendo as conditionais
// aninhadas com apenas uma dimensao, ao inves de um monte
// de if e else encadinhado, a gente so retorna antes no final de
// bloco de condicional

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

// A forma de lidar com erros em Go encouraja os desenvolvedores
// a fazer o uso das guard clauses, o que evita condicionais aninhadas
// o Exemplo abaixo demonstra uma funcao com conditionais aninhadas
func getInsuranceAmountNested(status insuranceStatus) int { // Voltar depois pra implementar o tipo
	amount := 0
	if !status.hasInsurance() {
		amount = 1
	} else {
		if status.isTotaled() {
			amount = 10000
		} else {
			if status.isDented() {
				amount = 160
				if status.isBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}

// A mesma logica acima, mas implementada usando guard clauses
func getInsuranceAmountEarlyReturn(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}

func hasInsurance() bool {
	return true
}

func isTotaled() bool {
	return false
}

func isDented() bool {
	return false
}

func isBigDent() bool {
	return false
}

// Com as guard clauses fica mais simples de implementar e de ler
// levando em consideracao que, quando escrevemos codigo, e importante
// tentar reduzir a carga cognitiva de quem esta lendo, reducindo o numero
// de entidades que o leitor precisa pensar sobre a qualquer momento
