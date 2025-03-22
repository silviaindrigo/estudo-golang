package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do arquivo main!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("silviaindriug8yuh")
	fmt.Println(erro)
	
}