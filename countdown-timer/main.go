package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds int
	fmt.Print("Введите секунды: ")
	fmt.Scan(&seconds)

	for i := seconds; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Время вышло!")
}
