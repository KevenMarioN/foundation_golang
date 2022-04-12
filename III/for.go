package main

import "fmt"

func main() {
	/*
		Loops
		for palavra reservada
		i := 0 -> declaração da váriavel e valor inical
		i < 5 -> enquanto i for menor que 5, vai fica executando
		a cada execução do bloco for é verificada essa condição por ultimo
		i++  -> operador de incrementação é a mesma coisa que i recebe i somado com 1
		deve aparecer no terminal
		0
		1
		2
		3
		4
		Cadê o 5?
		lembra da condição i < 5. no final da execução o i é iqual á 5,
		ou seja já satisfaz a condição e não mostra o 5
		no lugar do < colocar o <= o 5 será exibido!
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	carros := []string{"Ford", "Tesla"}
	/*
		usando a palavra reserda range em um slice ou array
		retorna dois tipos de dados
		1 - sendo o index
		2 - sendo o item
		usando o _, é a mesma coisa de não querer receber esse dado

	*/
	for _, b := range carros {
		fmt.Println("Carro: ", b)
	}
}
