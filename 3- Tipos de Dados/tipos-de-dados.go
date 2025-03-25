package main

import (
	"errors"
	"fmt"
)

func main (){
	var numero  = 100000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000000
	fmt.Println(numero2)
	
	// alias
	// int32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000000000.45
	fmt.Println(numeroReal2)	

	nuemroReal3 := 12345.67
	fmt.Println(nuemroReal3)

	// FIM NUMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	texto := 6 
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)




}