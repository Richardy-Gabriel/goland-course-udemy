package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias - INT32 = RUNE
	var num3 rune = 123456
	fmt.Println(num3)

	// BYTE = UINT8 - Números positivos
	var num4 byte = 123
	fmt.Println(num4)

	// Números reais
	var numreal float32 = 123.45
	fmt.Println(numreal)

	var numreal2 float64 = 123456.78
	fmt.Println(numreal2)

	numreal3 := 1234567.89
	fmt.Println(numreal3)

	// String
	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	// Pega o número da tabela ASP dependendo da letra que colocar
	char := '%'
	fmt.Println(char)

	var boleano bool
	fmt.Println(boleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
