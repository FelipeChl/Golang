package main

import "fmt"

func somar(n1 int, n2 int) int {

	return n1 + n2

}

func calculosMatematicos(n1, n2 int) (int, int) {

	soma := n1 + n2
	subtracao := n1- n2

	return soma, subtracao // go permite mais de um retorno uma vez declarado mais  de um tipo
	
}

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	resultadoSoma,resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}