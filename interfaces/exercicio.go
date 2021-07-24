package main

import (
	"fmt"
)

type Humano struct {
	nome         string
	corPreferida string
}

type Dog struct {
	nome      string
	brinquedo string
}

type Cat struct {
	nome  string
	sache string
}

//interface
type Apresentador interface {
	Apresente()
}

type Shower interface {
	SayHello()
}

func (h Humano) Apresente() {
	fmt.Printf("Olá! Meu nome é %s e minha cor preferida é %s\n", h.nome, h.corPreferida)
}

func (d Dog) Apresente() {
	fmt.Printf("Au Au! Meu nome é %s e meu brinquedo preferido é %s\n", d.nome, d.brinquedo)
}

func (c Cat) Apresente() {
	fmt.Printf("Miau! Meu nome é %s e meu sache preferido é %s\n", c.nome, c.sache)
}

//exercicio 2
type Animal struct {
	nome     string
	grunhido string
	dono     string
	tipo     string
}

func (a Animal) SayHello() {
	fmt.Printf("%s\t! Meu nome é %s, meu dono é %s e sou um %s\n\t", a.grunhido, a.nome, a.dono, a.tipo)
}

func main() {
	viviane := Humano{
		nome:         "Viviane",
		corPreferida: "Roxo",
	}
	chomp := Dog{
		nome:      "Chomp",
		brinquedo: "Pato",
	}
	cat := Cat{
		nome:  "Deivid",
		sache: "Pelinha",
	}
	cat2 := Cat{
		nome:  "Davy Francisco",
		sache: "Cookies",
	}

	cat3 := Cat{
		nome:  "Giovanna",
		sache: "Lasanha",
	}

	//ex 2
	flamingo := Animal{
		nome:     "Crispie",
		dono:     "Giovanna",
		grunhido: "Ahuhuhu",
		tipo:     "Flamingo",
	}

	cachorrinha := Animal{
		nome:     "Nutella",
		dono:     "Davy e Giovanna",
		grunhido: "Baw au",
		tipo:     "Cachorro",
	}

	leao := Animal{
		nome:     "Thor",
		dono:     "Viviane",
		grunhido: "Auuuu",
		tipo:     "Leão",
	}

	apresentadores := []Apresentador{viviane, chomp, cat, cat2, cat3}

	for _, apresentador := range apresentadores {
		apresentador.Apresente()
	}

	//ex 2
	shower := []Shower{leao, flamingo, cachorrinha}

	for _, value := range shower {
		value.SayHello()
	}

}
