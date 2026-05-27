package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func main() {

	var u usuario
	u.nome = "Gustavo"
	u.idade = 28
	fmt.Println(u)

}