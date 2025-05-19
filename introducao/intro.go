package introducao

import (
	"fmt"

	"github.com/LipeMachado/go-study/meet"
)

const Bar string = "================"

func Introducao() {
	fmt.Println(Bar)
	meet.Say("Introdução ao Go")
}

func Variaveis() {
	fmt.Println(Bar)
	fmt.Printf("Variáveis\n")
	//palavra reservada "var" nome da variavel tipo
	var x int = 10
	fmt.Println(x)
}

func DeclaracaoCurta() {
	fmt.Println(Bar)
	fmt.Printf("Declaração Curta\n")
	texto := "Olá"
	texto += "Mundo"
	fmt.Println(texto)
}

func Constantes() {
	fmt.Println(Bar)
	fmt.Printf("Constantes\n")
	const PI float64 = 3.1415
	fmt.Println(PI)
}
