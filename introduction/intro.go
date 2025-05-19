package introduction

import (
	"fmt"

	"github.com/LipeMachado/go-study/meet"
)

func Introducao() {
	meet.Bar()
	meet.Say("Introdução ao Go")
}

func Variaveis() {
	meet.Bar()
	fmt.Printf("Variáveis\n")
	//palavra reservada "var" nome da variavel tipo
	var x int = 10
	fmt.Println(x)
}

func DeclaracaoCurta() {
	meet.Bar()
	fmt.Printf("Declaração Curta\n")
	texto := "Olá"
	texto += "Mundo"
	fmt.Println(texto)
}

func Constantes() {
	meet.Bar()
	fmt.Printf("Constantes\n")
	const PI float64 = 3.1415
	fmt.Println(PI)
}
