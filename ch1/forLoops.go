package main

import "fmt"

func main() {
	// Traditinal for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// Idiomatic Go
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// For loop used as while loop
	// Because the variable i was created above
	// I went with the letter j for this example
	j := 0
	for {
		if j == 10 {
			break
		}
		fmt.Print(j*j, " ")
		j++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -2, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index: ", i, "value: ", v)
	}
}
