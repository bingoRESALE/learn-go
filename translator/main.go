package main

import "fmt"

func main() {
	var temp float64
	var from, to string

	fmt.Print("Температура: ")
	fmt.Scan(&temp)
	fmt.Print("Из (C/F/K): ")
	fmt.Scan(&from)
	fmt.Print("В (C/F/K): ")
	fmt.Scan(&to)

	if from == "C" && to == "F" {
		temp = temp*9/5 + 32
	} else if from == "F" && to == "C" {
		temp = (temp - 32) * 5 / 9
	} else if from == "C" && to == "K" {
		temp = temp + 273.15
	} else if from == "K" && to == "C" {
		temp = temp - 273.15
	} else if from == "F" && to == "K" {
		temp = (temp-32)*5/9 + 273.15
	} else if from == "K" && to == "F" {
		temp = (temp-273.15)*9/5 + 32
	}

	fmt.Printf("Результат: %.2f\n", temp)
}
