package main

import "fmt"

func main() {
	// Declaracao de if em go nao precisa de parenteses em volta da condicao
	height := 4
	height2 := 4
	email := "pedro@cacete.com"

	if height > 4 { // nao vai printar
		fmt.Println("You are tall enough!")
	}

	// Declarando com if else
	if height2 > 6 {
		fmt.Println("You are super tall!")
	} else if height2 > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough")
	}

	// Uma declaracao de if pode ter um statment inicial, a variavel criada
	// no statment inicial e definida somente no escopo do corpo do if
	// if INITIAL_STATEMENT; CONDITION {
	// }
	// Com isso a gente pode declarar uma variavel que vai ser usada somente
	// no contexto do if
	if length := len(email); length > 1 {
		fmt.Printf("Email is invalid, email: %v", length)
	}
}
