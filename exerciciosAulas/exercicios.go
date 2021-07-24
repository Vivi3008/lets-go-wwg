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
	exemplo2[3] = "Wednesday"
	fmt.Println("tamanho do slice: ", len(exemplo2))
	fmt.Println("capacidade do slice: ", cap(exemplo2))

	fmt.Println(exemplo2[:])                //imprime todos valores do slice
	fmt.Println(exemplo2[0:1])              //imprime o valor da posicao 1
	fmt.Println(exemplo2[len(exemplo2)-1:]) //imprime o ultimo elemento do slice

	//crie um slice do tamanho 12 usando make() e atribua os signos do zodiaco a cada um dos elementos individualmente
	//printe na tela essa slice
	signos := make([]string, 12, 12)

	signos[0] = "Aquario"
	signos[1] = "Peixes"
	signos[2] = "Aries"
	fmt.Println("Signos do zodiaco: ", signos[:])
	fmt.Println("tamanho do slice: ", len(signos))

	//crie um mapa onde as chaves(key) sao os numeros dos meses e o valor é o nome do mes
	//printe os meses janeiro e dezembro

	var meses = make(map[int16]string)
	meses[1] = "janeiro"
	meses[3] = "março"
	meses[4] = "abril"
	meses[5] = "maio"
	meses[6] = "junho"
	meses[7] = "JULHO"
	meses[8] = "agosto"
	meses[9] = "setembro"
	meses[10] = "outubro"
	meses[11] = "novembro"
	meses[12] = "dezembro"

	fmt.Println(meses)
	fmt.Println(meses[1])
	fmt.Println(meses[12])
}
