package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Aula struct {
	Descricao string
	Docente   string
	Discentes []Pessoa
}

func main() {
	vivi := Pessoa{Nome: "Viviane", Idade: 35}
	eliana := Pessoa{Nome: "Eliana", Idade: 47}
	joao := Pessoa{Nome: "João", Idade: 19}
	zeca := Pessoa{Nome: "Zeca", Idade: 34}
	aline := Pessoa{Nome: "Aline", Idade: 42}

	fmt.Printf("Pessoa %+v\n", vivi)
	fmt.Printf("Meu nome é %s e tenho %v anos\n", vivi.Nome, vivi.Idade)

	aula := Aula{
		Descricao: "Aula de Golang",
		Docente:   "Xenia",
		Discentes: []Pessoa{
			joao,
			zeca,
			aline,
			eliana,
			vivi,
		},
	}

	fmt.Printf("%+v\n", aula)
	fmt.Println("tamanho do slice discentes", len(aula.Discentes))
}
