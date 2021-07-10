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
}
