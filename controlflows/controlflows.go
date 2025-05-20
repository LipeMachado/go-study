package controlflows

import (
	"errors"
	"fmt"
	"time"

	"github.com/LipeMachado/go-study/meet"
)

func thisIsAnError() error {
	return errors.New("Error imprimindo")
}

func ControlFlows() {
	meet.Bar()
	meet.Say("Control Flows")
}

func IfElse() {
	meet.Bar()
	meet.Say("If Else")
	nota := 2

	if err := thisIsAnError(); err != nil {
		if nota >= 7 {
			fmt.Println("Aprovado")
		} else if nota <= 5 && nota >= 3 {
			fmt.Println("Recuperação")
		} else {
			fmt.Println("Reprovado")
		}
		fmt.Println(err)
	}

	players := map[string]int{
		"Ronaldo": 10,
		"Neymar":  9,
	}

	if value, ok := players["Ronaldo"]; ok {
		fmt.Println(value)
	}
}

func SwitchCase() {
	meet.Bar()
	meet.Say("Switch Case")
	fmt.Println("Quando é Sábado?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoje")
	case today + 1:
		fmt.Println("Amanhã")
	case today + 2:
		fmt.Println("Depois de amanhã")
	default:
		fmt.Println("Ta longe ainda")
	}
}
