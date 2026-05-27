package main

import (
	"fmt"
	"modulo/auxiliar" //importa o pacote auxiliar

	"github.com/badoux/checkmail"
)

func main(){

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("felipechavespsp@gmail.com")
	fmt.Println(erro)
}
	
	

