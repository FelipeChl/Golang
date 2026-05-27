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

	// Arrays Internos
	slice3 := make([] float32, 10, 15) // a função make é utilizada para criar um slice, passando o tipo de dado, o tamanho e a capacidade do slice. O tamanho é o número de elementos que o slice tem atualmente, e a capacidade é o número máximo de elementos que o slice pode ter antes de precisar ser redimensionado.
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // a função len retorna o tamanho do slice
	fmt.Println(cap(slice3)) // a função cap retorna a capacidade do slice
}