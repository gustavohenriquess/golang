package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	wellcome()
	for {
		menu()
	}
}

func readOption() int {
	var optionSelected int
	fmt.Scan(&optionSelected)
	fmt.Println("")
	fmt.Println("Option selected:", optionSelected)
	fmt.Println("")

	return optionSelected
}

func menu() {

	fmt.Println("")
	fmt.Println("1- Normal Calculate")
	fmt.Println("2- Recursive Calculate")
	fmt.Println("3- Math Lib Calculate")
	fmt.Println("0- Exit")

	option := readOption()

	switch option {
	case 1:
		nth := askFibonacci()
		fmt.Println("O numero escolhido foi", normalCalculate(nth))
	case 2:
		nth := askFibonacci()
		fmt.Println("O numero escolhido foi", recursiveCalculate(nth))
	case 3:
		nth := askFibonacci()
		fmt.Println("O numero escolhido foi", mathLibCalculate(nth))

	case 0:
		fmt.Println("Exit")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		menu()
	}
}

func wellcome() {
	fmt.Println("")
	fmt.Println("Hello!, I'm a program that will help you to calculate the fibonacci sequence")
}

func normalCalculate(nth int) int {
	fibonacci := []int{0, 1}

	if nth < 0 {
		return -1
	}

	for i := 0; i < nth; i++ {
		fibonnaci_length := len(fibonacci)
		fibonacci = append(fibonacci, fibonacci[fibonnaci_length-1]+fibonacci[fibonnaci_length-2])
	}

	return fibonacci[nth]

}

func recursiveCalculate(nth int) int {
	if nth < 0 {
		return -1
	}
	if nth == 0 {
		return 0
	}
	if nth == 1 {
		return 1
	}

	return recursiveCalculate(nth-1) + recursiveCalculate(nth-2)

}

func askFibonacci() int {
	fmt.Println("How many fibonacci numbers do you want to calculate?")
	nth := readOption()
	return nth
}

func mathLibCalculate(nth int) int {
	sqrt := math.Sqrt(5)
	phi := (1 + sqrt) / 2
	diff := math.Pow(phi, float64(nth)) + math.Pow(phi, float64(-nth))
	return int(math.Floor(diff / sqrt))
}
