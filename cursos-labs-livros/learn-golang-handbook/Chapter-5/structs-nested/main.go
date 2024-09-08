package main

import "fmt"

// Structs podem ser aninhadas entidades mais complexas
//

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

func main() {
	// Os campos de uma struct podem ser acessados via o operador ponto(.)
	myCar := car{}
	myCar.FrontWheel.Radius = 5
	fmt.Println(myCar)
}
