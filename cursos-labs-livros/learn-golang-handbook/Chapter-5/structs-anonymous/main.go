package main

// Uma estrutura anonima(anonymous struct) e igual uma estrutura
// regular, mas e definida sem um nome, com isso
// nao pode mais ser usada em nenhum outro lugar no codigo
func main() {
	//
	// usando um segundo par de chaves logo apos declarar o tipo
	// Para criar uma estrutura anonima, apenas inicianiza a instancia

	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}
	println(myCar)
}

// Podemos tambem aninhar estruturas anonimas como campos
// dentro de outras estruturas
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Wheel e um campo contendo uma estrutura anonima
	Wheel struct {
		Radius   int
		Material string
	}
}

// ** QUANDO USAR ESTRUTURAS ANONIMAS
// Por via de regra, e sempre melhor usar as estruturas nominadas
// fica mais facil de ler o seu codigo depois, e, uma outra vantagem
// e que elas ficam reutilizaveis, estruturas anonimas se encaixam melhor
// em contextos que a estruturas nao vai ser referenciada ou alterada com frequencia
// se for pra uso exclusico de uma funcao, por exemplo,
