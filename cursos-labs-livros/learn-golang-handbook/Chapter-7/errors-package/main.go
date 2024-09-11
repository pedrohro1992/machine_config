package main

import "errors"

// A biblioteca padrao do Go fornece o package
// "error", que deixa mais facil lidar com erros
// Segue um exemplo simples de como utiliza-lo
var err error = errors.New("something went wrong")
