package main

import (
	"fmt"
)

func main() {
	funcionariosSalarios := map[string]float64{
		"Leandro":       9750.67,
		"Luis Fernando": 10456.99,
		"Alex":          1456.00,
	}

	for f, s := range funcionariosSalarios {
		fmt.Println(f, s)
	}

	fmt.Println("----------------")

	funcionariosSalarios["João"] = 2500.00

	delete(funcionariosSalarios, "Alex")

	for f, s := range funcionariosSalarios {
		fmt.Println(f, s)
	}
}
