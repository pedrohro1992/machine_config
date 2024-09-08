package main

// Usamos estruturas em Go para representar dados estruturados
// geralmente e conveniente agrupar variaveis de diferentes tipos
// juntas,
// por exemplo, a representacao de um carro ficaria
// da seguinte forma
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

// Isso cria uma nova estrutura chamnada car
// todos os carros possuem:
// Make
// Model
// Height
// Width
// ** ----------- ***
// Em go, o que geralmente e representado por uma struct
// voce representaria com um dicionary no python ou
// um object literal no JavaScript(Nao sei o que isso quer dizer muito bem)
