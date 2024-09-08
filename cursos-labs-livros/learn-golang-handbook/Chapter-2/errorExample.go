package main

import "fmt"

func main() {
	var username string = "pramos"
	var password int = 9900

	// Isso vai dar erro porque string e int nao podem ser adicionadas
	fmt.Println("Authorization: Basic", username+":"+password)
}
