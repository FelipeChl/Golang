package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] int // é assim que se declara um array, adicionando ... o número de elementos que ele terá entre colchetes
	fmt.Println(array1)

	slice := []int {1, 2, 3, 4, 5} // slice tem um tamanho dinâmico, porem, não muda o tipo de dado
	fmt.Println(slice)

	slice = append(slice, 6) // para adicionar um elemento a um slice, utilizamos a função append, passando o slice e o elemento que queremos adicionar
	fmt.Println(slice)
}