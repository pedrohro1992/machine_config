package main

import "fmt"

// Como erros sao apenas interfaces, podemos construir
// tipos customizados que implementar a interface "error"
// Segue um exemplo de uma struct chamada "userError" que
// implementa a interface "error"
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with ther account", e.name)
}

// Agora pode ser usado como erro:
func sendSMS(msg, userName string) error {
	if !canSendToUser(userName) {
		return userError{name: userName}
	}
	// Implementa logica pra envio de mensagem
	fmt.Println(msg)
	return nil
}

func canSendToUser(userName string) bool {
	// faz alguma logica com o usuario
	fmt.Println(userName)
	// retorna um bool
	return true
}

// Programs em go expressao erros com valores "error". "Error-values"
// sao qualquer tipo que implementa a interface build-in "error interface"
//
// E preciso ter em mente que o jeito que Go lida com erro e bastante unico
// a maioria das linguagens trata erro como algo especial e diferente
// no Python, por exemplo, e leventada uma exception e em JavaScript
// joga e recebe erros
//
// No Go, um "error" e so outro valor que lidamos como qualquer outro
// valor, do jeito que achamos melhor. Nao existe nenhum tipo de keyword
// especial pra lidar com eles
