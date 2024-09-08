package main

import "fmt"

// Uma funcao pode retornar valores que quem esta chamando nao precisa
// da pra ignorar o retorno que a gente nao precisa usando usando o
// underscore(_)
func main() {
	x, _ := getPoint()
	fmt.Println(x)
}

func getPoint() (x int, y int) {
	return 3, 4
}

// No exemplo acima, a funcao getPoint retorna os valores 3 e 4
// na chamada dessa funcao na main, x recebe apenas o 3 e ignora
// o retorn 4
