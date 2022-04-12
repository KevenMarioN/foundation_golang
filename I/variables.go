package main

import "fmt"

func main() {
	/*
		Aqui temos a forma de declaração onde você "fala" para o Go,
		qual vai ser o tipo dessa váriavel
		/*********************************\
		var -> Palavra reservada do sistema
		age -> Nome da váriavel
		int -> Palavra reservada do sistema, usada para definir o tipo da váriavel,
		qual só recebe valores inteiros
		= -> Operador de atribuição de valor
		22 -> Valor que através do operador =, será atriuido a váriavel age
		/*********************************\
	*/
	var age int = 22
	fmt.Println(age)
	/*
		Aqui temos a forma de declaração onde o Go, escolhe o tipo da váriavel,
		correspondente ao que a váriavel receber
		/*********************************\
		people -> Nome da váriavel
		:= -> Operador de atribuição de valor e atribuição do tipo da váriavel
		"Mário" -> No Go, tipo texto, conhecidos com string, podem ser declarados com
		"" ou ``
		/*********************************\
	*/
	people := "Mário"
	fmt.Println(people)
}
