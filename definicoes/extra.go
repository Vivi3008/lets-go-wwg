package main

import (
	"fmt"
)

func main() {
	//ex 1
	fmt.Println("----------------------------------")
	amarelo := [5]string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := [4]string{"Helena", "Jonas", "José", "Juliana"}

	fmt.Printf("Time Amarelo: %v\n", amarelo)
	fmt.Printf("Time Vermelho: %v\n", vermelho)

	//ex 2
	fmt.Println("----------------------------------")
	yellow := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	red := []string{"Helena", "Jonas", "José", "Juliana"}

	red = append(red, "Luis Inácio")
	fmt.Printf("Time vermelho: %v\n", red)
	fmt.Printf("Time amarelo: %v\n", yellow)

	//ex 3
	fmt.Println("----------------------------------")
	num := 4
	mult2 := (num % 2) == 0
	mult3 := (num % 3) == 0
	if mult2 {
		fmt.Println("multiplo 2: ", mult2)
	}
	if mult3 {
		fmt.Println("multiplo 3: ", mult3)
	}

	//for
	fmt.Println("----------------------------------")
	for i := 13; i <= 27; i++ {
		fmt.Println(i)
	}

	// for simplificado
	fmt.Println("----------------------------------")
	i := 13
	for i <= 27 {
		fmt.Println(i)
		i++
	}

	//outro tipo
	fmt.Println("----------------------------------")
	seiGo := false

	for seiGo == false {
		fmt.Println("Continue estudando...")
		seiGo = true
	}
	fmt.Println("Sei go? ", seiGo)

	//ex 2
	fmt.Println("----------------------------------")

	horas := 0

	for horas <= 24 {
		fmt.Printf("%.2d:00\n", horas)
		horas++
	}

	//ex 3
	fmt.Println("----------------------------------")
	hora := 0
	for hora <= 3 {
		for minuto := 0; minuto <= 60; minuto++ {
			for segundos := 0; segundos <= 60; minuto++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundos)
				segundos++
			}
			minuto++
		}
		hora++
	}

	// range
	fatia := []string{"zero", "um", "dois", "tres", "quatro"}

	for i, v := range fatia {
		fmt.Println(i, "-", v)
	}

	//ex 4
	listaMercado := []string{"abacate", "sabonete", "azeite", "tomate", "banana"}
	for i, v := range listaMercado {
		fmt.Printf("%d) %s\n", i+1, v)
	}

	//funções
	Hello()
	Cumprimente("viviane", 12)
}

func Hello() {
	fmt.Println("Hello Golang!")
}

func Cumprimente(nome string, hora int) {
	cumprimento := ""

	switch {
	case hora >= 6 && hora < 12:
		cumprimento = "Bom dia!"
	case hora >= 12 && hora < 18:
		cumprimento = "Boa tarde!"
	case hora >= 18 && hora < 24:
		cumprimento = "Boa noite!"
	case hora >= 0 && hora < 6:
		cumprimento = "Boa madrugada!"
	default:
		cumprimento = "Olá"
	}

	fmt.Printf("%s, %s", hora, nome)
}
