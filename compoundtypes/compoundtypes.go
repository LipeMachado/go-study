package compoundtypes

import (
	"fmt"

	"github.com/LipeMachado/go-study/meet"
)

func CompoundTypes() {
	meet.Bar()
	meet.Say("Tipos Compostos")
}

func Arrays() {
	meet.Bar()
	meet.Say("Arrays")
	//Array é uma lista de valores do mesmo tipo
	var gavetas [2]string

	gavetas[0] = "Copos"
	gavetas[1] = "Pratos"
	fmt.Println(gavetas)
}

func Slices() {
	meet.Bar()
	meet.Say("Slices")
	//Slice é um array que pode ser redimensionado
	var gavetas []string
	//Append adiciona um valor ao slice
	gavetas = append(gavetas, "Copos", "Pratos", "Jarras", "Colheres")
	fmt.Println(len(gavetas))
	fmt.Println(gavetas)
	//Dividindo o slice
	//O primeiro valor é o índice inicial e o segundo é o índice final
	fmt.Println(gavetas[0:2])
}

func Maps() {
	meet.Bar()
	meet.Say("Maps")
	//Map é uma lista de valores do mesmo tipo
	//Map é formado por chave e valor
	var pessoas = map[string]int{}
	pessoas["João"] = 20
	pessoas["Maria"] = 25
	pessoas["Pedro"] = 30
	fmt.Println(pessoas)

	if idade, ok := pessoas["João"]; ok {
		fmt.Println("Esta pessoa existe e tem ", idade, "anos. Logo é", ok)
	} else {
		fmt.Println("Esta pessoa não existe. Logo é", ok)
	}

	delete(pessoas, "João")
	fmt.Println(pessoas)
}
