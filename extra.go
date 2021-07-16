package main

import "fmt"

func main() {
	//ex 1 array preencher?? bom la vai...
	var filhos [3]string
	filhos = [3]string{"Giovanna", "Davy", "Chomp"}
	fmt.Printf("Tenho %v filhos e são %s\n", len(filhos), filhos)

	//ex 3 slices
	var produtos []int
	var qtd []int
	var total []int

	produtos = append(produtos, 32, 35, 23, 65, 32, 1)
	qtd = append(qtd, 1, 3, 2, 5, 6, 2)

	totalProdutos := 0
	totalQtd := 0

	for _, num := range produtos {
		totalProdutos += num
	}

	for _, num := range produtos {
		totalQtd += num
	}

	total = append(total, totalProdutos, totalQtd)

	fmt.Printf("O valor total de produtos é %v e o total de qtd é %v: \n", totalProdutos, totalQtd)
	fmt.Printf("Valor total de tudo %v:\n", total)

}
