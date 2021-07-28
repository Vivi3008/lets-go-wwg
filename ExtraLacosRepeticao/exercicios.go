package main

import (
	"bufio"
	"fmt"
	"os"
)

//ex 5
//iterar no texto e procurar letras iguais
func main() {
	letraE := "e"
	fmt.Println("Escreva um texto:")
	fmt.Printf("Tanto de letra e %d: \n", Repetir(letraE, len(leiaLinha())))
	fmt.Printf("voce digitou %d\n ", len(leiaLinha()))
}

func leiaLinha() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func Repetir(caractere string, tamanhoArray int) int {
	var repeticoes string
	for i := 0; i < tamanhoArray; i++ {
		repeticoes = repeticoes + caractere
	}
	return int(len(repeticoes))
}
