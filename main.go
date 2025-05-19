package main

import (
	basictypes "github.com/LipeMachado/go-study/basictypes"
	"github.com/LipeMachado/go-study/compoundtypes"
	introduction "github.com/LipeMachado/go-study/introduction"
	"github.com/LipeMachado/go-study/meet"
)

func main() {
	meet.Say("Estudando Go")
	//Introdução
	introduction.Introducao()
	introduction.Variaveis()
	introduction.DeclaracaoCurta()
	introduction.Constantes()
	//Tipos Básicos
	basictypes.BasicTypes()
	basictypes.Int()
	basictypes.Float()
	basictypes.TipoBoleano()
	basictypes.String()
	//Tipos Compostos
	compoundtypes.CompoundTypes()
	compoundtypes.Arrays()
	compoundtypes.Slices()
	compoundtypes.Maps()
}
