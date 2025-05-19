package tiposbasicos

import (
	"fmt"

	"github.com/LipeMachado/go-study/meet"
)

const Bar string = "================"

func BasicTypes() {
	fmt.Println(Bar)
	meet.Say("Tipos Básicos")
}

func Int() {
	fmt.Println(Bar)
	meet.Say("Inteiros")
	var indice int8 = 1
	var contador int32 = 20
	var longint int64 = 3000000000000000000
	var byte uint8 = 255
	fmt.Println(indice, "\n", contador, "\n", longint, "\n", byte)
}

func Float() {
	fmt.Println(Bar)
	meet.Say("Floats")
	var salario float32 = 1000.50
	var altura float64 = 1.80
	fmt.Println(salario, "\n", altura)
}

func TipoBoleano() {
	fmt.Println(Bar)
	meet.Say("Boolean")
	var aprovado bool = true
	var reprovado bool = false
	fmt.Println(aprovado, "\n", reprovado)
}

func String() {
	fmt.Println(Bar)
	meet.Say("String")
	var nome string = "Lipe"
	var idade int = 20
	var versao float32 = 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade, "anos")
	fmt.Println("Este programa está na versão", versao)
}
