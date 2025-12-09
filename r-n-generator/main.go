package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(10)
	}

	sort.Ints(numbers)
	fmt.Println(numbers)
}
