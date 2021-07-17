package main

import "fmt"

func main() {
	//ex 1
	amarelo := [5]string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := [4]string{"Helena", "Jonas", "José", "Juliana"}

	fmt.Printf("Time Amarelo: %v\n", amarelo)
	fmt.Printf("Time Vermelho: %v\n", vermelho)

	//ex 2
	yellow := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	red := []string{"Helena", "Jonas", "José", "Juliana"}

	red = append(red, "Luis Inácio")
	fmt.Printf("Time vermelho: %v\n", red)
	fmt.Printf("Time amarelo: %v\n", yellow)

	//ex 3

}
