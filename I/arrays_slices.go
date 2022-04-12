package main

import "fmt"

func main() {
	/*
		Aqui temos a forma de declaração de arrays
		/*********************************\
		var -> Palavra reservada do sistema
		age -> Nome da váriavel
		[4] -> Numero de espaço do array
		float32 -> Palavra reservada do sistema, usada para definir o tipo da váriavel,
		qual recebe valores com casa decimais ou inteiros
		{5, 7.8, 10, 8.6} -> O array já inicializa com os 4 espaços ocupados
		do tipo float

		OBS : Na programação começa contar partindo do 0,
		0 1 2 3
		/*********************************\
	*/
	var notas [4]float32 = [4]float32{5, 7.8, 10, 8.6}
	fmt.Println("Notas :", notas)
	// função len é usada para retorna o valor de espaço ocupado do array
	fmt.Println("Tamanho preenchido do array de notas: ", len(notas))
	// função cap é usada para retorna a capacidade máxima do array
	fmt.Println("Tamanho da capacidade total do array de notas: ", cap(notas))
	/*
		Aqui temos a forma de declaração de slice, usando a segunda forma de atribuição
		mas poderia ser assim :
			var alunos []string = []string{"Mária", "João", "Enzo", "Pamela"}
		alunos -> Nome da váriavel
		:= -> Operador de atribuição de valor e atribuição do tipo da váriavel
		[] -> Tamanho não definido, o Go atribui essa váriavel com um slice
		string -> Palavra reservada do sistema, usada para definir o tipo da váriavel,
		qual recebe valores do tipo texto, string
		{"Mária", "João", "Enzo", "Pamela"} -> O array já inicializa
		com os 4 espaços ocupados do tipo texto

		OBS : Na programação começa contar partindo do 0,
		0 1 2 3
	*/
	fmt.Println("")
	alunos := []string{"Mária", "João", "Enzo", "Pamela"}
	fmt.Println("Alunos :", alunos)
	// função len é usada para retorna o valor de espaço ocupado do array
	fmt.Println("Tamanho preenchido do array de alunos: ", len(alunos))
	// função cap é usada para retorna a capacidade máxima do array
	fmt.Println("Tamanho da capacidade total do array de alunos: ", cap(alunos))
	/*
		append função que junta slices ou itens e retorna um novo slice com dados adicionados.
		Diferente do array, o slice não tem um limite de tamanho
		como iniciei a váriavel alunos com 4 alunos, depois adicionei 1 aluno
		o Go dobra a capacidade ou seja se rodamos o fmt.Println(cap(alunos)), teremos o resultado 8
		e se rodamos fmt.Println(len(alunos)), retorna o valor 5,
		que é total de espaço que realmente estão sendo usados
	*/
	alunos = append(alunos, "Keven")
	/*
		append(alunos, "Keven") -> Desta forma o append retorna um novo slice
		alunos = append(alunos, "Keven")
		então usamos a váriavel alunos já declarada como um slice de string
		usamos o valor de atribuição =, para que o retorno de append com novo dados e as modificações citadas acima
	*/
	fmt.Println("Tamanho preenchido do array de alunos, após realizar o append: ", len(alunos))
	fmt.Println("Tamanho da capacidade total do array de alunos, após realizar o append: ", cap(alunos))
}
