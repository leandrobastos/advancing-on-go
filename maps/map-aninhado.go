package main

import (
	"fmt"
)

func main() {
	funcionarios := map[string]map[string]float64{
		"L": {
			"Leandro":       12678.99,
			"Luis Fernando": 33564.34,
		},
		"M": {
			"Manuela": 7564.01,
			"Maria":   897.99,
		},
		"R": {
			"Renato": 756.87,
		},
	}

	fmt.Println(funcionarios)

	delete(funcionarios, "R")

	fmt.Println("---------------")

	fmt.Println(funcionarios)

	for letra, funcs := range funcionarios {
		fmt.Println(letra, funcs)

		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
		fmt.Println("####################")
	}
}
