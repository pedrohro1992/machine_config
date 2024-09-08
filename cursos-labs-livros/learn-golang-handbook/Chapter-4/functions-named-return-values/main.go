package main

func main() {
}

// valores de retorno podem receber nomes, e se eles receberem
// serao tradados como se fossem novas variaveis
// declaradas no topo da funcao
//
// valores de retorno nomimados sao a melhor maneira
// de documentar o proposito do valor retornado
//
// uma declaracao de retrorno sem argumentos retorna
// os valores de retorno nominal
// o que e conhecido como "Naked Return"
// naked returns devem ser usados apenas em funcoes curtas
// uma vez que eles podem prejudicar a reabilidade em funcoes
// mais longas
func getCoordsNaked() (x, y int) {
	//x e y vai ser retornados por default
	//
	return // retorna de forma automatica x e y
}

// E a mesma coisa que:
func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}

// No primeiro exemplo x e y sao valores de retorno, no final da funcao
// podemos escrever return para retornar os valores dessas duas variaveis
// ao inves de escrever "return x, y"
