package main

import "fmt"

func main() {
	//ex 1 criando arrays e atribuindo valores
	array := [7]string{"um", "dois", "tres", "quatro", "cinco", "seis", "sete"}

	fmt.Printf("O array é do tipo %T e o seus valores são: %v\n ", array, array)

	//ex 2
	filhos := [2]string{"Giovanna", "Davy"}
	tesouro := filhos

	fmt.Printf("array um : %v\n", filhos)
	fmt.Printf("array dois: %v\n", tesouro)

	//ex 3 criar slices e Juntar dois slices
	values := []int{1, 3, 6, 5, 2, 1}
	value2 := []int{7, 8, 9, 10, 12, 13}

	total := append(values, value2...)
	fmt.Printf("O valor das slices é %v\n", total)

	//ex 4 adicionando mais itens a um slice
	lista := []string{"Abacate", "Mamao", "Laranja"}
	fmt.Println("Slice lista: ", lista)

	lista = append(lista, "Sabao", "Carne", "Pasta de amendoim")
	fmt.Println("Lista: ", lista)

	//ex 5 criando um map e deletando uma entrada do mapa
	cores := map[string]string{
		"cinza":  "#424242",
		"roxo":   "#a378f8",
		"branco": "#ffffff",
	}
	fmt.Printf("Cores %v\n", cores)

	//verificando se um valor no map existe
	azul, existe := cores["azul"]
	fmt.Println("Existe azul? ", azul, existe)

	delete(cores, "cinza")
	fmt.Printf("Cores depois de deletar %v\n", cores)

	//ex 6 criando map
	ano := map[int]string{
		1:  "janeiro",
		2:  "fevereiro",
		3:  "março",
		4:  "abril",
		5:  "maio",
		6:  "junho",
		7:  "julho",
		8:  "agosto",
		9:  "setembro",
		10: "outubro",
		11: "novembro",
		12: "dezembro",
	}
	fmt.Printf("Valor do mes 1 e do ultimo: %v, %v\n", ano[1], ano[12])
	fmt.Println("Tamanho do mapa: ", len(ano))

	//ex 7

	type pessoa struct {
		nome  string
		idade int16
		cor   string
	}

	vivi := pessoa{
		nome:  "Viviane",
		idade: 35,
		cor:   "roxo",
	}

	gigi := pessoa{
		nome:  "Giovanna",
		idade: 9,
		cor:   "rosa",
	}

	bebe := pessoa{
		nome:  "Davy Francisco",
		idade: 7,
		cor:   "azul",
	}

	fmt.Printf("Nome: %v, idade: %v, cor preferida: %v\n", vivi.nome, vivi.idade, vivi.cor)
	fmt.Printf("Nome: %v, idade: %v, cor preferida: %v\n", gigi.nome, gigi.idade, gigi.cor)
	fmt.Printf("Nome: %v, idade: %v, cor preferida: %v\n", bebe.nome, bebe.idade, bebe.cor)

	//aula dia 17
	idade := 5
	if idade > 16 {
		fmt.Println("Pode votar")
		return
	}
	fmt.Println("Nao pode votar")

	//exercicios
	//2
	num := -5
	if num > 0 {
		fmt.Println("Numero positivo")
		return
	}
	fmt.Println("Numero negativo")

	//extra 1
	val1 := 15
	val2 := 10
	val3 := 13
	var maior int

	if val1 > val2 && val1 > val3 {
		maior = val1
	} else if val2 > val3 && val2 > val1 {
		maior = val2
	} else if val3 > val1 && val3 > val2 {
		maior = val3
	}
	fmt.Println("maior valor: ", maior)

}
