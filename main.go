package main

import (
	//Pode renomear o pacote colocando atras do "link"...
	"fmt"

	"github.com/LipeMachado/go-study/meet"
)

func main() {
	meet.Say("Hello Lipe!")
	meet.SayHello()

	//Variaveis
	//palavra reservada "var" nome da variavel tipo
	var x int = 10
	fmt.Println(x)

	//Declaração curta
	texto := "Olá"
	texto += "Mundo"
	fmt.Println(texto)

	//Constantes
	const PI float64 = 3.1415
	fmt.Println(PI)
}
