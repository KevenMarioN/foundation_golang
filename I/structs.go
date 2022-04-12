package main

import "fmt"

/*
	Usando struct podemos definir a estrutura de uma váriavel
	type -> palavra reservada para declara struct ou interfaces
	alunos -> Nome da estrutura
	struct -> Palavra reservada para gerar um struct
	{} -> delemita oque faz parte da estrutura
	name string -> váriavel do tipo string
	age int -> váriavel do tipo int
*/
type alunos struct {
	name string
	age  int
}

func main() {
	/*
		Usamos a mesma declaração usando int,string.
		Mas ao atribuir chamamos o a struct definida
		atribuindo os dados usando : em vez do atribuidor =
	*/
	var Pamela alunos = alunos{
		name: "Pamela",
		age:  16,
	}
	// Podemos imprimir a váriavel com toda estrutura ...
	fmt.Println("Pamela", Pamela)
	// ou parte dela, se precisar somente do nome
	fmt.Println("Meu nome é:", Pamela.name)
	fmt.Println("")

	/*
		Podemos declarar uma váriavel de forma simplificaada também,
		usando a struct que criamos
	*/
	Mario := alunos{
		name: "Mário",
		age:  22,
	}
	// Podemos imprimir a váriavel com toda estrutura ...
	fmt.Println("Mario", Mario)
	// ou parte dela, se precisar somente da idade
	fmt.Println("Mário tem :", Mario.age, "anos.")

}
