package main

import "fmt"

func main() {
	/*
		if -> Palavra reservada
			* if ou SE em português,
			* sempre em sua frente tera alguma comparação ou um valor que terá o peso
				verdadeiro ou falso
				true ou false
			* 3 > -1
			* podemos ler assim : 3 maior que -1
			* além do >, temos :
				< - Menor
				>= - Maior ou igual
				<= - Menor ou igual
				== - igual
				!= - diferente
			* Logo 3 > -1 ou 3 maior que -1.
			essa expressão será verdadeira e executara o que está dentro de chaves
	*/
	if 3 > -1 {
		fmt.Println("Verdade")
	}
	// O else serve caso a condição do if seja falsa,
	// será executado que estiver dentro do else
	if 1 > 1 {
		fmt.Println("Verdade")
	} else {
		fmt.Println("Mentira")
	}
}
