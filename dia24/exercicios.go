package main

import (
	"fmt"
	"math"
)

//Metodos

type Dog struct {
	name      string
	dono      string
	brinquedo string
}

type Circulo struct {
	raio float64
}

type conjunto []int

func (d Dog) Apresenta() {
	fmt.Printf("Baw! Meu nome é %s, minha mãe é %s, meu brinquedo preferido é %s!\n", d.name, d.dono, d.brinquedo)
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func (c Circulo) Perimetro() float64 {
	perimetro := 2 * math.Pi * c.raio
	return perimetro
}

func (c conjunto) Some() int {
	soma := 0
	for _, operando := range c {
		soma += operando
	}
	return soma
}

func (c conjunto) Media() float64 {
	soma := float64(c.Some())
	qtd := float64(len(c))
	media := soma / qtd
	return media
}

func main() {
	dog1 := Dog{
		name:      "Chomp",
		dono:      "Viviane",
		brinquedo: "Pato de Borracha",
	}
	dog2 := Dog{
		name:      "Gucci",
		dono:      "Hanna",
		brinquedo: "Ursinho de pelúcia",
	}
	dog1.Apresenta()
	dog2.Apresenta()

	circulo1 := Circulo{
		raio: 65.4,
	}
	area := circulo1.Area()
	perimetro := circulo1.Perimetro()
	fmt.Printf("Circulo: \n\tarea : %.4f\n\t perimetro: %.4f\n\t raio: %.2f\n\t", area, perimetro, circulo1.raio)

	conjunto1 := conjunto{1, 6, 5, 6, 96, 6, 3}
	conjunto2 := conjunto{4, 2, 8}

	var conjuntos = []conjunto{conjunto1, conjunto2}

	for _, valor := range conjuntos {
		soma := valor.Some()
		media := valor.Media()

		fmt.Printf("Soma do conjunto é %d\n", soma)
		fmt.Printf("Media do conjunto é %.2f\n", media)
	}

}
