package main

import (
	"fmt"
)

func main() {
	var numero int64 = 24
	var nome string = "Vivi"
	var idade float64 = 35.2
	fmt.Printf("Tipo %T e valor %v\n", numero, numero)
	fmt.Printf("Tipo %T e valor %v\n", nome, nome)
	fmt.Printf("Tipo %T e valor %v\n", idade, idade)

	a := 230
	b := 27

	var soma = a + b

	fmt.Printf("O tipo de a é %T e seu valor é %v\n", a, a)
	fmt.Printf("O tipo de b é %T e seu valor é %v\n", b, b)
	fmt.Printf("O tipo de soma é %T e seu valor é %v\n", soma, soma)
	fmt.Println(soma)

	nome2 := "viviane"
	cor := "roxo"

	fmt.Printf("Meu nome é %s e minha cor preferida é %s!\n", nome2, cor)

	anoNascimento := 1986
	anoAtual := 2021

	age := (anoAtual - anoNascimento)

	fmt.Printf("Minha idade é %v\n", age)

	//arrays

	var weekdays [7]string
	weekdays = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(weekdays[2])

	//slices
	var diasDaSemana []string
	fmt.Println("tamanho do slice: ", len(diasDaSemana))

	diasDaSemana = append(diasDaSemana, "Sunday", "Monday")

	//depois do append
	fmt.Println("tamanho do slice:", len(diasDaSemana))

	//make()
	exemplo := make([]string, 7)
	fmt.Println("tamanho: ", len(exemplo))

	exemplo2 := make([]string, 5, 7)
	fmt.Println("tamanho: ", len(exemplo2))
}
