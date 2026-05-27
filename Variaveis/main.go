// aula de variaveis

package main

import "fmt"

func main()  {

	var variavel1 string = "Variavel1" //com o var se declara o tipo de variavel
	variavel2 := "variavel2" // com os := não precisa declarar o tipo ele reconhecera, porem, precisa ser declarada em algum momento

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( //declarar mais de uma variavel usando o var
		variavel3 string = "llalala"
		variavel4 string = "alalal"
	)

	fmt.Println(variavel3,variavel4)
	
	variavel5, variavel6 := "akhdkhbh", "bkdbkd" // declarando mais de uma vairavel com o := 
	fmt.Println(variavel5,variavel6)

}