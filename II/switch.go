package main

import "fmt"

func main() {

	/*
		Já pensou ter várias opções ?
		1 - Casa
		2 - Trabalho
		8, 5 - Férias
		uma for simples de validar multiplas opções que uma váriavel pode receber e
		usando switch
		case 1:
			caso valor de numero for 1 execute isso
		case 8,5:
			caso valor de numero for 8 ou 5 execute isso
	*/
	numero := 2

	switch numero {
	case 1:
		fmt.Println("Ligando Casa...")

	case 2:
		fmt.Println("Ligando trabalho...")
	case 8, 5:
		fmt.Println("Férias")
	// Caso o valor não atender nenhumas das opções acima, o retorno será esse
	default:
		fmt.Println("Ganhou nada!")
	}
}
