package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	// Ponteiros são variáveis que armazenam o endereço de memória de outra variável. Eles são utilizados para acessar e modificar o valor de uma variável a partir de seu endereço de memória.
	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2) 

	variavel ++
	fmt.Println(variavel, variavel2)

	var variavel3 int 
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // o operador & é utilizado para obter o endereço de memória de uma variável

	fmt.Println(variavel3, ponteiro) // adicionando o * faz a desreferência
}