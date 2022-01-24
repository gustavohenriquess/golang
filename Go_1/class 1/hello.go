package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Gustavo"
	age := 25
	version := 1.1

	fmt.Println("Ola Sr.", name, "sua idade é", age)
	fmt.Println("Este programa esta na versão", version)

	fmt.Println("O tipo da variavel name é", reflect.TypeOf(name))
	fmt.Println("O tipo da variavel age é", reflect.TypeOf(age))
	fmt.Println("O tipo da variavel version é", reflect.TypeOf(version))
}
