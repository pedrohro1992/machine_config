// INTERFACES LIMPAS
// Escrever interfaces e dificil, sempre que se precisa lidar com
// abstracao quando se escreve codigo, o simples pode se tornar
// complexo rapido se o devido cuidado nao forem tomados
//
// # Melhores praticas para escrever interfaces
//
// 1 - Deixar interfaces pequenas:
//
//	 Interfaces sao feitas para definir
//	 o comportamento minimo necessario
//	 para representar com precisao
//	 uma ideia ou um conceito
//	Aqui está um exemplo do pacote HTTP padrão
//	de uma interface maior que é um bom exemplo de definição de comportamento mínimo:
package main

import (
	"io"
	"os"
)

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	stat() (os.FileInfo, error)
}

// Na interface acima, qualquer Type que satisfaca
// o comportamento da intrarface pode ser considerado
// pode ser considerado pelo package HTTP como File,
// isso e util porque o package HTTP nao precisa saber
// se esta lidando com um arquivo(file) no disco
// um buffer de network ou um simples "byte[]"
//
// ***************************************************
//
// 2 - Interfaces nao devem ter conhecimento dos types que a satisfacem
//
// Uma interface deve definer o que e necessario para outros
// types se classificarem como membro daquela interface
// Uma interface nao deve estar ciente de nenhum tipo que
// a satisfaca no momento de design
// Por exemplo, vamos assumir que estamos construindo uma
// interface que definem os componentes necesarios para
// descrever um carro
type car interface {
	Color() string
	Speed() int
	IsFiretruck() bool
}

// Colo() e Speed() fazem sentindos, sao metodos que
// estao confinados no escopo de um carro
// IsFiretruck e um anti-pattern. Estamos
// forcando todos os carros a a declarem, um
// IsFiretruck := false ou true, o que leva
// a uma infinidade de possibilidades  como
// IsPickup() IsSedan() IsTank()
// e assim por diante
//
// Ao inves disso, devemos contar com a funcionalidade
// nativa de afirmacao de tipo( type assertion) para
// entregar o type subjacente quando recebido uma
// instancia da interfcace "car". Ou se uma sub-interace
// e necessaria, pode ser definida como:
type firetruck interface {
	car
	HoseLength() int
}

// Essa interface vai herdar os metodos e adiciona
// mais um metodo requirido, para fazer de um "car" um "firetruck"
//
// ***************************************************
//
// 3 - Interfaces nao sao clasess
//
// - Interfaces nao sao classes, elas sao similares
// - Interfaces nao tem construdores ou desconstrudores que necessitam que os dados sejam criados ou destruidos
// - interfaces nao sao herarquicas por natureza, embora haja um recurso sintatico que cria
// interfaces que sao conjuntos de outras interfaces
// - Interfaces definem signatures de funcao mas nao o seu comportamento interno, criar uma interface nao vai
// deixar o seu codigo DRY(Don't Repeat Yourself) em relacao a estruturar os metodos, por exemplo, se um tipo satisfas
// a interface fmt.String, todos eles vao precisar da sua propria versao do da funcao String()
